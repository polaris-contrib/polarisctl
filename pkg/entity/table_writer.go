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
	"fmt"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/liushuochen/gotable"
	"github.com/liushuochen/gotable/table"
	"github.com/polarismesh/specification/source/go/api/v1/config_manage"
	"github.com/polarismesh/specification/source/go/api/v1/model"
	"github.com/polarismesh/specification/source/go/api/v1/service_manage"
	"google.golang.org/protobuf/types/known/anypb"
)

type Writer interface {
	Write(interface{})
}

// TableWriter table write
type TableWriter struct {
	table *table.Table
}

// NewTableWriter
func NewTableWriter(options ...PrintOption) Writer {

	for _, option := range options {
		option(printConf)
	}
	return &TableWriter{}
}

// print
func (p *TableWriter) Write(result interface{}) {
	switch result.(type) {
	case *service_manage.Response, *config_manage.ConfigResponse:
		p.table, _ = gotable.Create("type", "resource", "code", "info")
		p.response(result)
		fmt.Println(p.table)
	case *service_manage.BatchWriteResponse, *config_manage.ConfigBatchWriteResponse:
		p.batchWriteResponse(result)
	case *service_manage.BatchQueryResponse, *config_manage.ConfigBatchQueryResponse:
		p.batchQueryResponse(result)
	case *HttpFailed:
		p.httpFailed(result)
	default:
		p.defaultResp(result)
	}
}
func (p *TableWriter) httpFailed(result interface{}) {
	res := result.(*HttpFailed)
	p.table, _ = gotable.Create("http code", "http body")
	p.table.AddRow([]string{res.Code, res.Body})
	fmt.Println(p.table)
}

func (p *TableWriter) defaultResp(response interface{}) {

	resType, resValue := getReflect(response)
	p.table, _ = gotable.Create("type", "resource")

	if resType.Kind() == reflect.Slice {
		elemType := resValue.Type().Elem() // 获取 Slice 定义的元素类型
		name := strings.Split(elemType.String(), ".")[1]
		codec := &StructCodec{}
		if 0 == resValue.Len() {
			p.table.AddRow([]string{name, "nil"})
		}
		for i := 0; i < resValue.Len(); i++ {
			res := resValue.Index(i)
			p.table.AddRow([]string{name, string(codec.Codec(res.Interface()))})
		}
		fmt.Println(p.table)
		return
	}

	name := strings.Split(resType.String(), ".")[1]
	codec := &StructCodec{}
	p.table.AddRow([]string{name, string(codec.Codec(response))})

	fmt.Println(p.table)

}

func (p *TableWriter) response(response interface{}) {

	findRes := false
	if response == nil {
		p.table.AddRow([]string{"unkown", "nil", "0", "0"})
		return
	}

	resValue := reflect.ValueOf(response)
	if resValue.Kind() == reflect.Ptr {
		resValue = resValue.Elem()
	}

	if resValue.IsZero() {
		p.table.AddRow([]string{"unkown", "zero", "0", "0"})
		return
	}
	code := fmt.Sprintf("%v", resValue.FieldByName("Code").Elem().FieldByName("Value").Interface())
	info := fmt.Sprintf("%v", resValue.FieldByName("Info").Elem().FieldByName("Value").Interface())
	for rsName, fieldName := range getAllFieldNames(response) {

		if value, ok := printConf.Find(rsName); !ok || !value {
			continue
		}

		fieldValue := resValue.FieldByName(fieldName)

		if fieldValue.IsNil() {
			continue
		}

		if rsName == "google.protobuf.Any" {
			ret := p.anyRes(fieldValue.Interface().(*anypb.Any))
			if 0 != len(ret) {
				ret = append(ret, code)
				ret = append(ret, info)
				p.table.AddRow(ret)
				findRes = true
			}
			continue
		}

		codec := NewCodec(rsName)
		p.table.AddRow(
			[]string{rsName, string(codec.Codec(fieldValue.Interface())), code, info},
		)
		findRes = true
	}

	if !findRes {
		p.table.AddRow([]string{"nil", "nil", code, info})

	}

}

