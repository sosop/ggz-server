package middlewares

import (
	"net/http"
	"github.com/golang/glog"
	"github.com/urfave/negroni"
	"ggz-server/util"
	"ggz-server/object"
)

var ParseFormMiddlerware =  negroni.HandlerFunc(func (w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	err := r.ParseForm()
	if err != nil {
		glog.Error(err)
		util.WriteJsonString(w, object.NewServerErrReturnObj())
		return
	}
	next.ServeHTTP(w, r)
})
