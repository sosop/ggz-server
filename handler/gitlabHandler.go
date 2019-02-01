package handler

import (
	"net/http"
	"ggz-server/util"
	"ggz-server/object"
	"github.com/sosop/gitlabClient"
	"ggz-server/store"
	"github.com/golang/glog"
	"io/ioutil"
)


func init() {
	// 初始化gitinfo
	gitlabClient.GitInfo = gitlabClient.NewGitlabInfo(nil, "")
}


func CreateGitlab(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil || len(data) == 0 {
		glog.Error(err, "数据为空！")
		util.WriteJsonString(w, object.NewParamErrReturnObj())
		return
	}

	err = util.UnMarshal(data, gitlabClient.GitInfo)
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
		return
	}

	err = store.Store(object.Gitlab, data)
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
		return
	}

	util.WriteJsonString(w, object.NewSuccessReturnObj())
}

func GetGitlab(w http.ResponseWriter, r *http.Request) {
	data, err := store.View(object.Gitlab)
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
		return
	}
	err = util.UnMarshal(data, gitlabClient.GitInfo)
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
		return
	}
	util.WriteJsonString(w, object.NewSuccessWithDataReturnObj(gitlabClient.GitInfo))
}