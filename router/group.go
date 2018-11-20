package router

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	moehttp "github.com/light4d/yourfs/common/http"
	"github.com/light4d/yourfs/model"
	"net/http"

	"github.com/light4d/yourfs/service"
)

func group(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		group_get(resp, req)
	case http.MethodPost:
		group_post(resp, req)
	case http.MethodPut:
		group_put(resp, req)
	case http.MethodDelete:
		group_delete(resp, req)
	default:
		moehttp.Options(req, resp)
	}
}
func group_get(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}
	filter := moehttp.Getfilter(req)
	filter["type"] = "group"
	gs, err := service.SearchUserorGroup(filter)
	if err != nil {
		result.Error = err.Error()
	} else {
		result.Result = gs
	}
	moehttp.Endresp(result, resp)
}
func group_post(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}

	group := model.User{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		result.Error = err.Error()
		moehttp.Endresp(result, resp)
		return
	}
	err = json.Unmarshal(body, &group)
	if err != nil {
		result.Error = err.Error()
		moehttp.Endresp(result, resp)
		return
	}
	uid := getuid(req)
	groupid, err := service.CreateGroup(uid, group)
	if err != nil {
		result.Error = err.Error()
		moehttp.Endresp(result, resp)
		return
	}
	result.Result = struct {
		ID string
	}{ID: groupid}
	moehttp.Endresp(result, resp)
}

func group_put(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}

	id := req.URL.Query().Get("id")
	if id == "" {
		result.Error = errors.New("id不能为空")
		moehttp.Endresp(result, resp)
		return
	}
	updater := make(map[string]interface{})
	err := moehttp.Unmarshalreqbody(req, &updater)
	if err != nil {
		result.Error = err.Error()
		moehttp.Endresp(result, resp)
		return
	}
	uid := getuid(req)
	err = service.UpdateGroup(uid, id, updater)
	if err != nil {
		result.Error = err.Error()
	}
	moehttp.Endresp(result, resp)
}
func group_delete(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}
	uid := getuid(req)
	groupid := req.URL.Query().Get("id")

	if groupid != "" {
		err := service.DeleteGroup(uid, groupid)
		if err != nil {
			result.Error = err.Error()
		}
		moehttp.Endresp(result, resp)
		return
	} else {
		result.Error = errors.New("whick one do you want to delete?")
		moehttp.Endresp(result, resp)
		return
	}

}
