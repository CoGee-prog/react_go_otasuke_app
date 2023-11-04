package middlewares

import (
	"context"
	"net/http"
	"os"

	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
)

// FirebaseApp はFirebaseのアプリインスタンスを保持するためのグローバル変数
var FirebaseApp *firebase.App

// Firebase Admin SDKの初期化
func Init() {
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
}

// Firebaseで認証を行うMiddleware関数
func FirebaseAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// IDトークンをヘッダーから取得
		idToken := c.GetHeader("Authorization")
		if idToken == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			return
		}

		// Firebase Admin SDKを使ってIDトークンを検証
		ctx := context.Background()
		client, err := FirebaseApp.Auth(ctx)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Error getting Auth client"})
			return
		}
		token, err := client.VerifyIDToken(ctx, idToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID token"})
			return
		}

		// 検証が成功したら、リクエストコンテキストにユーザー情報を追加
		c.Set("firebaseUID", token.UID)
		c.Next()
	}
}
