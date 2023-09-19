package repo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/golang/protobuf/jsonpb"
	"github.com/polarismesh/specification/source/go/api/v1/service_manage"
)

// ResourceWrite 创建/删除/修改 资源
type ResourceRepo struct {
	resource    string
	queryParam  entity.QueryParam
	resourceApi ResourceAPI
	client      *ApiClient
	rsFile      string
	method      string
}

// NewResourceWriteRepo 创建/删除/修改
func NewResourceWriteRepo(resource, url, method, rsFile string) *ResourceRepo {
	return &ResourceRepo{
		resource:    resource,
		rsFile:      rsFile,
		resourceApi: url,
		method:      method,
		client:      GetApiClient(),
	}
}

// NewResourceListRepo 查询操作
func NewResourceListRepo(resource, url string, param entity.QueryParam) *ResourceRepo {
	return &ResourceRepo{
		resource:    resource,
		queryParam:  param,
		resourceApi: url,
		method:      "GET",
		client:      GetApiClient(),
	}
}

func (rsRepo ResourceRepo) Write() {
	jsonFile, err := os.Open(rsRepo.rsFile)
	if err != nil {
		fmt.Printf("[polarisctl err] input invalid, -f empty\n")
		os.Exit(1)
	}

	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		os.Exit(1)
	}

	rsRepo.client.BuildURL(rsRepo.resourceApi, "")
	body := []byte{}

	if rsRepo.method == "POST" {
		body = rsRepo.client.Post(bytes.NewBuffer(jsonData))
	} else if rsRepo.method == "PUT" {
		body = rsRepo.client.Put(bytes.NewBuffer(jsonData))
	} else {
		fmt.Printf("[polarisctl internal sys err] resource:%s unkown method:%s\n", rsRepo.resource, rsRepo.method)
		os.Exit(1)
	}

	var response service_manage.BatchWriteResponse
	err = jsonpb.Unmarshal(bytes.NewReader(body), &response)
	if err != nil {
		fmt.Printf("[polarisctl internal err]: unmarshal body failed:%v body:%s\n", err, string(body))
		return
	}
	ctlPrint := entity.NewPolarisPrint(response)
	ctlPrint.Print()
}

func (rsRepo ResourceRepo) Get() {
	rsRepo.client.BuildURL(rsRepo.resourceApi, rsRepo.queryParam.Encode())
	body := rsRepo.client.Get()

	var response service_manage.BatchQueryResponse
	err := jsonpb.Unmarshal(bytes.NewReader(body), &response)
	if err != nil {
		fmt.Printf("[polarisctl internal err]: unmarshal body failed:%v body:%s\n", err, string(body))
		return
	}

	ctlPrint := entity.NewPolarisPrint(response)
	ctlPrint.Print()
}
