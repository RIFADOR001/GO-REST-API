package middlewares

import (
	"net/http"

	"example.com/REST-API/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	// The idea for this function is to verify information before other handlers
	// We use abort to prevent sending multiple replies
	// Then if something goes wrong, the code stops here
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	context.Set("userId", userId)

	// This means that the next handler in line continues
	context.Next()
}
