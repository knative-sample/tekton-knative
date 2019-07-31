package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"math/rand"
)

type Log struct {
}

func handler(w http.ResponseWriter, r *http.Request) {
	l:= &Log{}
	l.Info("Hello world received a request.")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "World"
	}

	rand.Seed(time.Now().Unix())

	fmt.Fprintf(w, "Hello %s!\n", target)

	if rand.Intn(10) < 4 {
		l.Error("http handler error")
	} else {
		l.Info("http handler success")
	}

}

func main() {
	l:= &Log{}
	l.Info("Hello world sample started.")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatalf("ListenAndServe error:%s ", err.Error())
	}
}

func (log *Log) Infof(format string, a ...interface{})  {
	log.log("INFO", format, a...)
}

func (log *Log) Info(msg string)  {
	log.log("INFO", "%s", msg)
}

func (log *Log) Errorf(format string, a ...interface{})  {
	log.log("ERROR", format, a...)
}

func (log *Log) Error(msg string)  {
	log.log("ERROR", "%s", msg)
}

func (log *Log) Fatalf(format string, a ...interface{})  {
	log.log("FATAL", format, a...)
}

func (log *Log) Fatal(msg string)  {
	log.log("FATAL", "%s", msg)
}

func (log *Log) log(level, format string, a ...interface{})  {
	var cstSh, _ = time.LoadLocation("Asia/Shanghai")
	ft := fmt.Sprintf("%s %s %s\n", time.Now().In(cstSh).Format("2006-01-02 15:04:05"), level, format)
	fmt.Printf(ft, a...)
}

