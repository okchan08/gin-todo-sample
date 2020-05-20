package gin

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type SessionInfo struct {
	UserID interface{}
}

func IsAuthenticated() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		sessionInfo := SessionInfo{
			UserID: session.Get("UserID"),
		}

		if sessionInfo.UserID == nil {
			ctx.Redirect(http.StatusMovedPermanently, "/signin")
			ctx.Abort()
		} else {
			ctx.Set("UserID", sessionInfo.UserID)
			ctx.Next()
		}
	}
}
