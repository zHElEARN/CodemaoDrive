package main

import (
	"DriveHelper"
	"Utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

type METHOD uint8

const (
	UPLOAD   METHOD = 0
	DOWNLOAD METHOD = 1
)

func init() {
	Info = log.New(os.Stdout, "[Info] ", log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, "[Warning] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "[Error] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	var (
		method   METHOD
	)

	var _t int
	Info.Println("请输入模式(上传: 0/下载: 1)")
	_, _ = fmt.Scan(&_t)
	if _t == 0 {
		method = UPLOAD
	} else if _t == 1 {
		method = DOWNLOAD
	} else {
		Warning.Println("输出错误 将默认选择上传(0)")
		method = UPLOAD
	}

	if method == UPLOAD {
	UploadLabel:
		Info.Println("请输入要上传的文件完整路径")
		inputReader := bufio.NewReader(os.Stdin)
		filepath, _ := inputReader.ReadString('\n')
		filepath = strings.Replace(filepath, "\n", "", -1)

		if r, err := Utils.PathExists(filepath); !r {
			Warning.Println("未找到文件, 请重新输入", err)
			goto UploadLabel
		}

		Info.Println("获得Token中...")
		token, err := DriveHelper.GetToken()
		if err != nil {
			Error.Println("获得Token失败, 错误信息", err)
			os.Exit(-1)
		}
		Info.Println("获得Token成功", token)

		response, err := DriveHelper.Upload(token, filepath)
		if err != nil {
			Error.Println("上传文件错误", err)
			os.Exit(-1)
		}
		Info.Println(response.String())
	}
}
