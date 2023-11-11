package services

import (
	"context"
	"errors"
	"net/http"
	"os"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	"google.golang.org/api/option"
)

type UserService struct {
	db *database.GormDatabase
}

// 対戦相手募集のサービスを作成する
func NewUserService(db *database.GormDatabase) *UserService {
	// Firebase Admin SDKの初期化
	initFirebase()
	return &UserService{
		db: db,
	}
}

// 対戦相手募集の構造体の配列
var Users []*models.User

// Firebaseのアプリインスタンスを保持するためのグローバル変数
var FirebaseApp *firebase.App

// Firebase認証クライアント
var client *auth.Client

// Firebase Admin SDKの初期化
func initFirebase() *firebase.App {
	firebaseConfig := os.Getenv("FIREBASE_CONFIG")
	if firebaseConfig == "" {
		panic("Firebase config is empty")
	}
	opt := option.WithCredentialsJSON([]byte(firebaseConfig))
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		panic("Failed to initialize Firebase app: " + err.Error())
	}
	FirebaseApp = app
	return FirebaseApp
}


var (
	store = sessions.NewCookieStore([]byte("secret-key"))
)

// IDトークンの検証
func VerifyIDToken(c *gin.Context) (*auth.Token, error) {
	// IDトークンをヘッダーから取得
	idToken := c.GetHeader("Authorization")
	if idToken == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
		return nil, errors.New("Authorization header is required")
	}

	// Firebase Admin SDKを使ってIDトークンを検証
	client, err := FirebaseApp.Auth(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Error getting Auth client"})
		return nil, errors.New("Error getting Auth client")
	}
	// IDトークンを検証する
	token, err := client.VerifyIDToken(c, idToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID token"})
		return nil, errors.New("Invalid ID token")
	}
	// 検証が成功したら、リクエストコンテキストにユーザー情報を追加
	c.Set("user_id", token.UID)
	c.Next()
	return token, nil
}

// セッションCookieを作成する
func CreateSessionCookie(c *gin.Context) error {
	// 有効期限を7日間に設定
	expiresIn := time.Hour * 24 * 7
	// セッションCookieを作成
	sessionCookie, err := client.SessionCookie(c, c.GetHeader("Authorization"), expiresIn)
	if err != nil {
		return errors.New("Failed to create a session cookie")
	}

	// セッションCookieをクライアントに設定
	c.SetCookie("session", sessionCookie, int(expiresIn.Seconds()), "/", "", true, true)

	return nil
}

// Firebaseのアプリインスタンスを取得する
func (userService *UserService) GetFireBaseApp() *firebase.App{
	return FirebaseApp
}

// ユーザーを取得する
func (userService *UserService) GetUser(id string) (*models.User, error) {
	var user models.User
	result := userService.db.DB.Where("id = ?", id).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			// レコードが見つからない場合はnilを返す
			return nil, nil
		}
		// その他のエラーの場合
		return nil, result.Error
	}
	// レコードが見つかった場合
	return &user, nil
}

// 新規ユーザーを作成する
func (userService *UserService) CreateUser(user *models.User) error {
	result := userService.db.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
