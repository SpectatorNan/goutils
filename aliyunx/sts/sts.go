package sts

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type StsConfig struct {
	AccessKeyId     string
	AccessKeySecret string
	RoleAcs         string
	//BucketEndpoint  string
	EndPoint   string
	Expiration int
}
type StsClient struct {
	ChildAccountKeyId  string
	ChildAccountSecret string
	RoleAcs            string
	EndPoint           string
}

func NewStsClient(config StsConfig) *StsClient {
	return &StsClient{
		ChildAccountKeyId:  config.AccessKeyId,
		ChildAccountSecret: config.AccessKeySecret,
		RoleAcs:            config.RoleAcs,
		EndPoint:           config.EndPoint,
	}
}

/*
@title 获取阿里STS授权认证
@ds		有效时长（秒） 0~3600秒
*/
func (cli *StsClient) GetALiSTSCredentials(ds int) (*STSCredentials, error) {

	url, err := cli.GenerateSignatureUrl("client", fmt.Sprintf("%d", ds))
	if err != nil {
		return nil, err
	}

	data, err := cli.GetStsResponse(url)
	if err != nil {
		return nil, err
	}

	return data, nil
}
func (cli *StsClient) GenerateSignatureUrl(sessionName, durationSeconds string) (string, error) {
	assumeUrl := "SignatureVersion=1.0"
	assumeUrl += "&Format=JSON"
	assumeUrl += "&Timestamp=" + url.QueryEscape(time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	assumeUrl += "&RoleArn=" + url.QueryEscape(cli.RoleAcs)
	assumeUrl += "&RoleSessionName=" + sessionName
	assumeUrl += "&AccessKeyId=" + cli.ChildAccountKeyId
	assumeUrl += "&SignatureMethod=HMAC-SHA1"
	assumeUrl += "&Version=2015-04-01"
	assumeUrl += "&Action=AssumeRole"
	assumeUrl += "&SignatureNonce=" + uuid.New().String()
	assumeUrl += "&DurationSeconds=" + durationSeconds

	// 解析成V type
	signToString, err := url.ParseQuery(assumeUrl)
	if err != nil {
		return "", err
	}

	// URL顺序化
	result := signToString.Encode()

	// 拼接
	StringToSign := "GET" + "&" + "%2F" + "&" + url.QueryEscape(result)

	// HMAC
	hashSign := hmac.New(sha1.New, []byte(cli.ChildAccountSecret+"&"))
	hashSign.Write([]byte(StringToSign))

	// 生成signature
	signature := base64.StdEncoding.EncodeToString(hashSign.Sum(nil))

	// Url 添加signature
	//assumeUrl = "https://sts.cn-shenzhen.aliyuncs.com/?" + assumeUrl + "&Signature=" + url.QueryEscape(signature)
	assumeUrl = cli.EndPoint + "?" + assumeUrl + "&Signature=" + url.QueryEscape(signature)

	return assumeUrl, nil
}

// 请求构造好的URL,获得授权信息
// TODO: 安全认证 HTTPS
func (cli *StsClient) GetStsResponse(url string) (*STSCredentials, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	credentials := parseSTSResponse(body)
	return credentials, err
}

func parseSTSResponse(jsonData []byte) *STSCredentials {
	var jsonObject map[string]interface{}
	err := json.Unmarshal(jsonData, &jsonObject)
	if err != nil {
		fmt.Println(err)
	}

	creStr := jsonObject["Credentials"]

	data, err := json.Marshal(creStr)
	var credentials STSCredentials

	var tempCredential struct {
		AccessKeySecret string
		AccessKeyId     string
		Expiration      string
		SecurityToken   string
	}
	err = json.Unmarshal(data, &tempCredential)
	if err != nil {
		fmt.Println(err)
		return nil //errorx.NewUnknownError("数据签名序列化异常")
	}
	fmt.Println(credentials)
	return &STSCredentials{
		AccessKeySecret: tempCredential.AccessKeySecret,
		AccessKeyId:     tempCredential.AccessKeyId,
		Expiration:      tempCredential.Expiration,
		SecurityToken:   tempCredential.SecurityToken,
	}
}
