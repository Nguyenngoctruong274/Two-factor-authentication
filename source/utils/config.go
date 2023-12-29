package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"

	"github.com/golang-jwt/jwt"
	"gopkg.in/yaml.v1"
)

func ValidateToken(token string, signedJWTKey string) (jwt.MapClaims, error) {
	tok, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
		}

		return []byte(signedJWTKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalidate token: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	return claims, nil
}

func SetEnv(file string) {
	root := GetPkgRoot()

	yamlBuff, err := ioutil.ReadFile(root + "/" + file)
	if err != nil {
		log.Fatal(err)
	}
	yamlCfg := &YamlCfg{}
	if err = yaml.Unmarshal(yamlBuff, yamlCfg); err != nil {
		log.Fatal(err)
	}
	for k, v := range yamlCfg.Variables {
		os.Setenv(k, v)
	}

}
func GetPkgRoot() (basePath string) {

	pkgName := reflect.TypeOf(TestCtx{}).PkgPath()
	rootPkgName := strings.Split(pkgName, "/")[0]
	_, b, _, _ := runtime.Caller(0)
	basePath = filepath.Dir(b)
	para := strings.Split(basePath, "/")
	var i int
	for i = 0; i < len(para); i++ {
		if para[i] == rootPkgName {
			break
		}
	}
	basePath = strings.Join(para[:i+1], "/")
	return
}

type YamlCfg struct {
	Variables map[string]string `yaml:"Variables"`
}
type TestCtx struct {
	ReqId string
}
