package DriveHelper

import (
	"encoding/json"
	"github.com/guonaihong/gout"
)

type ResponseFile struct {
	Hash string `json:"hash"`
	Key  string `json:"key"`
}

func (responseFile *ResponseFile) String() string {
	content, _ := json.Marshal(responseFile)
	return string(content)
}

func Upload(token string, file string) (ResponseFile, error) {
	var (
		code         int
		response     []byte
		responseFile ResponseFile
	)
	err := gout.
		POST("https://upload.qiniup.com/").
		SetForm(
			gout.H{
				"token": token,
				"file":  gout.FormFile(file),
			},
		).
		Code(&code).
		BindBody(&response).
		Do()

	if err != nil {
		return ResponseFile{}, err
	}

	err = json.Unmarshal(response, &responseFile)
	if err != nil {
		return ResponseFile{}, err
	}

	return responseFile, nil
}
