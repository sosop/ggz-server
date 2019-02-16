package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"
	"net/http"
	"ggz-server/route"
	"github.com/golang/glog"
	"context"
	"flag"
	"ggz-server/store"
	"ggz-server/object"
)


var wait time.Duration

func init() {
	flag.DurationVar(&wait, "graceful-timeout", time.Second * 15, "the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")
	flag.Parse()
}


func main() {

	srv := &http.Server{
		Addr:         "0.0.0.0:8989",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler: route.R,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			glog.Error(err)
		}
		object.WebhookAddr = srv.Addr
	}()

	gracefulStop := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(gracefulStop, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)
	<-gracefulStop


	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	// 关闭数据连接
	store.Close()
	glog.Info("server shutting down")
	os.Exit(0)
}
