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
	resource string
	client   *ApiClient
	method   string
	rsFile   string
	ctlPrint *entity.PolarisPrint
	batch    bool
}

// NewResourceRepo 查询操作
func NewResourceRepo(resource, url string) *ResourceRepo {
	return &ResourceRepo{
		resource: resource,
		client:   NewApiClient(url),
		ctlPrint: entity.NewPolarisPrint(),
		batch:    true,
	}
}

// Batch set batch write
func (rsRepo *ResourceRepo) Batch(value bool) *ResourceRepo {
	rsRepo.batch = value
	return rsRepo
}

// Print set print
func (rsRepo *ResourceRepo) Print(ctlPrint *entity.PolarisPrint) *ResourceRepo {
	rsRepo.ctlPrint = ctlPrint
	return rsRepo
}

// Param set get url param
func (rsRepo *ResourceRepo) Param(value string) *ResourceRepo {
	rsRepo.client.queryParam = value
	return rsRepo
}

// File set put/post/del resources description file
func (rsRepo *ResourceRepo) File(value string) *ResourceRepo {
	rsRepo.rsFile = value
	return rsRepo
}

// Method set http method:GET/PUT/POST/PUT/DEL
func (rsRepo *ResourceRepo) Method(value string) *ResourceRepo {
	rsRepo.method = value
	return rsRepo
}

// Build execute
func (rsRepo ResourceRepo) Build() {
	if rsRepo.method == "GET" {
		rsRepo.get()
		return
	}
	rsRepo.write()
}

// write put/post/del resources
func (rsRepo ResourceRepo) write() {
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

	body := []byte{}

	if rsRepo.method == "POST" {
		body = rsRepo.client.Post(bytes.NewBuffer(jsonData))
	} else if rsRepo.method == "PUT" {
		body = rsRepo.client.Put(bytes.NewBuffer(jsonData))
	} else {
		fmt.Printf("[polarisctl internal sys err] resource:%s unkown method:%s\n", rsRepo.resource, rsRepo.method)
		os.Exit(1)
	}

	if rsRepo.batch {

		var response service_manage.BatchWriteResponse
		err = jsonpb.Unmarshal(bytes.NewReader(body), &response)
		if err != nil {
			fmt.Printf("[polarisctl internal err]: unmarshal body failed:%v body:%s\n", err, string(body))
			return
		}
		rsRepo.ctlPrint.Response(response).BatchWrite().Print()
		return
	}

	var response service_manage.Response
	err = jsonpb.Unmarshal(bytes.NewReader(body), &response)
	if err != nil {
		fmt.Printf("[polarisctl internal err]: unmarshal body failed:%v body:%s\n", err, string(body))
		return
	}
	rsRepo.ctlPrint.Response(response).Write().Print()
}

// get query resources
func (rsRepo ResourceRepo) get() {
	body := rsRepo.client.Get()

	if rsRepo.batch {

		var response service_manage.BatchQueryResponse
		err := jsonpb.Unmarshal(bytes.NewReader(body), &response)
		if err != nil {
			fmt.Printf("[polarisctl internal err]: unmarshal body failed:%v body:%s\n", err, string(body))
			return
		}
		rsRepo.ctlPrint.Response(response).BatchQuery().Print()
		return
	}
	var response service_manage.Response
	err := jsonpb.Unmarshal(bytes.NewReader(body), &response)
	if err != nil {
		fmt.Printf("[polarisctl internal err]: unmarshal body failed:%v body:%s\n", err, string(body))
		return
	}
	rsRepo.ctlPrint.Response(response).Query().Print()
}
