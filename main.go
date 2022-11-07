package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

var baseDir string = "/Users/stan/Downloads"

// var uploadDir string = "/Users/stan/Downloads"
var uploadDir string = "/data/docker/data/volumes/owncloud-docker-server_files/_data/files/"

func main() {
	fmt.Println("Hi, There!")
	//flag.StringVar(&baseDir, "base", "/opt", "Share files in this directory")
	//flag.StringVar(&uploadDir, "upload", "/Users/stan/Downloads", "Upload files to this directory")
	flag.Parse()

	http.HandleFunc("/", handlePage)
	http.HandleFunc("/upload", handleUploadPage)
	http.ListenAndServe("127.0.0.1:8099", nil)
}

func handlePage(res http.ResponseWriter, req *http.Request) {
	//uri := r.RequestURI
	//fmt.Fprintf(w, uri)
	key := req.FormValue("key")
	fmt.Fprintf(res, key)
}

func handleUploadPage(w http.ResponseWriter, r *http.Request) {

	UserName := r.FormValue("uname")
	fmt.Println(UserName)
	SysFlag := 0

	//判断请求方式
	if r.Method == "POST" {
		//设置内存大小
		r.ParseMultipartForm(32 << 20)
		w.Write([]byte("System Setup ok."))
		//获取上传的第一个文件
		file, header, err := r.FormFile("file")
		// 判断文件有效性
		if err != nil {
			w.Write([]byte("File check error"))
		} else {
			defer file.Close()
			//创建上传目录
			os.Mkdir(uploadDir, os.ModePerm)
			//创建上传文件
			cur, err := os.Create(uploadDir + UserName + "/files/" + header.Filename)
			if err != nil {
				println(err.Error())
			} else {
				defer cur.Close()
				//把上传文件数据拷贝到我们新建的文件
				io.Copy(cur, file)
				w.Write([]byte("Upload SUCCESS"))
				SysFlag = 1
			}
		}
	}

	if SysFlag == 0 {
		t, errPar := template.ParseFiles("./html/upload.html")
		if errPar != nil {
			fmt.Println("ErrorPar = ", errPar)
			return
		}
		errExe := t.Execute(w, UserName)
		if errExe != nil {
			fmt.Println("ErrorExe = ", errExe)
			return
		}
		SysFlag = 0
	}
}
