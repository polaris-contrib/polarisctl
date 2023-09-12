package namespaces

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/tabwriter"

	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// fileName create/delete/update 的资源描述文件:json
var fileName string

// NewCmdNamespaces 构建 namespaces 的跟命令
func NewCmdNamespaces() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "namespaces subcommand",
		Short: "namespaces cmd",
		Long:  "namespaces [list/create/delete/update]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.AddCommand(NewCmdNamespacesCreate())
	cmd.AddCommand(NewCmdNamespacesDelete())
	cmd.AddCommand(NewCmdNamespacesUpdate())
	cmd.AddCommand(NewCmdNamespacesList())
	return cmd
}

func nsBatchOp(file string, method string, uriPath string) {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		fmt.Printf("[polarisctl err] input invalid, -f empty\n")
		os.Exit(1)
	}

	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		os.Exit(1)
	}

	client := repo.GetApiClient()
	client.BuildURL(uriPath)
	body := []byte{}

	if method == "post" {
		body = client.Post(bytes.NewBuffer(jsonData))
	} else if method == "put" {
		body = client.Put(bytes.NewBuffer(jsonData))
	} else {
		fmt.Printf("[polarisctl internal sys err] unkown method:%s\n", method)
		os.Exit(1)
	}

	result := &entity.BatchResult{}
	err = json.Unmarshal(body, result)
	if err != nil {
		fmt.Printf("[polarisctl internal err]: unmarshal body failed:%v body:%s\n", err, string(body))
		return
	}

	result.CheckRes()
	result.Dump()

	if 0 == len(result.Responses) {
		return
	}
	fmt.Printf("\n====================== responses ========================\n")

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.StripEscape|tabwriter.AlignRight|tabwriter.Debug|tabwriter.TabIndent|tabwriter.DiscardEmptyColumns|tabwriter.TabIndent)
	fmt.Fprintf(w, "namespace\tcode\tinfo\t\n")
	for _, res := range result.Responses {
		fmt.Fprintf(w, "%s\t%d\t%s\t\n", res.Namespace.Name, res.Code, res.Info)
	}
	w.Flush()
}
