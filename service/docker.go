package service

import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
)

func Echo(ws *websocket.Conn) {

	var err error

	var reply string

	if err = websocket.Message.Receive(ws, &reply); err != nil {

		fmt.Println("receive failed:", err)

		return

	}

	fmt.Println("reveived from client: " + reply)

	msg := "received:" + reply

	fmt.Println("send to client:" + msg)

	for {
		if runMsg != "" {
			if err = websocket.Message.Send(ws, runMsg); err != nil {
				log.Println("send failed:", err)
			}
			runMsg = ""
		}
	}
	log.Println("websocket end...")
}

func Web(w http.ResponseWriter, r *http.Request) {

	//打印请求的方法

	fmt.Println("method", r.Method)

	if r.Method == "GET" { //如果请求方法为get显示login.html,并相应给前端

		t, _ := template.ParseFiles("static/websocket.html")

		t.Execute(w, nil)

	} else {

		//否则走打印输出post接受的参数username和password

		fmt.Println(r.PostFormValue("username"))

		fmt.Println(r.PostFormValue("password"))

	}

}
