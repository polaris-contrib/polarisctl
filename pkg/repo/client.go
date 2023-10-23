/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */
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
const namingV1Api string = "/naming/v1/"
const namingV2Api string = "/naming/v2/"

const confV1Api string = "/config/v1/"
const coreV1API string = "/core/v1/"
const maintainV1API string = "/maintain/v1/"
const (
	//KNamespaceUrl namespaces 操作的 url 前缀
	API_NAMESPACES     ResourceAPI = namingV1Api + "namespaces"
	API_NAMESPACES_DEL ResourceAPI = namingV1Api + "namespaces/delete"

	// service
	API_SERVICES     ResourceAPI = namingV1Api + "services"
	API_SERVICES_ALL ResourceAPI = namingV1Api + "services/all"
	API_SERVICES_DEL ResourceAPI = namingV1Api + "services/delete"

	// service alias
	API_ALIAS      ResourceAPI = namingV1Api + "service/alias"
	API_ALIAS_LIST ResourceAPI = namingV1Api + "service/aliases"
	API_ALIAS_DEL  ResourceAPI = namingV1Api + "service/aliases/delete"

	// instances
	API_INSTANCES              ResourceAPI = namingV1Api + "instances"
	API_INSTANCES_DEL          ResourceAPI = namingV1Api + "instances/delete"
	API_INSTANCES_COUNT        ResourceAPI = namingV1Api + "instances/count"
	API_INSTANCES_LABELS       ResourceAPI = namingV1Api + "instances/labels"
	API_INSTANCES_HOST_DEL     ResourceAPI = namingV1Api + "instances/delete/host"
	API_INSTANCES_HOST_ISOLATE ResourceAPI = namingV1Api + "instances/isolate/host"

	// routings
	API_ROUTINGS        ResourceAPI = namingV2Api + "routings"
	API_ROUTINGS_DEL    ResourceAPI = namingV2Api + "routings/delete"
	API_ROUTINGS_ENABLE ResourceAPI = namingV2Api + "routings/enable"

	// circuitbreaker
	API_CIRCUITBREAKER        ResourceAPI = namingV1Api + "circuitbreaker/rules"
	API_CIRCUITBREAKER_DEL    ResourceAPI = namingV1Api + "circuitbreakers/delete"
	API_CIRCUITBREAKER_ENABLE ResourceAPI = namingV1Api + "circuitbreaker/rules/enable"

	// ratelimits
	API_RATELIMITS        ResourceAPI = namingV1Api + "ratelimits"
	API_RATELIMITS_DEL    ResourceAPI = namingV1Api + "ratelimits/delete"
	API_RATELIMITS_ENABLE ResourceAPI = namingV1Api + "ratelimits/enable"

	// faultdetectors
	API_FAULTDETECTORS     ResourceAPI = namingV1Api + "faultdetectors"
	API_FAULTDETECTORS_DEL ResourceAPI = namingV1Api + "faultdetectors/delete"

	// contracts
	API_CONTRACTS               ResourceAPI = namingV1Api + "service/contracts"
	API_CONTRACTS_DEL           ResourceAPI = namingV1Api + "service/contracts/delete"
	API_CONTRACTS_METHOD_DEL    ResourceAPI = namingV1Api + "service/contract/method/delete"
	API_CONTRACTS_METHOD_APPEND ResourceAPI = namingV1Api + "service/contract/methods/append"
	API_CONTRACTS_METHOD        ResourceAPI = namingV1Api + "service/contract/methods"

	// configfiles
	API_CONFIGFILES         ResourceAPI = confV1Api + "configfiles"
	API_CONFIGFILES_DEL     ResourceAPI = confV1Api + "configfiles/batchdelete"
	API_CONFIGFILES_EXPORT  ResourceAPI = confV1Api + "configfiles/export"
	API_CONFIGFILES_IMPORT  ResourceAPI = confV1Api + "configfiles/import"
	API_CONFIGFILES_BYGROUP ResourceAPI = confV1Api + "configfiles/by-group"
	API_CONFIGFILES_SEARCH  ResourceAPI = confV1Api + "configfiles/search"
	API_CONFIGFILES_PUB     ResourceAPI = confV1Api + "configfiles/createandpub"

	// configgroup
	API_CONFIGGROUPS     ResourceAPI = confV1Api + "configfilegroups"
	API_CONFIGGROUPS_DEL ResourceAPI = confV1Api + "configfilegroups"

	// config release
	API_RELEASE      ResourceAPI = confV1Api + "configfiles/release"
	API_RELEASES     ResourceAPI = confV1Api + "configfiles/releases"
	API_RELEASE_VER  ResourceAPI = confV1Api + "configfiles/release/versions"
	API_RELEASE_ROLL ResourceAPI = confV1Api + "configfiles/releases/rollback"
	API_RELEASE_DEL  ResourceAPI = confV1Api + "configfiles/releases/delete"
	API_RELEASE_HIST ResourceAPI = confV1Api + "configfiles/releasehistory"

	// maintain
	API_MAINTAIN_LOG    ResourceAPI = maintainV1API + "log/outputlevel"
	API_MAINTAIN_LEADER ResourceAPI = maintainV1API + "leaders"
	API_MAINTAIN_CMDB   ResourceAPI = maintainV1API + "cmdb/info"
	API_MAINTAIN_CLIENT ResourceAPI = maintainV1API + "report/clients"
)

