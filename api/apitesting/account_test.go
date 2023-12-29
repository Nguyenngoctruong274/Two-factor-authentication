package apitesting

import (
	"authentication/api/authenition"
	"authentication/api/dataApplication"
	"authentication/middleware"
	callback "authentication/source/callBack"
	"authentication/source/db"
	"authentication/source/model"
	AUtils "authentication/source/utils"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	AuthenURI string = "mongodb://localhost:27017/authentication"
	AuthenDB  string = "authentication"
)

func init() {
	db.AuthenConnection = db.NewConnection(
		AuthenURI, AuthenDB, nil, 15*time.Second)
}

func Test_HashPassword(t *testing.T) {
	password := "FEOL@123"
	passwordHash := AUtils.HashPassword(password)
	t.Fatal(passwordHash)
}

func Test_LogIn(t *testing.T) {
	AUtils.SetEnv("config.yaml")
	c := &gin.Context{}
	account := []*model.LoginUserInput{
		{
			Email:    "truong@gmail.com",
			Password: "uMg3ZS8EsaHgFLJDrvCwaoUakhx7Y2GkStfS3GzTjI8=",
		},
		{
			Email:    "anhthinh@gmail.com",
			Password: "uMg3ZS8EsaHgFLJDrvCwaoUakhx7Y2GkStfS3GzTjI8=",
		},
	}
	for _, ac := range account {
		resp, err := authenition.NewService().Login(c, ac)
		if err != nil {
			t.Fatal(err)
		}
		log.Println(resp)
	}
}

func Test_GetData(t *testing.T) {
	id, _ := primitive.ObjectIDFromHex("649d5ba3ddb0294d80497697")
	user := model.User{
		ID:       id,
		Email:    "truong@gmail.com",
		Password: "uMg3ZS8EsaHgFLJDrvCwaoUakhx7Y2GkStfS3GzTjI8=",
	}
	c := &gin.Context{}
	c.Set("currentUser", user)
	resp, err := dataApplication.NewService().GetJobInfo(c, user.ID)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(resp)
}

func TestCheckAuth(t *testing.T) {
	AUtils.SetEnv("config.yaml")
	// Tạo một router mới của Gin
	router := gin.Default()

	// Đăng ký middleware CheckAuth()
	router.Use(middleware.CheckAuth())

	// Định nghĩa một route để kiểm tra hàm CheckAuth()
	router.GET("/protected", func(ctx *gin.Context) {
		user, exists := ctx.Get("currentUser")
		if !exists {
			t.Error("Expected 'currentUser' to be set in the context")
		}
		t.Log("Current user:", user)
		ctx.Status(http.StatusOK)
	})

	// Tạo một HTTP request với header Authorization và cookie token
	req := httptest.NewRequest(http.MethodGet, "/protected", nil)
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NDlkNWJhM2RkYjAyOTRkODA0OTc2OTciLCJleHAiOjE2OTIxNzEyMzgsImlhdCI6MTY4OTU3OTIzOH0.8N0QZTlTh_4TMxy2eJ0Rm1B5Lhm-WfVcyZ2y97mk_lM")
	req.AddCookie(&http.Cookie{Name: "token", Value: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NDlkNWJhM2RkYjAyOTRkODA0OTc2OTciLCJleHAiOjE2OTIxNzEyMzgsImlhdCI6MTY4OTU3OTIzOH0.8N0QZTlTh_4TMxy2eJ0Rm1B5Lhm-WfVcyZ2y97mk_lM"})

	// Tạo một HTTP response recorder để nhận response từ router
	recorder := httptest.NewRecorder()

	// Gửi request đến router
	router.ServeHTTP(recorder, req)

	// Kiểm tra response status code
	if recorder.Code != http.StatusOK {
		t.Errorf("Expected status %d, but got %d", http.StatusOK, recorder.Code)
	}

	// Kiểm tra nội dung log
	expectedLog := "Current user: YOUR_USER"
	if !strings.Contains(recorder.Body.String(), expectedLog) {
		t.Errorf("Expected log '%s' not found in response body", expectedLog)
	}
}

func Test_Token(t *testing.T) {
	AUtils.SetEnv("config.yaml")

	c := &gin.Context{}
	req := &model.CheckTokenReq{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NDlkNWJhM2RkYjAyOTRkODA0OTc2OTciLCJleHAiOjE2OTIxNzY1ODEsImlhdCI6MTY4OTU4NDU4MX0.FUvT_aKhQ7NFoDEsGmpw_4lgLgoh2PYIRI49PTsWmj8"}

	resp, err := authenition.NewService().CheckToken(c, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(resp)
}
func Test_CallBackLogout(t *testing.T) {
	AUtils.SetEnv("config.yaml")

	ctx := &gin.Context{}
	token :=
		"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NDlkNWJhM2RkYjAyOTRkODA0OTc2OTciLCJleHAiOjE2ODk5MjUwMjEsImlhdCI6MTY4OTgzODYyMX0.EpAR4V7k1kzWlypGThQMK4jC9qOt_2fIuk5RFKwTU1I"

	resp, err := dataApplication.NewService().Logout(ctx, token)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(resp)
}

func Test_CallBackCheckAuth(t *testing.T) {
	c := &gin.Context{}
	req := &model.CheckTokenReq{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJfaWQiOiI2NDlkNWJhM2RkYjAyOTRkODA0OTc2OTciLCJleHAiOjE2OTIyNTg4MjIsImlhdCI6MTY4OTY2NjgyMn0.GeIZ-nPj69EJQ0cNK_ezZBC-m6qDxpMnE6dj8reQYcs",
	}
	resp, err := callback.CallBackCheckAuth(c, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal(resp)
}
