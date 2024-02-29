package middlewares

import (
	"net/http"
	"react_go_otasuke_app/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// リクエストごとに処理が成功した場合はコミット、エラーが発生した場合はロールバックを行うミドルウェア
func Transaction(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// トランザクション開始
		tx := db.Begin()
		if tx.Error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, utils.NewResponse(
				http.StatusInternalServerError,
				"トランザクションの開始に失敗しました",
				nil,
			))
			return
		}

		// コンテキストにトランザクションをセット
		c.Set("tx", tx)

		// リクエストを処理
		c.Next()

		// レスポンスが2xx系以外ならロールバック
		statusCode := c.Writer.Status() 
		if statusCode < 200 || statusCode >= 300 {
			tx.Rollback()
			return
		}

		// 成功した場合はコミット
		tx.Commit()
	}
}
