package main

import (
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	http.HandleFunc("/uploadone", uploadOne)
	http.ListenAndServe("127.0.0.1:8099", nil)
}

func handlePage(res http.ResponseWriter, req *http.Request) {
	//uri := r.RequestURI
	//fmt.Fprintf(w, uri)
	key := req.FormValue("key")
	fmt.Fprintf(res, key)
}

func uploadOne(w http.ResponseWriter, r *http.Request) {

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

func handleUploadPage(w http.ResponseWriter, req *http.Request) {

	var names []string

	w.WriteHeader(http.StatusOK)
	uname := req.FormValue("uname")
	fmt.Println(uname)

	var userDir = uploadDir + "/" + uname

	fmt.Println(userDir)

	r, err := req.MultipartReader()
	fmt.Println(err)
	if err == nil {
		fmt.Println("2")
		f, err := r.ReadForm(20 * 1024 * 1024)
		if err == nil {
			for k, v := range f.File {
				fmt.Printf("File:%s\n", k)
				for i := 0; i < len(v); i++ {
					var filename string
					for n := 1; true; n++ {
						filename = filepath.Join(uploadDir, fmt.Sprintf("%v-%v", n, v[i].Filename))
						_, err := os.Stat(filename)
						if err != nil {
							break
						}
					}
					fmt.Println(filename)
					of1, _ := os.Create(filename)
					if1, _ := v[i].Open()
					for n, _ := io.Copy(of1, if1); n > 0; n, _ = io.Copy(of1, if1) {
					}
					names = append(names, v[i].Filename)
					fmt.Printf("%s\n", filename)
					of1.Close()
					if1.Close()
				}
			}
		} else {
			log.Println("ReadForm:" + err.Error())
			//resp.Write([]byte("Error"))
			//return
		}
	}
	t := template.New("")
	_, err = t.Parse(`<html><head>
	<meta http-equiv="content-type" content="text/html;charset=utf-8"/>
	<meta name="viewport" content="width=device-width,initial-scale=1.0">
	<title>上传文件</title>
</head>
<body>
<form method="post" action="/uploadone" enctype="multipart/form-data">
    <input type="file" name="file" multiple="multiple" />
    <input type="submit" value="上传">
</form>
</body>
</html>`)
	if err != nil {
		panic(err)
	}
	t.Execute(w, names)
}
