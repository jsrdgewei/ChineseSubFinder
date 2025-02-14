package middle

import (
	"github.com/allanpk716/ChineseSubFinder/internal/backend/common"
	"github.com/allanpk716/ChineseSubFinder/internal/types/backend"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func CheckAuth() gin.HandlerFunc {

	return func(context *gin.Context) {
		authHeader := context.Request.Header.Get("Authorization")
		if len(authHeader) <= 1 {
			context.JSON(http.StatusUnauthorized, backend.ReplyCheckAuth{Message: "Request Header Authorization Error"})
			context.Abort()
			return
		}
		nowAccessToken := strings.Fields(authHeader)[1]
		if nowAccessToken != common.GetAccessToken() {
			context.JSON(http.StatusUnauthorized, backend.ReplyCheckAuth{Message: "AccessToken Error"})
			context.Abort()
			return
		}
		// 向下传递消息
		context.Next()
	}
}
