package imgx

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

func fetchImgByUrl(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return bytes, err

}

func FetchBase64RemoteImg(url string) (string, error) {
	bys, err := fetchImgByUrl(url)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bys), nil
}
