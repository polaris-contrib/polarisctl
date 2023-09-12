package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/0226zy/polarisctl/pkg/cmd/namespaces"
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configPath 配置路径
var configPath string

// clusterName 当前的集群名
var clusterName string

// polarisctl 配置
var polarisCtlConf entity.PolarisCtlConf

// cluster 当前集群配置
var cluster entity.PolarisClusterConf

var debug bool

// NewDefaultPolarisCommand 构建 root 命令
func NewDefaultPolarisCommand() *cobra.Command {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("[polarisctl internal sys err] UserHomeDir failed:%v\n", err)
		os.Exit(1)
	}

	root := &cobra.Command{
		Use:   "polarisctl",
		Short: "polarisctl is used to quickly initiate related OpenAPI requests",
		Long: `polarisctl polaris command line tool is used to quickly initiate related OpenAPI requests
 Find OpenApi doc at:
 https://polarismesh.cn/docs/%E5%8F%82%E8%80%83%E6%96%87%E6%A1%A3/%E6%8E%A5%E5%8F%A3%E6%96%87%E6%A1%A3/open_api`,
		Run: runHelp,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			initConfig()
			repo.InitApiClient(cluster)
		},
	}

	defaultPath := home + "/.polarisctl/polarisctl_config.json"
	root.PersistentFlags().StringVar(&configPath, "config", defaultPath, "cluster config path")
	root.PersistentFlags().StringVar(&clusterName, "cluster", "", "current cluster")
	root.PersistentFlags().BoolVar(&debug, "debug", false, "debug log")
	viper.BindPFlag("debug", root.PersistentFlags().Lookup("debug"))

	// register namespaces
	root.AddCommand(namespaces.NewCmdNamespaces())
	return root
}

func initConfig() {
	file, err := os.Open(configPath)
	if err != nil {
		fmt.Printf("[polarisctl internal sys err] open config file failed:%v\n", err)
		os.Exit(1)
	}

	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("[polarisctl internal sys err] read config file failed:%v\n", err)
		os.Exit(1)
	}

	if err = json.Unmarshal(data, &polarisCtlConf); err != nil {
		fmt.Printf("[polarisctl internal sys err] parse config failed:%v\n", err)
		os.Exit(1)
	}

	if err = polarisCtlConf.Check(); err != nil {
		fmt.Printf("[polarisctl internal sys err] conf invalid:%v\n", err)
		os.Exit(1)
	}

	if len(clusterName) == 0 {
		clusterName = polarisCtlConf.DefaultCluster
	}

	if cluster, err = polarisCtlConf.FindCluster(clusterName); err != nil {
		fmt.Printf("[polarisctl internal sys err] cannot find cluster:%d config\n", clusterName)
		os.Exit(1)
	}

	if viper.GetBool("debug") {
		fmt.Printf("[polarisctl debug] use cluster:%+v\n\n", cluster)
	}
}

// runHelp 输出帮助信息
func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
