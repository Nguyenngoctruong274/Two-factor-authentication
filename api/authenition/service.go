package authenition

import (
	"authentication/source/auth"
	"authentication/source/db"
	"authentication/source/model"
	"authentication/source/utils"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Service struct {
	db *db.DB
}

// func (s *Service) SignUp(ctx context.Context, request *model.User) (
// 	model.UserResponse, error) {
// 	//validate
// 	if err := request.Validate(); err != nil {
// 		return model.UserResponse{Code: 1, Message: "INVALID_REQUEST"}, err
// 	}
// 	//hash Password

// 	if err := s.db.AccountDb.CreateUser(ctx, request.Email, request.Password); err != nil {
// 		return model.UserResponse{Code: 2, Message: "CREATE_USER_FAILED"}, err
// 	}

// 	return model.UserResponse{Code: 0, Message: "SIGN_UP_SUCCESS"}, nil

// }

func (s *Service) Login(c *gin.Context, user *model.LoginUserInput) (
	model.UserResponse, error) {

	if err := user.Validate(); err != nil {
		return model.UserResponse{Code: 1, Message: "INVALID_REQUEST"}, err
	}
	//Hash Password
	passWord := utils.HashPassword(user.Password)
	//Get User
	account, err := s.db.AccountDb.GetAccount(c, user.Email, passWord)
	if err != nil {
		return model.UserResponse{Code: 3, Message: "Get_USER_FAILED"}, err
	}
	token, err := s.db.AccountDb.GenerateToken(c, account.ID, user.Email, user.Password)
	if err != nil {
		return model.UserResponse{Code: 4, Message: "GEN_TOKEN_FAILED"}, err
	}
	//Insert JWT
	if err := s.db.AccountDb.UpdateJWT(c, account.ID, token.Token); err != nil {
		return model.UserResponse{Code: 5, Message: "UPDATE_TOKEN_FAILED"}, err
	}

	result := model.Result{
		ID:           account.ID,
		Token:        token.Token,
		RefreshToken: token.RefreshToken,
	}
	return model.UserResponse{Code: 0, Message: "SUCCESS", Result: result}, err
}

func (s *Service) CheckToken(c *gin.Context, request *model.CheckTokenReq) (
	*model.CheckTokenRespond, error) {

	jwtSecret := os.Getenv("JWT_SECRET")
	if len(request.Token) == 0 {
		return &model.CheckTokenRespond{Code: http.StatusUnauthorized, Message: "JWT_NOT_FOUND"}, nil
	}

	//verify token
	id, err := auth.ValidateAT(request.Token, jwtSecret)
	if err != nil {
		return &model.CheckTokenRespond{Code: http.StatusUnauthorized, Message: "JWT_INVALID"}, nil
	}

	//checkJWT
	accountJwt, err := s.db.AccountDb.GetAccountByToken(c, request.Token)
	if err != nil {
		return &model.CheckTokenRespond{Code: http.StatusUnauthorized, Message: "JWT_EXPIRED"}, err
	}
	if id != accountJwt.ID {
		return &model.CheckTokenRespond{Code: http.StatusUnauthorized, Message: "THE_ACCOUNT_IS_ALREADY_LOGGED_IN_ON_ANOTHER_DEVICE"}, nil
	}

	return &model.CheckTokenRespond{Code: 0, Message: "SUCCESS"}, err
}

func (s *Service) Logout(c *gin.Context, req *model.CheckTokenReq) (resp *model.LogoutRespond, err error) {

	accountJwt, err := s.db.AccountDb.GetAccountByToken(c, req.Token)
	if err != nil {
		return &model.LogoutRespond{Code: http.StatusUnauthorized, Message: "THE_ACCOUNT_IS_ALREADY_LOGGED_IN_ON_ANOTHER_DEVICE"}, err
	}
	if err := s.db.AccountDb.UpdateJWT(c, accountJwt.ID, ""); err != nil {
		return &model.LogoutRespond{Code: http.StatusUnauthorized, Message: "THE_ACCOUNT_IS_ALREADY_LOGGED_IN_ON_ANOTHER_DEVICE"}, err
	}
	return &model.LogoutRespond{Code: 0, Message: "LOG_OUT_SUCCESS"}, err
}

func NewService() *Service {
	return &Service{
		db: db.AuthenConnection,
	}
}
