# SimpleHttpFileServer

## 简介

本项目主要用于搭建一个基于 web 的文件服务，提供类似于 Apache index 的文件查看功能，同时也提供文件上传入口。

## 安装与使用

可以使用源码编译安装，亦可直接使用编译好的二进制文件直接运行。

### 使用`go`编译

`go get github.com/rocket049/fileserver`

### 直接下载可执行程序

[https://github.com/rocket049/fileserver/releases](https://github.com/rocket049/fileserver/releases)

下载解压后，把可执行程序移动到 `PATH` 中使用。

参数：

```
  -share string
    	Share files in this DIR (default ".")。分享这个目录中的文件。
  -upload string
    	Upload files to this DIR (default ".")。上传文件到这个目录。
```

运行后会显示访问 URL ，输入浏览器就可以下载上传。

同时也会显示一个二维码，可以用手机浏览器扫描访问。

#### 专用图片显示程序

目录中的`showImg`程序是专用图片显示程序，编译后随便软链接到到某个包含在`PATH`中的目录就可以被`fileserver`调用。


## 声明

本项目 Fork 自
[https://github.com/rocket049/fileserver/releases](https://github.com/rocket049/fileserver/releases)