func (p *TableWriter) batchQueryResponse(response interface{}) {

	resValue := reflect.ValueOf(response)
	if resValue.Kind() == reflect.Ptr {
		resValue = resValue.Elem()
	}

	code := p.getString(resValue, "Code")
	info := p.getString(resValue, "Info")
	amount := p.getString(resValue, "Amount")
	size := p.getString(resValue, "Size")
	total := p.getString(resValue, "Total")

	if size == "nil" && total != "nil" {
		size = total
	}

	p.table, _ = gotable.Create("code", "info", "amount", "size")
	p.table.AddRow([]string{code, info, amount, size})
	fmt.Println(p.table)

	p.table, _ = gotable.Create("type", "resource")
	p.batchQueryResp(getAllFieldNames(response), resValue)
	fmt.Println(p.table)
}

func (p *TableWriter) batchWriteResponse(response interface{}) {

	resValue := reflect.ValueOf(response).Elem()
	code := p.getString(resValue, "Code")
	info := p.getString(resValue, "Info")
	size := p.getString(resValue, "Size")
	total := p.getString(resValue, "Total")

	succ, failed := uint32(0), uint32(0)
	responses := resValue.FieldByName("Responses")
	if responses.IsValid() {
		for i := 0; i < responses.Len(); i++ {
			res := responses.Index(i).Elem()
			codeVal := res.FieldByName("Code").Elem()
			if !codeVal.IsValid() {
				continue
			}
			if codeVal.FieldByName("Value").Interface().(uint32) == uint32(model.Code_ExecuteSuccess) {
				succ += 1
			} else {
				failed += 1
			}
		}
	}

	if size == "nil" && total != "nil" {
		size = total
	}

	p.table, _ = gotable.Create("code", "info", "size", "succ", "failed")
	p.table.AddRow([]string{
		code,
		info,
		size,
		codeToStr(succ),
		codeToStr(failed),
	})
	fmt.Println(p.table)

	p.table, _ = gotable.Create("type", "resource", "code", "info")

	if responses.IsValid() {
		for i := 0; i < responses.Len(); i++ {
			res := responses.Index(i).Elem().Interface()
			p.response(res)
		}
	}
	fmt.Println(p.table)

}

func (p *TableWriter) anyRes(data *anypb.Any) []string {
	ret := []string{}
	if data == nil {
		return ret
	}

	var err error
	msgName := ""
	var res proto.Message

	if err, msgName, res = unmarshalAny(data); err != nil {
		return ret
	}

	value, ok := printConf.Find(msgName)
	if !ok {
		return []string{"unkown type", "unkown resource"}
	}
	if !value {
		return ret
	}

	codec := NewCodec(msgName)
	ret = append(ret, msgName)
	ret = append(ret, string(codec.Codec(res)))
	return ret

}

func (p *TableWriter) batchQueryResp(fieldNames map[string]string, response reflect.Value) {

	for rsName, fieldName := range fieldNames {
		// 跳过未配置的 message
		// 跳过配置不输出的 message
		if value, ok := printConf.Find(rsName); !ok || !value {
			continue
		}

		fieldValue := response.FieldByName(fieldName)

		if fieldValue.IsNil() {
			continue
		}

		if fieldValue.Type().Kind() != reflect.Slice {
			continue
		}

		if fieldValue.Len() == 0 {
			continue
		}

		for i := 0; i < fieldValue.Len(); i++ {

			if rsName == "google.protobuf.Any" {
				ret := p.anyRes(fieldValue.Index(i).Interface().(*anypb.Any))
				if 0 != len(ret) {
					p.table.AddRow(ret)
				}
				continue
			}
			codec := NewCodec(rsName)
			p.table.AddRow([]string{rsName, string(codec.Codec(fieldValue.Index(i).Interface()))})
		}
	}
}

func (p *TableWriter) getString(value reflect.Value, name string) string {
	fieldValue := value.FieldByName(name)
	if !fieldValue.IsValid() {
		return "nil"
	}
	if fieldValue.IsZero() {
		return "nil"
	}
	return fmt.Sprintf("%v", fieldValue.Elem().FieldByName("Value").Interface())
}
