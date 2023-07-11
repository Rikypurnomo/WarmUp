package middleware

import (
	"net/http"
	"strconv"
	"strings"

	authdto "github.com/Rikypurnomo/warmup/internal/api/dto/auth"
	"github.com/Rikypurnomo/warmup/internal/api/validator"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusBadRequest, "message": "unauthorized"})
			c.Abort()
			return
		}
		tokenString := strings.Replace(token, "Bearer ", "", 1)
		res, err := validator.VerifyToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": http.StatusUnauthorized, "message": "unauthorized"})
			c.Abort()
			return
		}

		clm, ok := res.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}
		uuid := clm["uuid"].(string)
		result, err := validator.ExtractAuth(c, &authdto.LoginTokenResponse{
			NewUuid: uuid,
		})
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}
		// fmt.Println(uuid, "<<<the uid")
		resultInt, err := strconv.Atoi(result)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			return
		}
		c.Set("userLogin", resultInt)
		c.Next()
	}
}
