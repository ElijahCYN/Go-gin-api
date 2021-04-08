package main

import (
	"fmt"
	"github.com/ElijahCYN/Go-gin-api/pkg/setting"
	"github.com/ElijahCYN/Go-gin-api/routers"
	"net/http"
)

func main()  {
	router := routers.InitRouter()

	s := &http.Server{
		Addr: fmt.Sprintf(":%d", setting.HTTPPort),
		Handler: router,
		ReadTimeout: setting.ReadTimeout,
		WriteTimeout: setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
