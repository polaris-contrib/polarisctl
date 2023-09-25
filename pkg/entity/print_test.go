package entity_test

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/golang/protobuf/jsonpb"
	"github.com/polarismesh/specification/source/go/api/v1/service_manage"
)

func TestPrint(t *testing.T) {

	data := `{"code": 200000,
		"info": "execute success",
		"amount": 1,
		"size": 1,
		"namespaces": [
		],
		"services": [
		{
			"name": "polarisctl_test_svr1",
			"namespace": "default",
			"metadata": {
			},
			"ports": "9092,9091",
			"business": "",
			"department": "",
			"cmdb_mod1": "",
		"cmdb_mod2": "",
	 "cmdb_mod3": "",
	 "comment": "test polarisctl namespaces cmd",
   "owners": "100020293116",
   "token": null,
   "ctime": "2023-09-20 17:01:08",
   "mtime": "2023-09-20 17:07:13",
   "revision": "72edcb4a4097494cb44846510d1098d3",
   "platform_id": "",
   "total_instance_count": 1,
   "healthy_instance_count": 0,
   "user_ids": ["123","124"]}]}`
	response := service_manage.BatchQueryResponse{}
	err := jsonpb.Unmarshal(bytes.NewReader([]byte(data)), &response)
	if err != nil {
		fmt.Printf("unmarshal failed:%v\n", err)
		os.Exit(1)
	}
	fmt.Printf("res:%v\n", response)
	print := entity.NewPolarisPrint().BatchQuery().Response(response)
	print.Print()
}
