package router

import (
	"encoding/json"

	"io/ioutil"

	moehttp "github.com/light4d/yourfs/common/http"
	"github.com/light4d/yourfs/model"
	"net/http"

	"github.com/light4d/yourfs/service"
)

func login(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodPost:
		login_post(resp, req)
	default:
		moehttp.Options(req, resp)
	}
}

func login_post(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}
	m := map[string]string{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		result.Error = err.Error()
		moehttp.Endresp(result, resp)
		return
	}
	err = json.Unmarshal(body, &m)
	if err != nil {
		result.Error = err
		moehttp.Endresp(result, resp)
		return
	}
	t, err := service.Login(m["id"], m["password"])
	if err != nil {
		result.Error = err.Error()
		moehttp.Endresp(result, resp)
		return
	}
	result.Result = struct {
		UserID string
		Token  string
	}{UserID: m["id"], Token: t}

	moehttp.Endresp(result, resp)
}
