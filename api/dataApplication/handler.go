package dataApplication

import (
	"authentication/source/model"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	s *Service
}

func (h *Handler) GetJobDecription(c *gin.Context) {
	userClaim, ok := c.Get("currentUser")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User Not Login"})
		return
	}
	user, ok := userClaim.(*model.User)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "User Not Login"})
		return
	}

	resp, err := h.s.GetJobInfo(c, user.ID)
	if err != nil {
		log.Printf("Path: %s, Response:%+v, Error: %s", c.Request.RequestURI, resp, err.Error())
	}
	if resp.Code != 0 {
		// c.HTML(http.StatusUnauthorized, resp.Message, resp)
		c.JSON(http.StatusUnauthorized, resp)

	} else {
		// c.HTML(http.StatusOK, "web.html", gin.H{
		// 	"title": "Dashboard",
		// 	"resp":  resp,
		// })
		c.JSON(http.StatusOK, resp)

	}
}

func (h *Handler) CallBackCheckAuth(c *gin.Context) {
	var request *model.CheckTokenReq

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	resp, err := h.s.CallBackCheckAuth(c, request)
	if err != nil {
		log.Printf("Path: %s, Response:%+v, Error: %s", c.Request.RequestURI, resp, err.Error())
	}
	if resp.Code != 0 {
		c.JSON(http.StatusUnauthorized, resp)
	} else {
		c.JSON(http.StatusCreated, resp)
	}
}

func (h *Handler) Logout(c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("Authorization")
	fields := strings.Fields(authorizationHeader)
	token := ""
	if len(fields) != 0 && fields[0] == "Bearer" {
		token = fields[1]
	}
	resp, err := h.s.Logout(c, token)
	if err != nil {
		c.JSON(http.StatusOK, resp)
		return

	}
	if resp.Code == 0 {
		// c.SetCookie("token", "", -1, "/", "/", false, true)
		c.JSON(http.StatusOK, resp)
		return
	} else {
		c.JSON(http.StatusUnauthorized, resp)
	}

}
func NewHandler(s *Service) *Handler {
	return &Handler{s: s}
}
