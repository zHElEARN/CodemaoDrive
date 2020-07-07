package main

import (
	"DriveHelper"
	"bufio"
	"log"
	"os"
	"strings"
)

var (
	Info  *log.Logger
	Error *log.Logger
)

func init() {
	Info = log.New(os.Stdout, "[Info] ", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "[Error] ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	Info.Println("请输入要上传的文件完整路径")
	inputReader := bufio.NewReader(os.Stdin)
	filepath, _ := inputReader.ReadString('\n')
	filepath = strings.Replace(filepath, "\n", "", -1)

	Info.Println("获得Token中...")
	token, err := DriveHelper.GetUploadToken()
	if err != nil {
		Error.Println("获得Token失败, 错误信息", err)
		os.Exit(-1)
	}
	Info.Println("获得Token成功")

	fileInfo, err := DriveHelper.UploadFile(token, filepath)
	if err != nil {
		Error.Println("上传文件错误", err)
		os.Exit(-1)
	}
	Info.Println("上传成功 CD码", fileInfo.BuildUri())
}