const (
	RS_NAMESPACES     ResourceName = "namespace"
	RS_SERVICES       ResourceName = "service"
	RS_ALIAS          ResourceName = "alias"
	RS_INSTANCES      ResourceName = "instances"
	RS_ROUTINGS       ResourceName = "routings"
	RS_CIRCUITBREAKER ResourceName = "circuitbreaker"
	RS_RATELIMITS     ResourceName = "ratelimits"
	RS_FAULTDETECTORS ResourceName = "faultdetectors"
	RS_CONTRACTS      ResourceName = "contracts"
	RS_CONFIGGROUP    ResourceName = "configgroups"
	RS_CONFIGFILES    ResourceName = "configfiles"
	RS_RELEASE        ResourceName = "release"
	RS_MAINTAIN       ResourceName = "maintain"
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
func (client *ApiClient) Delete(body io.Reader) (int, []byte) {
	client.buildURL()
	client.buildReq("DELETE", client.murl.String(), body)
	client.req.Header.Add("Content-Type", "application/json")
	return client.do()
}

// Put 更新/修改操作
func (client *ApiClient) Put(body io.Reader) (int, []byte) {
	client.buildURL()
	client.buildReq("PUT", client.murl.String(), body)
	client.req.Header.Add("Content-Type", "application/json")
	return client.do()
}

// Post 创建
func (client *ApiClient) Post(body io.Reader) (int, []byte) {
	client.buildURL()
	client.buildReq("POST", client.murl.String(), body)
	client.req.Header.Add("Content-Type", "application/json")
	return client.do()
}

// Get 查询
func (client *ApiClient) Get() (int, []byte) {
	client.buildURL()
	client.buildReq("GET", client.murl.String(), nil)
	return client.do()
}

// do 执行 http do ，处理 http code
func (client *ApiClient) do() (int, []byte) {
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
		fmt.Printf("[polarisctl internal sys err] http do %s failed:%v\n", method, err)
		os.Exit(1)
	}
	defer res.Body.Close()

	// parse body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("[polarisctl internal sys err] http %s read body failed:%v\n", method, err)
		os.Exit(1)
	}

	if debuglog {
		fmt.Printf("[polarisctl debug] resp :%+v\n", res)
		fmt.Printf("[polarisctl debug] body :%s\n", string(body))
	}
	return res.StatusCode, body
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
