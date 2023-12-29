package dataApplication

import (
	"authentication/api/dataApplication"
	"authentication/middleware"
	"authentication/source/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Server02() http.Handler {
	router := gin.Default()
	router.Use(gin.Recovery())
	router.GET("/healthchecker", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Implement Google OAuth2 in Golang"})
	})
	s := dataApplication.NewService()
	h := dataApplication.NewHandler(s)
	root := utils.GetPkgRoot()
	router.LoadHTMLGlob(root + "/templates/*.html")
	router.GET("/getData", middleware.CheckAuth(), h.GetJobDecription)
	router.POST("/callBackCheckAuth", h.CallBackCheckAuth)
	router.POST("/logout", h.Logout)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "Dashboard",
		})
	})

	// router.StaticFS("/images", http.Dir("public"))
	log.Fatal(router.Run(":" + "8081"))
	return router
}
