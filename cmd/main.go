package main

import (
	"flag"
	"fmt"
	"graphs/handlers"
	"graphs/middleware"
	"log"
	"net/http"
	"os"
	"strconv"
)

var (
	port    int
	logfile string
	ver     bool
)

func init() {
	flag.IntVar(&port, "port", 9090, "The port to listen on.")
	flag.StringVar(&logfile, "logfile", "", "Location of the logfile.")
	flag.BoolVar(&ver, "version", false, "Print server version.")
}

const (
	version     = "1.0.0"
	apiVersion  = "v1"
	apiBasePath = "/api/" + apiVersion + "/"

	//http path .
	apiPath = apiBasePath + "task2"
)

func main() {
	flag.Parse()
	if ver {
		fmt.Printf("HTTP Server v%s", version)
		os.Exit(0)
	}
	var logger *log.Logger

	if logfile == "" {
		logger = log.New(os.Stdout, "", log.LstdFlags)
	} else {
		f, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			panic(err)
		}
		logger = log.New(f, "", log.LstdFlags)
	}
	http.Handle(apiPath, middleware.ServiceLoader(http.HandlerFunc(handlers.GraphInput), middleware.RequestMetrics(logger)))
	strPort := ":" + strconv.Itoa(port)
	logger.Printf("started server on :%d", port)
	logger.Fatal(http.ListenAndServe(strPort, nil))
}
