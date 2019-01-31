package handler

import (
	"net/http"
	"ggz-server/util"
	"ggz-server/object"
	"github.com/sosop/gitlabClient"
	"ggz-server/store"
	"github.com/golang/glog"
	"fmt"
)

func CreateGitlab(w http.ResponseWriter, r *http.Request) {
	gitlabAddr := r.FormValue("gitlabAddr")
	if gitlabAddr == "" {
		glog.Error("gitlabAddr为空")
		util.WriteJsonString(w, object.NewParamErrReturnObj())
		return
	}

	gitlabClient.GitInfo = gitlabClient.NewGitlabInfo(nil, gitlabAddr)

	data, err := util.Marshal(gitlabClient.GitInfo)
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
	fmt.Println("========" + string(data))
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
		return
	}
	util.WriteJsonString(w, object.NewSuccessWithDataReturnObj(data))
}