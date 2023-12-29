package callback

import (
	"authentication/source/model"
	"authentication/source/utils"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	url        = "http://127.0.0.1:8080/checkAuth"
	urlLougout = "http://127.0.0.1:8080/logout"
)

func CallBackCheckAuth(ctx *gin.Context, req *model.CheckTokenReq) (
	*model.CheckTokenRespond, error) {
	var (
		body       []byte
		statusCode int
	)
	jsonData, err := json.Marshal(req)
	if err != nil {
		return &model.CheckTokenRespond{
			Code:    1,
			Message: "Marshal error",
		}, err
	}
	headers := make(map[string]string)

	body, statusCode, err = utils.PostJson(url, jsonData, headers)
	if err != nil {
		return &model.CheckTokenRespond{
			Code:    2,
			Message: "http error",
		}, err
	}
	if statusCode != http.StatusOK {
		return &model.CheckTokenRespond{
			Code:    3,
			Message: "THE_ACCOUNT_IS_ALREADY_LOGOUT",
		}, errors.New("status != 200")
	}
	resp := &model.CheckTokenRespond{}
	if err := json.Unmarshal(body, resp); err != nil {
		return &model.CheckTokenRespond{
			Code:    4,
			Message: "Unmarshal Error",
		}, err
	}
	return resp, nil

}

func Logout(ctx *gin.Context, token string) (
	*model.CheckTokenRespond, error) {
	var (
		body       []byte
		statusCode int
	)
	req := &model.CheckTokenReq{
		Token: token,
	}
	jsonData, err := json.Marshal(req)
	if err != nil {
		return &model.CheckTokenRespond{
			Code:    1,
			Message: "Marshal error",
		}, err
	}
	headers := make(map[string]string)

	body, statusCode, err = utils.PostJson(urlLougout, jsonData, headers)
	if err != nil {
		return &model.CheckTokenRespond{
			Code:    2,
			Message: "http error",
		}, err
	}
	if statusCode != http.StatusOK {
		return &model.CheckTokenRespond{
			Code:    3,
			Message: "THE_ACCOUNT_IS_ALREADY_LOGOUT",
		}, errors.New("status != 200")
	}
	resp := &model.CheckTokenRespond{}
	if err := json.Unmarshal(body, resp); err != nil {
		return &model.CheckTokenRespond{
			Code:    4,
			Message: "Unmarshal Error",
		}, err
	}
	return resp, nil

}
