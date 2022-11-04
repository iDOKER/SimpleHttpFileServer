package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello, World!")
	//var basedir = flag.String("share", ".", "Share files in this DIR")
	flag.Parse()

	http.HandleFunc("/", UploadPage)
	http.ListenAndServe("localhost:8080", nil)
}

func UploadPage(w http.ResponseWriter, r *http.Request) {
	uri := r.RequestURI
	uname := r.FormValue("uname")
	fmt.Fprintf(w, uname)
	fmt.Fprintf(w, uri)
}

//func main() {
//	// 解析url地址
//	u, err := url.Parse("http://bing.com/search?q=dotnet")
//	if err != nil {
//		panic(err)
//	}
//
//	// 打印格式化的地址信息
//	fmt.Println(u.Scheme) // 返回协议
//	fmt.Println(u.Host) // 返回域名
//	fmt.Println(u.Path) // 返回路径部分
//	fmt.Println(u.RawQuery) // 返回url的参数部分
//
//	params := u.Query() // 以url.Values数据类型的形式返回url参数部分,可以根据参数名读写参数
//
//	fmt.Println(params.Get("q")) // 读取参数q的值
//}
