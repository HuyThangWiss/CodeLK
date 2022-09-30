package middlewaresPrac
import (
	"awesomeProject/Practice/AuthPrac"
	"github.com/gin-gonic/gin"
)
func AuthLt() gin.HandlerFunc{
	return func(context *gin.Context) {
		tokenString := context.GetHeader("User")
		if tokenString == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		err:= AuthPrac.ValidateTokenLt(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": err.Error()})
			context.Abort()
			return
		}
		context.Next()
	}
}
