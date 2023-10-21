package entity

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"github.com/polarismesh/specification/source/go/api/v1/config_manage"
	"github.com/polarismesh/specification/source/go/api/v1/service_manage"
)

var resourceCache map[string]interface{}

func init() {
	resourceCache = map[string]interface{}{
		"v1.BatchQueryResponse":       &service_manage.BatchQueryResponse{},
		"v1.BatchWriteResponse":       &service_manage.BatchWriteResponse{},
		"v1.Response":                 &service_manage.Response{},
		"v1.ConfigBatchQueryResponse": &config_manage.ConfigBatchQueryResponse{},
		"v1.ConfigBatchWriteResponse": &config_manage.ConfigBatchWriteResponse{},
		"v1.ConfigResponse":           &config_manage.ConfigResponse{},
		"Maintain.LogLevel":           &[]LogLevelResponse{},
		"Maintain.CMDB":               &[]CMDBResponse{},
		"Maintain.Leader":             &[]LeadersResponse{},
		"Maintain.SDKClient":          &SDKClient{},
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
	target, ok := resourceCache[parse.responseKind]
	if !ok {
		fmt.Printf("[polarisctl internal sys err]: unknown key: %s\n", parse.responseKind)
		return nil
	}

	if strings.HasPrefix(parse.responseKind, "Maintain.") {
		err := json.Unmarshal(data, target)
		if err != nil {
			fmt.Printf("[polarisctl internal err] unmarshal data failed:%v data:%s\n", err, string(data))
			return nil
		}
		return target
	} else {
		instance := target.(proto.Message)
		err := jsonpb.Unmarshal(bytes.NewReader(data), instance)
		if err != nil {
			fmt.Printf("[polarisctl internal err] unmarshal data failed:%v data:%s\n", err, string(data))
			return nil
		}
		return instance
	}
	return nil
}

// LogLevelResponse
type LogLevelResponse struct {
	Name  string `json:"Name"`
	Level string `json:"Level"`
}

// LeadersResponse
type LeadersResponse struct {
	ElectKey string `json:"ElectKey"`
	Host     string `json:"Host"`
}

// CMDBResponse
type CMDBResponse struct {
	Campus   string `json:"Campus"`
	CampusID string `json:"CampusID"`
	IP       string `json:"IP"`
	Region   string `json:"Region"`
	RegionID string `json:"RegionID"`
	Zone     string `json:"Zone"`
	ZoneID   string `json:"ZoneID"`
}

// ClientsResponse
type ClientsResponse struct {
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
