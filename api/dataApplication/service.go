package dataApplication

import (
	callback "authentication/source/callBack"
	"authentication/source/db"
	"authentication/source/model"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Service struct {
	db *db.DB
}

func (s *Service) GetJobInfo(ctx *gin.Context, id primitive.ObjectID) (
	*model.JobDescriptionResponse, error) {

	result, err := s.db.CompanyDb.Job(ctx, id)
	if err != nil {
		return &model.JobDescriptionResponse{Code: 1, Message: "GET_JOB_FAILED", Result: model.JobDescription{}}, err
	}

	return &model.JobDescriptionResponse{Code: 0, Message: "GET_JOB_SUCCESS", Result: model.JobDescription{
		Company:  result.Company,
		Age:      result.Age,
		Position: result.Position,
	}}, nil

}

func (s *Service) CallBackCheckAuth(ctx *gin.Context, request *model.CheckTokenReq) (
	*model.CheckTokenRespond, error) {
	resp, err := callback.CallBackCheckAuth(ctx, request)
	if err != nil {
		return &model.CheckTokenRespond{
			Code:    1,
			Message: "Authentication failed",
		}, err
	}
	return resp, nil
}

func (s *Service) Logout(ctx *gin.Context, token string) (
	*model.CheckTokenRespond, error) {
	//checkAuth
	resp, err := callback.Logout(ctx, token)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

// func LogoutUser(ctx *gin.Context) {
// 	ctx.SetCookie("token", "", -1, "/", "localhost", false, true)
// 	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
// }

func NewService() *Service {
	return &Service{
		db: db.AuthenConnection,
	}
}
