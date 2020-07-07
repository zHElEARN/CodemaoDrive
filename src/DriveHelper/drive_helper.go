package DriveHelper

import (
	"encoding/base64"
	"fmt"
	"github.com/guonaihong/gout"
	"github.com/tidwall/gjson"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

type FileHashInfo struct {
	Key  string `json:"key"`
	Hash string `json:"hash"`
}

type FileFullInfo struct {
	FileHashInfo FileHashInfo
	FileName     string
	FileSize     int64
}

type UploadToken string

const (
	UriHeader = "cdrive://"
)

func (fileInfo *FileFullInfo) BuildUri() string {
	detail := fmt.Sprintf("%s|%s|%d", fileInfo.FileHashInfo.Key, fileInfo.FileName, fileInfo.FileSize)
	return fmt.Sprintf("%s%s", UriHeader, base64.StdEncoding.EncodeToString([]byte(detail)))
}

func (fileInfo *FileFullInfo) FromUri(uri string) FileFullInfo {

	uri = strings.Replace(uri, UriHeader, "", -1)
	_uri, err := base64.StdEncoding.DecodeString(uri)
	if err != nil {
		return *fileInfo
	}

	segment := strings.Split(string(_uri), "|")
	{
		fileInfo.FileHashInfo.Key = segment[0]
		fileInfo.FileHashInfo.Hash = segment[0]
		fileInfo.FileName = segment[1]
		fileInfo.FileSize, _ = strconv.ParseInt(segment[2], 10, 64)
	}
	return *fileInfo
}

func FileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	if err != nil {
		return false
	}
	return true
}

func GetUploadToken() (UploadToken, error) {
	var (
		responseBody = ""
	)
	err := gout.
		GET("https://api.codemao.cn/api/v2/cdn/upload/token/1").
		BindBody(&responseBody).
		Do()

	if err != nil {
		return "", err
	}
	return UploadToken(gjson.Get(responseBody, "data.0.token").String()), nil
}

func UploadFile(token UploadToken, file string) (FileFullInfo, error) {
	var (
		statusCode   int
		responseBody string
		fileInfo     FileFullInfo
	)

	if !FileExist(file) {
		return FileFullInfo{}, fmt.Errorf("file not found")
	}

	err := gout.
		POST("https://upload.qiniup.com/").
		SetForm(
			gout.H{
				"token": token,
				"file":  gout.FormFile(file),
			},
		).
		Code(&statusCode).
		BindBody(&responseBody).
		Do()
	if err != nil {
		return FileFullInfo{}, err
	}

	fileInfo.FileHashInfo.Hash = gjson.Get(responseBody, "hash").String()
	fileInfo.FileHashInfo.Key = gjson.Get(responseBody, "hash").String()
	_stat, _ := os.Stat(file)
	fileInfo.FileName = _stat.Name()
	fileInfo.FileSize = _stat.Size()

	return fileInfo, nil
}

func DownloadFile(cdCode string) (bool, error) {
	fileInfo := FileFullInfo{}
	fileInfo.FromUri(cdCode)

	fileUrl, err := url.Parse("https://static.codemao.cn/" + fileInfo.FileHashInfo.Key)
	if err != nil {
		return false, err
	}

	response, err := http.Get(fileUrl.String())
	if err != nil {
		return false, err
	}

	file, err := os.Create(fileInfo.FileName)
	if err != nil {
		return false, err
	}
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return false, err
	}

	return true, nil
}
