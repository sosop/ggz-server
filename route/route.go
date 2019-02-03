package route

import (
	"github.com/gorilla/mux"
	"ggz-server/handler"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"ggz-server/middleware"
)

var R *mux.Router

func init() {
	R = mux.NewRouter()

	c := cors.AllowAll()

	R.Handle("/config/global", negroni.New(c, middlewares.ParseFormMiddlerware, negroni.WrapFunc(handler.CreateGitlab))).Methods("POST", "OPTIONS")
	R.Handle("/config/global", negroni.New(c, negroni.WrapFunc(handler.GetGitlab))).Methods("GET", "OPTIONS")
	R.Handle("/config/project/setting/{group}/{token}", negroni.New(c, middlewares.ParseFormMiddlerware, negroni.WrapFunc(handler.CreateGitlabClient))).Methods("POST", "OPTIONS")
	R.Handle("/config/project/setting/{group}", negroni.New(c, middlewares.ParseFormMiddlerware, negroni.WrapFunc(handler.GetTokens))).Methods("GET", "OPTIONS")
	R.Handle("/config/project/setting/{group}/{token}", negroni.New(c, middlewares.ParseFormMiddlerware, negroni.WrapFunc(handler.DelToken))).Methods("DELETE", "OPTIONS")
}
