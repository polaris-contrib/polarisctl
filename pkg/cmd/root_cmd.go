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
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/polaris-contrilb/polarisctl/pkg/cmd/circuitbreaker"
	"github.com/polaris-contrilb/polarisctl/pkg/cmd/config"
	"github.com/polaris-contrilb/polarisctl/pkg/cmd/faultdetectors"
	"github.com/polaris-contrilb/polarisctl/pkg/cmd/instances"
	"github.com/polaris-contrilb/polarisctl/pkg/cmd/maintain"
	"github.com/polaris-contrilb/polarisctl/pkg/cmd/namespaces"
	"github.com/polaris-contrilb/polarisctl/pkg/cmd/ratelimits"
	"github.com/polaris-contrilb/polarisctl/pkg/cmd/routings"
	"github.com/polaris-contrilb/polarisctl/pkg/cmd/services"
	"github.com/polaris-contrilb/polarisctl/pkg/entity"
	"github.com/polaris-contrilb/polarisctl/pkg/repo"
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
		fmt.Printf("[polarisctl internal sys err] UserHomeDir failed:%v\n", err)
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
			repo.RegisterCluster(cluster)
		},
	}

	defaultPath := home + "/.polarisctl/polarisctl_config.json"
	root.PersistentFlags().StringVar(&configPath, "config", defaultPath, "cluster config path")
	root.PersistentFlags().StringVar(&clusterName, "cluster", "", "current cluster")
	root.PersistentFlags().BoolVar(&debug, "debug", false, "debug log")
	viper.BindPFlag("debug", root.PersistentFlags().Lookup("debug"))

	// register namespaces
	root.AddCommand(namespaces.NewCmdNamespaces())
	root.AddCommand(services.NewCmdServices())
	root.AddCommand(instances.NewCmdInstances())
	root.AddCommand(routings.NewCmdRoutings())
	root.AddCommand(circuitbreaker.NewCmdCircuitbreaker())
	root.AddCommand(ratelimits.NewCmdRatelimits())
	root.AddCommand(faultdetectors.NewCmdFaultdetectors())
	root.AddCommand(config.NewCmdConfig())
	root.AddCommand(maintain.NewCmdMaintain())
	return root
}

// initConfig 初始化集群配置
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
		fmt.Printf("[polarisctl internal sys err] cannot find cluster:%s config\n", clusterName)
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
