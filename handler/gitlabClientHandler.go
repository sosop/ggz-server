/**
Package classification gitlab client API

gitlab客户端操作接口

Host: localhost
Version: 0.1.0

swagger:meta
*/
package handler

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/golang/glog"
	"ggz-server/object"
	"ggz-server/util"
	"ggz-server/store"
	"github.com/dgraph-io/badger"
	"github.com/sosop/gitlabClient"
	"io/ioutil"
)

// 初始化gitlabinfo
func init() {
	gitlabClient.GitInfo = gitlabClient.NewGitlabInfo(nil, "")

	data, err := store.View(object.Gitlab)
	if err != nil {
		glog.Error(err)
		if err == badger.ErrKeyNotFound {
			return
		}
		panic(err)
	}
	err = util.UnMarshal(data, gitlabClient.GitInfo)
	if err != nil {
		panic(err)
	}
}

// 初始化gitlab.client
func init() {
	for g := range object.Group {
		tokens, err := getTokens(g)
		if err != nil {
			panic(err)
		}
		for token := range tokens {
			gitlabClient.PushGitlabClient(token)
		}
	}
}


func CreateGitlabClient(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /config/project/setting/{group}/{token} [group token] createGitlabClient
	vars := mux.Vars(r)
	group := vars["group"]
	token := vars["token"]

	if group == "" || token == "" {
		glog.Error("group or token 为空")
		util.WriteJsonString(w, object.NewParamErrReturnObj())
		return
	}

	tokens, err := getTokens(group)
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
		return
	}
	tokens[token] = struct{}{}

	err = saveTokens(group, tokens)
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
	}

	if _, exist := gitlabClient.GitlabClients[token]; !exist {
		gitlabClient.PushGitlabClient(token)
	}

	util.WriteJsonString(w, object.NewSuccessReturnObj())
}

func GetTokens(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	group := vars["group"]

	if group == "" {
		glog.Error("group为空")
		util.WriteJsonString(w, object.NewParamErrReturnObj())
		return
	}
	tokens, err := getTokens(group)
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
		return
	}
	util.WriteJsonString(w, object.NewSuccessWithDataReturnObj(tokens))
}

func DelToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	group := vars["group"]
	token := vars["token"]

	if group == "" || token == "" {
		glog.Error("group or token 为空")
		util.WriteJsonString(w, object.NewParamErrReturnObj())
		return
	}

	tokens, err := getTokens(group)
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
		return
	}

	delete(tokens, token)
	err = saveTokens(group, tokens)
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
		return
	}
	util.WriteJsonString(w, object.NewSuccessWithDataReturnObj(tokens))
}

func SearchProject(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
		return
	}

	// 获取group
	cache := make([]string, 16)
	err = util.UnMarshal(data, &cache)
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
		return
	}

	allTokens := make(object.Set, 16)
	for _, g := range cache {
		// 获取token
		tokens, err := getTokens(g)
		if err != nil {
			glog.Error(err)
			util.WriteJsonString(w, object.NewServerErrReturnObj())
			return
		}
		object.PushSet(allTokens, tokens)
	}

	// 获取所有项目
	allProjs := make([]gitlabClient.ProjectInfo, 0, 1024)
	for token := range allTokens {
		projs, err := gitlabClient.GitlabClients[token].ListProjects()
		if err != nil {
			glog.Error(err)
			util.WriteJsonString(w, object.NewServerErrReturnObj())
			return
		}

		if projs != nil && len(projs) > 0 {
			allProjs = append(allProjs, projs...)
		}
	}
	util.WriteJsonString(w, object.NewSuccessWithDataReturnObj(util.DistictProject(allProjs)))
}

func SelectBranch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	token := vars["token"]

	if id == "" || token == "" {
		glog.Error("project id or token 为空")
		util.WriteJsonString(w, object.NewParamErrReturnObj())
		return
	}

	if git, ok := gitlabClient.GitlabClients[token]; ok {
		data, err := git.ListBranch(id)
		if err != nil {
			glog.Error(err)
			util.WriteJsonString(w, object.NewServerErrReturnObj())
			return
		}
		result := make([]interface{}, 0, 512)
		err = util.UnMarshal(data, &result)
		if err != nil {
			glog.Error(err)
			util.WriteJsonString(w, object.NewServerErrReturnObj())
			return
		}
		util.WriteJsonString(w, object.NewSuccessWithDataReturnObj(result))
		return
	}

	util.WriteJsonString(w, object.NewParamErrReturnObj())

}

func getTokens(group string) (object.Set, error) {
	data, err := store.View(group + object.GitClient)
	if err != nil {
		if err == badger.ErrKeyNotFound {
			return object.Set{}, nil
		}
		return nil, err
	}
	var tokens object.Set
	err = util.UnMarshal(data, &tokens)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

func saveTokens(group string, tokens object.Set) error {
	data, err := util.Marshal(tokens)
	if err != nil {
		return err
	}

	err = store.Store(group + object.GitClient, data)
	if err != nil {
		return err
	}
	return nil
}