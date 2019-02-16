package handler

import (
	"net/http"
	"io/ioutil"
	"github.com/golang/glog"
	"ggz-server/util"
	"ggz-server/object"
	"github.com/sosop/gitlabClient"
	"ggz-server/store"
)

func  CreateProject(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
		return
	}

	project := gitlabClient.ProjectInfo{}
	err = util.UnMarshal(data, &project)
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
		return
	}

	projsBytes, err :=  store.View(object.BuildProjList)
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewParamErrReturnObj())
		return
	}

	projs := make([]gitlabClient.ProjectInfo, 0, 1024)
	err = util.UnMarshal(projsBytes, &projs)
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
		return
	}

	projs = util.DistictProject(append(projs, project))

	projsBytes, err = util.Marshal(projs)
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
		return
	}

	store.Store(object.BuildProjList, projsBytes)
}
