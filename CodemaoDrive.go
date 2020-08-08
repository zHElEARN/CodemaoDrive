package main

import (
	"CodemaoDrive/Drive"
	"fmt"
	"github.com/mitchellh/colorstring"
	"github.com/schollz/progressbar/v3"
	"os"
)

func init() {
	logo := []string{
		"   _____          _                            _____       _           ",
		"  / ____|        | |                          |  __ \\     (_)          ",
		" | |     ___   __| | ___ _ __ ___   __ _  ___ | |  | |_ __ ___   _____ ",
		" | |    / _ \\ / _` |/ _ \\ '_ ` _ \\ / _` |/ _ \\| |  | | '__| \\ \\ / / _ \\",
		" | |___| (_) | (_| |  __/ | | | | | (_| | (_) | |__| | |  | |\\ V /  __/",
		"  \\_____\\___/ \\__,_|\\___|_| |_| |_|\\__,_|\\___/|_____/|_|  |_| \\_/ \\___|"}
	for _, line := range logo {
		_, _ = colorstring.Println("[blue]" + line)
	}
}

func main() {
	var (
		input      string
		filePath   string
		cdriveCode string
	)

	fmt.Println("Welcome to CodemaoDrive")
MainLabel:
	fmt.Print("请输入你要进行的操作(upload: 上传, download: 下载): ")
	_, _ = fmt.Scan(&input)

	switch input {
	case "upload":
		{
			fmt.Print("请输入要上传的文件路径: ")
			_, _ = fmt.Scan(&filePath)

			fmt.Println("正在获得Token中...")
			token, err := Drive.GetUploadToken()
			if err != nil {
				_, _ = colorstring.Println("[red] 获得Token失败")
				panic(err)
			}

			var (
				bar       *progressbar.ProgressBar
				lastBytes = 0
				first     = true
			)

			ch := Drive.UploadFile(token, filePath)
			for progress := range ch {
				if progress.Error != nil {
					_, _ = colorstring.Println("[red] 上传错误")
					panic(err)
				}
				if progress.Success == true {
					fmt.Println()
					_, _ = colorstring.Println("上传成功 CD码为[green]" + progress.Data.BuildUri())
					_, _ = colorstring.Println("在线链接为: https://static.codemao.cn/" + progress.Data.FileHashInfo.Key)
					os.Exit(0)
				}
				if first {
					bar = progressbar.Default(int64(progress.TotalBytes))
					first = false
				}
				_ = bar.Add(progress.CurrentBytes - lastBytes)
				lastBytes = progress.CurrentBytes
			}

			/*
				fileFullInfo, err := Drive.UploadFileOld(token, filePath)
				if err != nil {
					_, _ = colorstring.Println("[red] 上传文件失败")
					panic(err)
				}

				_, _ = colorstring.Println("上传成功 CD码为[green]" + fileFullInfo.BuildUri())
			*/
		}
	case "download":
		{
			fmt.Print("请输入CD码: ")
			_, _ = fmt.Scan(&cdriveCode)

			result, err := Drive.DownloadFile(cdriveCode)
			if err != nil {
				_, _ = colorstring.Println("[red]下载失败")
				panic(err)
			}
			if result {
				_, _ = colorstring.Println("[blue]下载成功")
			} else {
				_, _ = colorstring.Println("[red]下载失败")
			}
		}
	default:
		{
			_, _ = colorstring.Println("输入错误 请重新输入")
			goto MainLabel
		}
	}

}
