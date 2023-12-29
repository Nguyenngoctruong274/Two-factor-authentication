package middleware

import (
	callback "authentication/source/callBack"
	"authentication/source/model"
	"authentication/source/utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var token string
		// cookie, err := c.Cookie("auth")
		authorizationHeader := c.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)
		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}
		// else if err == nil {
		// 	token = cookie
		// }
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Token not exist"})
			return
		}
		//callBackCheckAuth
		resp, err := callback.CallBackCheckAuth(c, &model.CheckTokenReq{
			Token: token,
		})
		if err != nil || resp.Code != 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		// ValidateTolen
		jwtSecret := os.Getenv("JWT_SECRET")
		userClaims, err := utils.ValidateToken(token, jwtSecret)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		id, err := primitive.ObjectIDFromHex(userClaims["_id"].(string))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}
		user := &model.User{
			ID: id,
		}

		c.Set("currentUser", user)
		c.Next()
	}
}

// user, err := db.AuthenConnection.AccountDb.GetAccountById(c, mapClaim["_id"].(string))
// if err != nil {
// 	c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
// 	return
// }
/*
 */
