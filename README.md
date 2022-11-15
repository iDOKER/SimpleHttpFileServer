# SimpleHttpFileServer

## 简介

本项目主要用于搭建一个基于 web 的文件服务，提供类似于 Apache index 的文件查看功能，同时也提供文件上传入口。

FUNC:

- [ ] 使用自定义二级目录名上传文件，默认为当前程序运行目录
- [ ] 绑定服务 IP 与端口

TODO:

- [ ] 目录文件查看

## 安装与使用

可以使用源码编译安装，亦可直接使用编译好的二进制文件直接运行。
需要将本目录下的 html 模板文件放在对应目录。

### 使用`go`编译

`go get github.com/iDOKER/SimpleHttpFileServer`

### 直接下载可执行程序

[https://github.com/iDOKER/SimpleHttpFileServer/releases/](https://github.com/iDOKER/SimpleHttpFileServer/releases/)

下载解压后，把可执行程序移动到系统 `PATH` 路径中使用。

参数：

```
  -dir string
    	Share files in this DIR (default ".")。分享这个目录中的文件。类似 index 功能
  -upload string
    	Upload files to this DIR (default ".")。上传文件到这个目录。
```

运行后会显示访问 URL ，输入浏览器就可以下载上传。

## 声明

本项目 参考 自下方项目并依照需求做出改动

[https://github.com/rocket049/fileserver/releases](https://github.com/rocket049/fileserver/releases)
