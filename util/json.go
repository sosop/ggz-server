package util

import (
	"github.com/json-iterator/go"
	"net/http"
	"github.com/golang/glog"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func Marshal(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

func UnMarshal(data []byte, obj interface{}) error {
	return json.Unmarshal(data, obj)
}

func WriteJsonString(w http.ResponseWriter, obj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	data, err := Marshal(obj)
	if err != nil {
		glog.Error(err)
		return
	}
	w.Write(data)
}