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
package entity

import (
	"errors"
	"strings"
)

var printConf *PrintConf

func init() {
	printConf = NewPrintConf()
}

// PolarisClusterConf polaris 集群的控制面配置
type PolarisClusterConf struct {
	Name  string `json:"name"`  // 集群名
	Host  string `json:"host"`  // 集群控制面域名，或者"ip:port"
	Token string `json:"token"` // 集群控制面的鉴权 token
}

type PrintConf struct {
	ResourceUrl     map[string]bool
	ResourceTagName map[string]bool
}

type PrintOption func(printConf *PrintConf)

func WithTags(tags string) PrintOption {
	return func(conf *PrintConf) {
		if 0 == len(tags) {
			return
		}
		for _, tag := range strings.Split(tags, ",") {
			conf.ResourceTagName[tag] = true
		}
	}
}

func (conf *PrintConf) FindTag(name string) bool {
	// 未配置，默认全部输出
	if 0 == len(conf.ResourceTagName) {
		return true
	}
	ok, isPrint := conf.ResourceTagName[name]
	if !ok || !isPrint {
		return false
	}
	return true
}

func (conf *PrintConf) Find(url string) (bool, bool) {
	value, ok := conf.ResourceUrl[url]
	return value, ok

}

func NewPrintConf() *PrintConf {
	return &PrintConf{
		ResourceTagName: map[string]bool{},
		ResourceUrl: map[string]bool{
			// base message
			"google.protobuf.UInt32Value": false,
			"google.protobuf.StringValue": false,

			// any
			"google.protobuf.Any": true,

			// v1
			"v1.Response":                 true,
			"v1.BatchQueryResponse":       true,
			"v1.BatchWriteResponse":       true,
			"v1.ConfigResponse":           true,
			"v1.ConfigBatchQueryResponse": true,
			"v1.ConfigBatchWriteResponse": true,

			// v1 resources
			"v1.Namespace":                true,
			"v1.Service":                  true,
			"v1.Instance":                 true,
			"v1.Routing":                  true,
			"v1.ServiceAlias":             true,
			"v1.Rule":                     true,
			"v1.ConfigWithService":        true,
			"v1.AuthStrategy":             true,
			"v1.Summary":                  false,
			"v1.Client":                   true,
			"v1.CircuitBreaker":           true,
			"v1.ConfigRelease":            true,
			"v1.User":                     true,
			"v1.UserGroup":                true,
			"v1.UserGroupRelation":        true,
			"v1.LoginResponse":            true,
			"v1.ModifyAuthStrategy":       true,
			"v1.ModifyUserGroup":          true,
			"v1.StrategyResources":        true,
			"v1.OptionSwitch":             true,
			"v1.InstanceLabels":           true,
			"v1.ConfigFileGroup":          true,
			"v1.ConfigFile":               true,
			"v1.ConfigFileRelease":        true,
			"v1.ConfigFileReleaseHistory": true,
			"v1.ConfigFileTemplate":       true,
			"v1.CircuitBreakerRule":       true,
			"v1.RuleRoutingConfig":        true,
			"v1.RouteRule":                true,
		},
	}
}

// PolarisCtl cli 配置
type PolarisCtlConf struct {
	DefaultCluster string               `json:"default_culster"`
	Clusters       []PolarisClusterConf `json:"clusters"`
}

// Check 检查必填配置
func (cnf PolarisCtlConf) Check() error {
	if len(cnf.DefaultCluster) == 0 {
		return errors.New("default_culster empty")
	}

	if len(cnf.Clusters) == 0 {
		return errors.New("clusters empty")
	}

	if _, err := cnf.FindCluster(cnf.DefaultCluster); err != nil {
		return err
	}

	return nil

}

// FindCluster 找到某个集群的配置
func (cnf PolarisCtlConf) FindCluster(clusterName string) (PolarisClusterConf, error) {
	for _, cluster := range cnf.Clusters {
		if cluster.Name == clusterName {
			return cluster, nil
		}
	}
	return PolarisClusterConf{}, errors.New("cannot find cluster conf:" + clusterName)
}
