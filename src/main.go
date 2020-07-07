package main

import (
	"DriveHelper"
	"fmt"
	"os"
)

func main() {
	// Logo
	fmt.Println("\n $$$$$$\\                  $$\\                                             $$$$$$$\\            $$\\                      \n$$  __$$\\                 $$ |                                            $$  __$$\\           \\__|                     \n$$ /  \\__| $$$$$$\\   $$$$$$$ | $$$$$$\\  $$$$$$\\$$$$\\   $$$$$$\\   $$$$$$\\  $$ |  $$ | $$$$$$\\  $$\\ $$\\    $$\\  $$$$$$\\  \n$$ |      $$  __$$\\ $$  __$$ |$$  __$$\\ $$  _$$  _$$\\  \\____$$\\ $$  __$$\\ $$ |  $$ |$$  __$$\\ $$ |\\$$\\  $$  |$$  __$$\\ \n$$ |      $$ /  $$ |$$ /  $$ |$$$$$$$$ |$$ / $$ / $$ | $$$$$$$ |$$ /  $$ |$$ |  $$ |$$ |  \\__|$$ | \\$$\\$$  / $$$$$$$$ |\n$$ |  $$\\ $$ |  $$ |$$ |  $$ |$$   ____|$$ | $$ | $$ |$$  __$$ |$$ |  $$ |$$ |  $$ |$$ |      $$ |  \\$$$  /  $$   ____|\n\\$$$$$$  |\\$$$$$$  |\\$$$$$$$ |\\$$$$$$$\\ $$ | $$ | $$ |\\$$$$$$$ |\\$$$$$$  |$$$$$$$  |$$ |      $$ |   \\$  /   \\$$$$$$$\\ \n \\______/  \\______/  \\_______| \\_______|\\__| \\__| \\__| \\_______| \\______/ \\_______/ \\__|      \\__|    \\_/     \\_______|\n                                                                                                                       \n                                                                                                                       \n                                                                                                                       \n")

	var (
		_result  = ""
		filePath = ""
	)

	fmt.Println("输入upload上传文件 输入download下载文件")
	_, _ = fmt.Scan(&_result)
	if _result == "upload" {
		fmt.Println("请输入文件路径: ")
		_, _ = fmt.Scan(&filePath)

		token, err := DriveHelper.GetUploadToken()
		if err != nil {
			fmt.Println("错误", err)
			os.Exit(-1)
		}

		fileInfo, err := DriveHelper.UploadFile(token, filePath)
		if err != nil {
			fmt.Println("错误", err)
			os.Exit(-1)
		}
		fmt.Println("上传成功, CD码", fileInfo.BuildUri())

		_pause := ""
		_, _ = fmt.Scanln(&_pause)
	} else if _result == "download" {
		cdCode := ""
		fmt.Println("请输入cd码: ")
		_, _ = fmt.Scan(&cdCode)

		result, err := DriveHelper.DownloadFile(cdCode)
		if err != nil {
			fmt.Println("错误", err)
			os.Exit(-1)
		}
		fmt.Println("成功下载", result)

		_pause := ""
		_, _ = fmt.Scanln(&_pause)
	} else {
		fmt.Println("错误输入")

		_pause := ""
		_, _ = fmt.Scanln(&_pause)
	}
}
