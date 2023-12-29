package authentication

import (
	"authentication/api/authenition"
	"authentication/source/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server01() http.Handler {
	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.Use(gin.Recovery())
	router.GET("/healthchecker", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Implement Google OAuth2 in Golang"})
	})

	s := authenition.NewService()
	h := authenition.NewHandler(s)
	utils.SetEnv("config.yaml")
	root := utils.GetPkgRoot()

	router.LoadHTMLGlob(root + "/templates/*.html")

	router.POST("/login", h.Login)
	router.POST("/checkAuth", h.CheckToken)
	router.POST("/logout", h.Logout)

	log.Fatal(router.Run(":" + "8080"))
	return router
}
