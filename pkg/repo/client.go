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

// ResourceAPI 资源的 http url 前缀
type ResourceAPI = string
type ResourceName = string

// url_v1 polaris api v1 url
const v1Api string = "/naming/v1/"
const (
	//KNamespaceUrl namespaces 操作的 url 前缀
	API_NAMESPACES    ResourceAPI = v1Api + "namespaces"
	API_NAMESPACESDEL ResourceAPI = v1Api + "namespaces/delete"

	// service
	API_SERVICES    ResourceAPI = v1Api + "services"
	API_SERVICESALL ResourceAPI = v1Api + "services/all"
	API_SERVICESDEL ResourceAPI = v1Api + "services/delete"

	// service alias
	API_ALIAS     ResourceAPI = v1Api + "service/alias"
	API_ALIASLIST ResourceAPI = v1Api + "service/aliases"
	API_ALIASDEL  ResourceAPI = v1Api + "service/aliases/delete"
)

const (
	RS_NAMESPACES ResourceName = "namespace"
	RS_SERVICES   ResourceName = "service"
	RS_ALIAS      ResourceName = "alias"
)

// ApiClient http 请求处理
type ApiClient struct {
	host       string
	token      string
	httpClient *http.Client

	req  *http.Request
	murl *url.URL

	api        string
	queryParam string
}

var cluster entity.PolarisClusterConf

func RegisterCluster(conf entity.PolarisClusterConf) {
	cluster = conf
}

// NewApiClient build api client
func NewApiClient(api string) *ApiClient {
	apiClient = &ApiClient{
		host:       cluster.Host,
		token:      cluster.Token,
		httpClient: &http.Client{},
		api:        api,
	}
	return apiClient
}

// buildURL 构造 net.url
func (client *ApiClient) buildURL() {
	var err error
	client.murl, err = url.Parse("http://" + client.host)

	if err != nil {
		fmt.Printf("[polarisctl internal sys err] build url failed:%v\n", err)
		os.Exit(1)
	}
	client.murl.Path += client.api
	if len(client.queryParam) != 0 {
		client.murl.RawQuery = client.queryParam
	}
}

// Put 更新/修改操作
func (client *ApiClient) Put(body io.Reader) []byte {
	client.buildURL()
	client.buildReq("PUT", client.murl.String(), body)
	client.req.Header.Add("Content-Type", "application/json")
	return client.do()
}

// Post 创建
func (client *ApiClient) Post(body io.Reader) []byte {
	client.buildURL()
	client.buildReq("POST", client.murl.String(), body)
	client.req.Header.Add("Content-Type", "application/json")
	return client.do()
}

// Get 查询
func (client *ApiClient) Get() []byte {
	client.buildURL()
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
	if debuglog {
		fmt.Printf("[polarisctl debug] resp :%+v\n", res)
		fmt.Printf("[polarisctl debug] body :%s\n", string(body))
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
