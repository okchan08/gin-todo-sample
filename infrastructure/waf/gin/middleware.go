package gin

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func errorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		err := c.Errors.ByType(gin.ErrorTypePublic).Last()
		if err != nil {
			log.Print(err.Err)

			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"Error": err.Error(),
			})
		}
	}
}
