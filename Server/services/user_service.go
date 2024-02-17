package services

import (
	"context"
	"errors"
	"net/http"
	"react_go_otasuke_app/config"
	"react_go_otasuke_app/database"
	"react_go_otasuke_app/models"
	"react_go_otasuke_app/utils"
	"time"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"gorm.io/gorm"
)

type UserService struct {
	db *database.GormDatabase
}

// ユーザーサービスを作成する
func NewUserService(db *database.GormDatabase) *UserService {
	// Firebase Admin SDKの初期化
	initFirebase()
	return &UserService{
		db: db,
	}
}

// Firebaseのアプリインスタンスを保持するためのグローバル変数
var FirebaseApp *firebase.App

// Firebase認証クライアント
var firebaseClient *auth.Client

// Firebase Admin SDKの初期化
func initFirebase() *firebase.App {
	c := config.Get()
	firebaseConfig := c.GetString("firebase.config")
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
	// サービスの変数に代入
	firebaseClient = client
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Error getting Auth client"})
		return nil, errors.New("Error getting Auth client")
	}
	// IDトークンを検証する
	token, err := firebaseClient.VerifyIDToken(c, idToken)
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
	sessionCookie, err := firebaseClient.SessionCookie(c, c.GetHeader("Authorization"), expiresIn)
	if err != nil {
		return errors.New("Failed to create a session cookie")
	}

	conf := config.Get()
	// セッションCookieをクライアントに設定
	c.SetCookie("session", sessionCookie, int(expiresIn.Seconds()), "/", conf.GetString("client.domain"), true, true)
	c.SetSameSite(http.SameSiteLaxMode)

	return nil
}

// Firebaseのセッショントークンを無効化する
func (us *UserService) RevokeRefreshTokens(c *gin.Context) error {
	return firebaseClient.RevokeRefreshTokens(c, utils.GetUserID())
}

// Firebaseのアプリインスタンスを取得する
func (us *UserService) GetFireBaseApp() *firebase.App {
	return FirebaseApp
}

// ユーザーを取得する
func (us *UserService) GetUser(id string) (*models.User, error) {
	var user models.User
	result := us.db.DB.Where("id = ?", id).First(&user)
	// レコードが見つからない場合はnilを返す
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
		// その他のエラーの場合
	} else if result.Error != nil {
		return nil, result.Error
	}
	// レコードが見つかった場合
	return &user, nil
}

// 新規ユーザーを作成する
func (us *UserService) CreateUser(user *models.User) error {
	result := us.db.DB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// 現在のチームを変更する
func (us *UserService) UpdateCurrentTeam(teamId uint) error {
	// チームに所属していなければエラー
	var userTeam models.UserTeam
	result := us.db.DB.Where("user_id = ? AND team_id = ?", utils.GetUserID(), teamId).First(userTeam)
	if result.Error != nil {
		return errors.New("所属チーム以外に切り替えられません")
	}

	result = us.db.DB.Model(&models.User{}).Where("id = ?", utils.GetUserID()).Update("current_team_id", teamId)
	if result.Error != nil {
		return errors.New("チーム切り替えに失敗しました")
	}
	return nil
}
