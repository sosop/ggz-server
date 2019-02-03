package object

import "github.com/sosop/gitlabClient"

const (
	Success			= 0
	ServerERR 		= -1
	ParamERR 		= -2

	Gitlab			= "keyForGitlab"
	GitClient 		= "-keyForTokens"
)

var GitlabClients =  make(map[string]*gitlabClient.GitLabClient, 1024)

type OwnGroup map[int]string

var Group = OwnGroup{0: "信贷组", 1: "消金组", 2: "金融组", 3: "利卡组", 100: "非业务项目"}

type Set map[string]struct{}


// 响应客户端对象
type ReturnObj struct {
	Code 	int 		`json:"code"`
	Msg 	string 		`json:"msg"`
	Data 	interface{} `json:"data"`
}

func NewServerErrReturnObj() ReturnObj {
	return ReturnObj{Code: ServerERR, Msg: "服务器错误"}
}

func NewParamErrReturnObj() ReturnObj {
	return ReturnObj{Code: ParamERR, Msg: "参数错误"}
}

func NewSuccessReturnObj() ReturnObj {
	return ReturnObj{Code: Success, Msg: "success"}
}

func NewSuccessWithDataReturnObj(data interface{}) ReturnObj {
	return ReturnObj{Code: Success, Msg: "success", Data: data}
}

func NewReturnObj(code int, msg string, data interface{}) ReturnObj {
	return ReturnObj{Code: code, Msg: msg, Data: data}
}


