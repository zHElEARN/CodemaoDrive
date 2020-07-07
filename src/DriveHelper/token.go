package DriveHelper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Token struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	Description string `json:"description"`
	Data        []struct {
		Token     string `json:"token"`
		BucketUrl string `json:"bucketUrl"`
	} `json:"data"`
}

func GetToken() (string, error) {
	tokenUrl, err := url.Parse("https://api.codemao.cn/api/v2/cdn/upload/token/1")
	if err != nil {
		return "", err
	}

	_query := tokenUrl.Query()
	_query.Add("from", "static")

	client := new(http.Client)
	request, err := http.NewRequest("GET", tokenUrl.String(), nil)
	if err != nil {
		return "", err
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/39.0.2171.71 Safari/537.36")
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var token Token
	err = json.Unmarshal(content, &token)
	if err != nil {
		return "", err
	}
	if token.Code != 200 {
		return "", fmt.Errorf("未知错误")
	}

	return token.Data[0].Token, nil
}
