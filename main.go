package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)

var baseDir string = "/Users/stan/Downloads"

// var uploadDir string = "/Users/stan/Downloads"
var uploadDir string = "/data/docker/data/volumes/owncloud-docker-server_files/_data/files/"

func main() {
	fmt.Println("Hi, There!")
	// flag.StringVar(&uploadDir, "upload", "/Users/stan/Downloads", "Upload files to this directory")
	flag.Parse()

	http.HandleFunc("/", handlePage)
	http.HandleFunc("/upload", handleUploadPage)
	err := http.ListenAndServe("127.0.0.1:8099", nil)
	if err != nil {
		return
	}
}

func handlePage(res http.ResponseWriter, req *http.Request) {
	key := req.FormValue("key")
	if key == "" {
		http.Error(res, "Key is empty", http.StatusBadRequest)
		return
	}
	_, err := res.Write([]byte(key))
	if err != nil {
		http.Error(res, "Error writing response", http.StatusInternalServerError)
	}
}

func handleUploadPage(w http.ResponseWriter, r *http.Request) {

	UserName := r.FormValue("uname")
	fmt.Println(UserName)
	SysFlag := 0

	//判断请求方式
	if r.Method == "POST" {
		//设置内存大小
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}
		_, err = w.Write([]byte("System Setup ok."))
		if err != nil {
			return
		}
		//获取上传的第一个文件
		file, header, err := r.FormFile("file")
		// 判断文件有效性
		if err != nil {
			http.Error(w, "Error retrieving file", http.StatusBadRequest)
			return
		} else {
			defer func(file multipart.File) {
				err := file.Close()
				if err != nil {
					log.Print(err)
				}
			}(file)
			//创建上传目录
			err := os.Mkdir(uploadDir, os.ModePerm)
			if err != nil {
				return
			}
			//创建上传文件
			cur, err := os.Create(uploadDir + UserName + "/files/" + header.Filename)
			if err != nil {
				println(err.Error())
				http.Error(w, "Error creating file", http.StatusInternalServerError)
				return
			} else {
				defer func(cur *os.File) {
					err := cur.Close()
					if err != nil {
						log.Print(err)
					}
				}(cur)
				//把上传文件数据拷贝到我们新建的文件
				_, err = io.Copy(cur, file)
				if err != nil {
					http.Error(w, "Error copying file", http.StatusInternalServerError)
					return
				}
				_, err = w.Write([]byte("Upload SUCCESS"))
				if err != nil {
					http.Error(w, "Error writing response", http.StatusInternalServerError)
					return
				}
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
