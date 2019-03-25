package main

import (
	"github.com/94192448/java-docker-tools/docker"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/uploadMore", docker.UploadMore)
	http.HandleFunc("/uploadOne", docker.UploadOne)

	http.HandleFunc("/test", docker.Test)

	http.Handle("/websocket", websocket.Handler(docker.Echo))
	http.HandleFunc("/web", docker.Web)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
