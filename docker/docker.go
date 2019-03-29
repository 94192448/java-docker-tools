package docker

import (
	"fmt"
	"golang.org/x/net/websocket"
	"html/template"
	"log"
	"net/http"
)

func Echo(ws *websocket.Conn) {

	var err error

	for {

		var reply string

		//websocket接受信息

		if err = websocket.Message.Receive(ws, &reply); err != nil {

			fmt.Println("receive failed:", err)

			break

		}

		fmt.Println("reveived from client: " + reply)

		msg := "received:" + reply

		fmt.Println("send to client:" + msg)

		for !runEnd {
			if runMsg != "" {
				if err = websocket.Message.Send(ws, runMsg); err != nil {
					//fmt.Println("send failed:", err)
				}
			}
		}

		if runEnd {
			log.Println("Is running end:", runEnd)
		}

	}

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
