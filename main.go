package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/everyday-items/gin-example/library/logging"
	"github.com/everyday-items/gin-example/library/db"
	"github.com/everyday-items/gin-example/library/setting"
	"github.com/everyday-items/gin-example/routers"
	"github.com/gin-gonic/gin"
)

// go run main.go  -env dev
func init() {
	var env string
	flag.StringVar(&env, "env", "prod", "")
	flag.Parse()
	setting.Setup(env)
	logging.Setup()
	db.Setup()
}

func main() {
	gin.SetMode(setting.ServerSetting.RunMode)
	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	server.ListenAndServe()
}
