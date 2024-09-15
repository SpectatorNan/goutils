package miniprogram

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
)

func (c *client) fetchUrl(apiUrl, path string, params map[string]string) string {
	targetPath := fmt.Sprintf("%s%s", apiUrl, path)
	if len(params) == 0 {
		return targetPath
	}
	urlPs := url.Values{}
	for k, v := range params {
		urlPs.Add(k, v)
	}
	baseUrl, _ := url.Parse(targetPath)
	baseUrl.RawQuery = urlPs.Encode()

	return baseUrl.String()
}
func (c *client) getResponse(path string, params map[string]string) (*http.Response, error) {
	// Build the URL with query parameters
	fullURL := c.fetchUrl(apiServer, path, params)

	// Make the GET request
	resp, err := http.Get(fullURL)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func handleResponse[T ResponseErrCode](resp *http.Response) (*T, error) {
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.Errorf("GetAccessToken error, StatusCode = %v", resp.StatusCode)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Unmarshal the response body into AccessTokenResponse
	var result T
	if err = json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if result.GetErrCode() != 0 {
		return nil, errors.Errorf("GetAccessToken error, errcode = %v, errmsg = %v", result.GetErrCode(), result.GetErrMsg())
	}

	return &result, nil
}
