package docker

import (
	"fmt"
	log "github.com/cihub/seelog"
	"html/template"
	"io"
	"net/http"
	"os"
)

func UploadOne(w http.ResponseWriter, r *http.Request) {

	fmt.Println("method", r.Method)

	if r.Method == "POST" {
		//设置内存大小
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

		fmt.Print("upload one end...")

		BuildDockerImages(r.FormValue("appName"), r.FormValue("appVersion"))

	} else {
		//解析模板文件
		t, _ := template.ParseFiles("./uploadOne.html")
		//输出文件数据
		t.Execute(w, nil)
	}
}

func UploadMore(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//设置内存大小
		r.ParseMultipartForm(32 << 20)
		//获取上传的文件组
		files := r.MultipartForm.File["file"]
		len := len(files)
		for i := 0; i < len; i++ {
			//打开上传文件
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				log.Fatal(err)
			}
			//创建上传目录
			os.Mkdir("./upload", os.ModePerm)
			//创建上传文件
			cur, err := os.Create("./upload/" + files[i].Filename)
			defer cur.Close()
			if err != nil {
				log.Fatal(err)
			}
			io.Copy(cur, file)
		}
	} else {
		//解析模板文件
		t, _ := template.ParseFiles("./uploadMore.html")
		//输出文件数据
		t.Execute(w, nil)
	}
}
