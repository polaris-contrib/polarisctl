package namespaces

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/golang/protobuf/jsonpb"
	"github.com/polarismesh/specification/source/go/api/v1/service_manage"
	"github.com/spf13/cobra"
)

var name string
var offset int
var limit int

// NewCmdNamespacesList 构建获取 namespace 列表命令
func NewCmdNamespacesList() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list naemspaces",
		Long:  "list namespaces",
		Run: func(cmd *cobra.Command, args []string) {
			runGetNs()
		},
	}
	cmd.Flags().StringVarP(&name, "name", "n", "", "naemspaces for get")
	cmd.Flags().IntVarP(&offset, "offset", "o", 0, "get offset")
	cmd.Flags().IntVarP(&limit, "limit", "l", 10, "get limit")
	return cmd
}

func runGetNs() {

	client := repo.GetApiClient()
	u := client.BuildURL(repo.NamespaceURL)

	query := u.Query()
	if len(name) != 0 {
		// 遍历所有的 namespaces
		query.Add("name", name)
	}
	query.Add("offset", strconv.Itoa(offset))
	query.Add("limit", strconv.Itoa(limit))
	u.RawQuery = query.Encode()

	body := client.Get()

	var response service_manage.BatchQueryResponse
	//var result entity.Result
	//err := json.Unmarshal(body, &response)
	err := jsonpb.Unmarshal(bytes.NewReader(body), &response)
	if err != nil {
		fmt.Printf("[polarisctl internal err]: unmarshal body failed:%v body:%s\n", err, string(body))
		return
	}

	ctlPrint := entity.NewPolarisPrint(response)
	ctlPrint.Print()
	//result.Dump()
	//fmt.Printf("\n====================== responses ========================\n")
	//entity.NamespaceDump(result.Namespaces)

}
