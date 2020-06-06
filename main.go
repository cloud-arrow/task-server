package main

import (
	"fmt"
	"github.com/cloud-arrow/task-server/pkg/config"
	"github.com/cloud-arrow/task-server/router"
	"net/http"
	"strconv"
	"time"
)

func main() {
	addr := fmt.Sprintf(":%s", strconv.Itoa(config.Config.Server.Httpport))
	s := &http.Server{
		Addr:           addr,
		Handler:        router.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
