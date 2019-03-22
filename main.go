package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func uploadOne(w http.ResponseWriter, r *http.Request) {
	//判断请求方式
	if r.Method == "POST" {
		//设置内存大小
		r.ParseMultipartForm(32 << 20)
		//获取上传的第一个文件
		file, header, err := r.FormFile("file")
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}
		//创建上传目录
		os.Mkdir("./upload", os.ModePerm)
		//创建上传文件
		cur, err := os.Create("./upload/" + header.Filename)
		defer cur.Close()
		if err != nil {
			log.Fatal(err)
		}
		//把上传文件数据拷贝到我们新建的文件
		io.Copy(cur, file)
	} else {
		//解析模板文件
		t, _ := template.ParseFiles("./uploadOne.html")
		//输出文件数据
		t.Execute(w, nil)
	}
}

func uploadMore(w http.ResponseWriter, r *http.Request) {
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

func main() {
	http.HandleFunc("/uploadMore", uploadMore)
	http.HandleFunc("/uploadOne", uploadOne)

	http.HandleFunc("/shell", exec_shell)

	http.Handle("/", http.FileServer(http.Dir("./static")))

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func exec_shell(writer http.ResponseWriter, request *http.Request) {
	var s = "uname"
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out.String())

}
