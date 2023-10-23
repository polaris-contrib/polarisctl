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
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/polarismesh/specification/source/go/api/v1/config_manage"
	"github.com/polarismesh/specification/source/go/api/v1/service_manage"
)

var resourceBuilder map[string]func() interface{}

func init() {
	resourceBuilder = map[string]func() interface{}{

		"v1.BatchQueryResponse":       func() interface{} { return &service_manage.BatchQueryResponse{} },
		"v1.BatchWriteResponse":       func() interface{} { return &service_manage.BatchWriteResponse{} },
		"v1.Response":                 func() interface{} { return &service_manage.Response{} },
		"v1.ConfigBatchQueryResponse": func() interface{} { return &config_manage.ConfigBatchQueryResponse{} },
		"v1.ConfigBatchWriteResponse": func() interface{} { return &config_manage.ConfigBatchWriteResponse{} },
		"v1.ConfigResponse":           func() interface{} { return &config_manage.ConfigResponse{} },
		"Maintain.CMDB":               func() interface{} { return &[]CMDB{} },
		"Maintain.LogLevel":           func() interface{} { return &[]LogLevel{} },
		"Maintain.Leader":             func() interface{} { return &[]Leaders{} },
		"Maintain.Clients":            func() interface{} { return &Clients{} },
	}
}

// ResponseParse parse json body
type ResponseParse struct {
	responseKind string
}

// NewResponseParse build
func NewResponseParse(responseKind string) *ResponseParse {
	return &ResponseParse{
		responseKind: responseKind,
	}
}

// ResponseKind se response kind
func (parse *ResponseParse) ResponseKind(kind string) *ResponseParse {
	parse.responseKind = kind
	return parse
}

// Parse unmarshal json body to pb message /strcut
func (parse *ResponseParse) Parse(data []byte) interface{} {
	build, ok := resourceBuilder[parse.responseKind]
	if !ok {
		fmt.Printf("[polarisctl internal sys err]: unknown key: %s\n", parse.responseKind)
		os.Exit(1)
	}

	instance := build()
	if strings.HasPrefix(parse.responseKind, "Maintain.") {
		err := json.Unmarshal(data, instance)
		if err != nil {
			fmt.Printf("[polarisctl internal err] unmarshal maitain data failed:%v data:%s\n", err, string(data))
			os.Exit(1)
		}
	} else {
		target := instance.(proto.Message)
		err := jsonpb.Unmarshal(bytes.NewReader(data), target)
		if err != nil {
			fmt.Printf("[polarisctl internal err] unmarshal data failed:%v data:%s\n", err, string(data))
			os.Exit(1)
		}
		instance = target
	}
	return instance
}

// LogLevel
type LogLevel struct {
	Name  string `json:"Name"`
	Level string `json:"Level"`
}

// Leaders
type Leaders struct {
	ElectKey   string `json:"ElectKey"`
	Host       string `json:"Host"`
	Ctime      int64  `json:"Ctime"`
	CreateTime string `json:"CreateTime"`
	Mtime      int64  `json:"Mtime"`
	ModifyTime string `json:"ModifyTime"`
	Valid      bool   `json:"Valid"`
}

// CMDB
type CMDB struct {
	Campus   string `json:"Campus"`
	CampusID string `json:"CampusID"`
	IP       string `json:"IP"`
	Region   string `json:"Region"`
	RegionID string `json:"RegionID"`
	Zone     string `json:"Zone"`
	ZoneID   string `json:"ZoneID"`
}

// Clients
type Clients struct {
	// 响应码
	Code int64 `json:"Code"`
	// 客户端列表
	Clients []SDKClient `json:"Response"`
}

// SDKClient
type SDKClient struct {
	Labels  map[string]interface{} `json:"labels"`
	Targets []string               `json:"targets"`
}

// HttpFailed http failed resp
type HttpFailed struct {
	Code string `json:"http code"`
	Body string `json:"http body"`
}
