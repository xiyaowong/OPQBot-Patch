package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/gin-gonic/gin"
)

type SendMsgTask struct {
	Path string
	Data []byte
}

var (
	// 请求转发给bot端的代理
	BotProxy = httputil.ReverseProxy{
		Director: func(req *http.Request) {
			req.URL.Scheme = "http"
			req.URL.Host = fmt.Sprintf("127.0.0.1:%d", BotServerPort)
		},
	}
	// 需要延时发送的消息的队列
	SendMsgQueue = make(chan *SendMsgTask, 1000)
)

func SendMsg() {
	for msg := range SendMsgQueue {
		url := fmt.Sprintf("http://127.0.0.1:%d%s", BotServerPort, msg.Path)
		// fmt.Println("SendMsg: ", url)
		_, err := http.Post(url, "application/json", bytes.NewReader(msg.Data))
		if err != nil {
			log.Printf("SendMsg Request Error: %v\n", err)
		}
		time.Sleep(1 * time.Second)
	}
}

func BotProxyHandler(ctx *gin.Context) {
	if ctx.Query("funcname") == "SendMsg" && ctx.Query("_queue") == "1" {
		if data, err := ioutil.ReadAll(ctx.Request.Body); err == nil {
			task := &SendMsgTask{
				Path: fmt.Sprintf("%s?%s", ctx.Request.URL.Path, ctx.Request.URL.Query().Encode()),
				Data: data,
			}
			SendMsgQueue <- task
		}
		ctx.JSON(200, &Response{0, "ok"})
	} else {
		BotProxy.ServeHTTP(ctx.Writer, ctx.Request)
	}
}
