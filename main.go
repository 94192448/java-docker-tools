package main

import (
	"github.com/94192448/java-docker-tools/service"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/uploadMore", service.UploadMore)
	http.HandleFunc("/uploadOne", service.UploadOne)

	http.HandleFunc("/test", service.TestOne)

	http.Handle("/websocket", websocket.Handler(service.Echo))
	http.HandleFunc("/web", service.Web)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
