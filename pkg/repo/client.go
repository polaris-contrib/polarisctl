package repo

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"

	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/spf13/viper"
)

// apiClient http api
var apiClient *ApiClient

// PolarisURL 资源的 http url 前缀
type PolarisURL = string

// url_v1 polaris api v1 url
const url_v1 string = "/naming/v1/"
const (
	//KNamespaceUrl namespaces 操作的 url 前缀
	NamespaceURL    PolarisURL = url_v1 + "namespaces"
	NamespaceDelURL PolarisURL = url_v1 + "namespaces/delete"
)

// ApiClient http 请求处理
type ApiClient struct {
	host       string
	token      string
	httpClient *http.Client

	req  *http.Request
	murl *url.URL
}

// GetApiClient 获取 http 全局句柄
func GetApiClient() *ApiClient {
	return apiClient
}

// InitApiClient 构造 http client
func InitApiClient(cluster entity.PolarisClusterConf) {
	apiClient = &ApiClient{
		host:       cluster.Host,
		token:      cluster.Token,
		httpClient: &http.Client{},
	}
}

// BuildURL 构造 net.url
func (client *ApiClient) BuildURL(value PolarisURL) *url.URL {
	var err error
	client.murl, err = url.Parse("http://" + client.host)

	if err != nil {
		fmt.Printf("[polarisctl internal sys err] build url failed:%v\n", err)
		os.Exit(1)
	}
	client.murl.Path += value
	return client.murl
}

// Put 更新/修改操作
func (client *ApiClient) Put(body io.Reader) []byte {
	client.buildReq("PUT", client.murl.String(), body)
	client.req.Header.Add("Content-Type", "application/json")
	return client.do()
}

// Post 创建
func (client *ApiClient) Post(body io.Reader) []byte {
	client.buildReq("POST", client.murl.String(), body)
	client.req.Header.Add("Content-Type", "application/json")
	return client.do()
}

// Get 查询
func (client *ApiClient) Get() []byte {
	client.buildReq("GET", client.murl.String(), nil)
	return client.do()
}

// do 执行 http do ，处理 http code
func (client *ApiClient) do() []byte {
	debuglog := viper.GetBool("debug")
	defer func() {
		// clear req
		client.req = nil
		client.murl = nil
	}()

	method := client.req.Method

	if debuglog {
		fmt.Printf("[polarisctl debug] req:%+v\n", client.req)
	}
	// send
	res, err := client.httpClient.Do(client.req)
	if err != nil {
		fmt.Printf("[polarisctl internal sys err] http %s failed:%v\n", method, err)
		os.Exit(1)
	}
	defer res.Body.Close()

	if !client.statusCheck(res.StatusCode) {
		fmt.Printf("[polarisctl internal sys err] http %s failed:%s\n", method, res.Status)
	}

	// parse body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("[polarisctl internal sys err] http %s failed:%v\n", method, err)
		os.Exit(1)
	}
	return body
}

// buildReq 构造 req 设置 token
func (client *ApiClient) buildReq(method string, url string, body io.Reader) {
	// build req
	var err error
	client.req, err = http.NewRequest(method, url, body)
	if err != nil {
		fmt.Printf("[polarisctl internal sys err] http build Req failed:%v\n", err)
		os.Exit(1)
	}
	client.req.Header.Add("X-Polaris-Token", client.token)
}

// statusCheck 检查 http code
func (client *ApiClient) statusCheck(code int) bool {
	if code == 415 {
		return false
	}
	if code >= 200 && code <= 299 {
		return true
	}
	if code >= 400 && code <= 499 {
		return true
	}
	if code == 500 {
		return true
	}
	return false
}
