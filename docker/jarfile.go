package docker

import (
	"fmt"
	log "github.com/cihub/seelog"
	"html/template"
	"io"
	"net/http"
	"os"
)

func TestOne(w http.ResponseWriter, r *http.Request) {

	fmt.Println("test one method--", r.Method)

}

func UploadOne(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method", r.Method)

	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)

		file, header, err := r.FormFile("file")
		defer file.Close()
		if err != nil {
			log.Error(err)
		}
		os.Mkdir("./upload", os.ModePerm)

		cur, err := os.Create("./upload/" + header.Filename)
		defer cur.Close()
		if err != nil {
			log.Error(err)
		}
		io.Copy(cur, file)

		fmt.Println("upload one end...")

		BuildDockerImages(r.FormValue("appName"), r.FormValue("appVersion"))

	} else {
		t, _ := template.ParseFiles("./static/uploadOne.html")
		t.Execute(w, nil)
	}
}

func UploadMore(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		r.ParseMultipartForm(32 << 20)

		files := r.MultipartForm.File["file"]
		len := len(files)
		for i := 0; i < len; i++ {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				log.Error(err)
			}

			os.Mkdir("./upload", os.ModePerm)

			cur, err := os.Create("./upload/" + files[i].Filename)
			defer cur.Close()
			if err != nil {
				log.Error(err)
			}
			io.Copy(cur, file)
		}
	} else {
		t, _ := template.ParseFiles("./uploadMore.html")
		t.Execute(w, nil)
	}
}
