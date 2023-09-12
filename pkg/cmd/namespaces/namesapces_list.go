package namespaces

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"
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

	var result entity.Result
	err := json.Unmarshal(body, &result)
	if err != nil {
		fmt.Printf("[polarisctl internal err]: unmarshal body failed:%v body:%s\n", err, string(body))
		return
	}

	result.Dump()
	fmt.Printf("\n====================== responses ========================\n")
	entity.PrintNamespaces(result.Namespaces)

}
