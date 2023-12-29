package authenition

import (
	"authentication/source/model"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	s *Service
}

func (h *Handler) Login(c *gin.Context) {
	var user *model.LoginUserInput

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	resp, err := h.s.Login(c, user)
	if err != nil {
		log.Printf("Path: %s, Response:%+v, Error: %s", c.Request.RequestURI, resp, err.Error())
	}
	if resp.Code != 0 {
		c.JSON(http.StatusUnauthorized, resp)
	} else {
		// c.SetCookie("token", resp.Result.Token, 30*60, "/", "http://127.0.0.1:801/web", false, true)
		c.JSON(http.StatusOK, resp)
		// c.HTML(http.StatusOK, "web.html", gin.H{
		// 	"title": "Dashboard",
		// 	"resp":  resp.Result.Token,
		// })
	}
}

func (h *Handler) CheckToken(c *gin.Context) {
	var req *model.CheckTokenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	result, err := h.s.CheckToken(c, req)
	if err != nil {
		log.Printf("Path: %s, result:%+v, Error: %s", c.Request.RequestURI, result, err.Error())
	}
	buff, err := json.Marshal(result)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	resp := &model.Response{
		StatusCode: 200,
		Body:       buff,
	}
	if result.Code == 0 {
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusUnauthorized, resp)
	}
}

func (h *Handler) Logout(c *gin.Context) {
	var req *model.CheckTokenReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": err.Error()})
		return
	}
	//checkAuth
	result, err := h.s.CheckToken(c, req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, result)

	}
	if result.Code != 0 {
		c.JSON(http.StatusUnauthorized, result)
	}

	resp, err := h.s.Logout(c, req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, resp)
		return
	}
	if resp.Code == 0 {
		c.JSON(http.StatusOK, resp)
	} else {
		c.JSON(http.StatusUnauthorized, resp)
	}
}

func NewHandler(s *Service) *Handler {
	return &Handler{s: s}
}
