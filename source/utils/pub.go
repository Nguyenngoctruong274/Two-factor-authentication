package utils

import (
	"authentication/source/model"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"
)

func Curl(url string, option *model.Option) (*model.Response, error) {
	if option == nil {
		return nil, ErrorWithStr("option must not be nil")
	}
	if strings.HasPrefix(url, "ws") {
		return nil, ErrorWithStr("the `websocket` protocol is not supported")
	}
	if strings.HasPrefix(url, "http") {
		//pass
	} else {
		url = "http://" + url
	}
	if len(option.Params) > 0 {
		url += "/" + option.Params
	}
	request, err := http.NewRequest(option.Method, url, option.Body)
	if err != nil {
		return nil, ErrorWithStr(err.Error())
	}
	for key, arr := range option.Header {
		for i, value := range arr {
			if i == 0 {
				request.Header.Set(key, value)
			} else {
				request.Header.Add(key, value)
			}
		}
	}
	client := &http.Client{
		Timeout: option.Timeout,
	}
	response, err := client.Do(request)
	if err != nil {
		return nil, ErrorWithStr(err.Error())
	}
	buff, err := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return &model.Response{
		StatusCode:    response.StatusCode,
		Header:        response.Header.Clone(),
		Body:          buff,
		ContentLength: request.ContentLength,
	}, err
}

func PostJson(url string, jsonByte []byte, reqHeader map[string]string) ([]byte, int, error) {
	headers := make(map[string][]string)
	headers["content-type"] = []string{"application/json"}
	for k, v := range reqHeader {
		headers[k] = []string{v}
	}
	body := bytes.NewBuffer(jsonByte)
	opt := &model.Option{
		Method:  http.MethodPost,
		Header:  headers,
		Body:    body,
		Timeout: 120 * time.Second,
	}
	resp, err := Curl(url, opt)
	if err != nil {
		return nil, -1, err
	}
	return resp.Body, resp.StatusCode, err
}

func ErrorWithStr(message string) (logErr error) {
	file, line := getFileLine()
	return fmt.Errorf(fmt.Sprintf("%v line:%v | %v", file, line, message))
}

func getFileLine() (file string, line int) {
	_, file, line, _ = runtime.Caller(2)
	para := strings.Split(file, "/")
	size := len(para)
	if size > 2 {
		file = fmt.Sprintf("%v/%v", para[size-2], para[size-1])
	}
	return
}
func StartServer(addr string, server http.Handler) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", server.ServeHTTP)

	return http.ListenAndServe(addr, mux)
}
