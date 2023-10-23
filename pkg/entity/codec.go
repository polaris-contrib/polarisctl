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
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type TableCodec interface {
	Codec(interface{}) []byte
	PassFilter(bool) TableCodec
}

type ProtoMessageCodec struct {
	rsName     string
	passFilter bool
}

type StructCodec struct {
	passFilter bool
}

func NewCodec(rsName string) TableCodec {
	return &ProtoMessageCodec{rsName, false}
}

func (codec *ProtoMessageCodec) filterTag(tagName string) bool {
	if codec.passFilter {
		return false
	}
	return !printConf.FindTag(tagName)
}

func (codec *ProtoMessageCodec) PassFilter(value bool) TableCodec {
	codec.passFilter = value
	return codec
}

func (codec *ProtoMessageCodec) Codec(resource interface{}) []byte {
	if nil == resource {
		return []byte("nil")
	}

	ret := bytes.NewBuffer([]byte{})

	resType, resVal := getReflect(resource)
	filedNum := resType.NumField()

	for i := 0; i < filedNum; i++ {

		typeField := resType.Field(i)
		field := resVal.Field(i)

		protoTag := typeField.Tag.Get("protobuf")
		if len(protoTag) == 0 {
			continue
		}

		tagName := strings.Split(typeField.Tag.Get("json"), ",")[0]
		if tagName == "" {
			fmt.Printf("[polarisctl internal sys err] cannot find json tag in %v\n", resource)
			continue
		}

		isFilter := codec.filterTag(tagName)
		if isFilter {
			continue
		}
		if field.Kind() == reflect.Ptr && (field.IsNil() || field.IsZero()) {
			continue
		}
		if field.Kind() == reflect.Slice && field.Len() == 0 && !isFilter {
			continue
		}
		if field.Kind() == reflect.Map && field.Len() == 0 && !isFilter {
			continue
		}

		ret.WriteString(tagName)
		ret.WriteByte(':')
		switch field.Kind() {
		case reflect.Ptr:
			if field.Type().String() == "*anypb.Any" {
				var err error
				msgName := ""
				var res proto.Message
				if err, msgName, res = unmarshalAny(field.Interface().(*anypb.Any)); err != nil {
					continue
				}
				ret.Write(NewCodec(msgName).PassFilter(true).Codec(res))

				continue
			}
			if strings.Contains(field.Type().String(), "wrapperspb") {
				ret.WriteString(fmt.Sprintf("%v", field.Elem().FieldByName("Value").Interface()))
			} else {
				ret.WriteByte('{')
				ret.Write((NewCodec("default")).PassFilter(true).Codec(field.Interface()))
				ret.WriteByte('}')
			}
		case reflect.Slice:
			ret.WriteByte('[')
			for i := 0; i < field.Len(); i++ {
				if strings.Contains(field.Index(i).Type().String(), "wrapperspb") {
					ret.WriteString(fmt.Sprintf("%v", field.Index(i).Elem().FieldByName("Value").Interface()))
				} else {
					ret.Write((NewCodec("default")).PassFilter(true).Codec(field.Index(i).Interface()))
				}
				if i != field.Len()-1 {
					ret.WriteByte(',')
				}
			}
			ret.WriteByte(']')
		case reflect.Map:
			ret.WriteByte('{')
			ret.WriteString(fmt.Sprintf("%v", field.Interface()))
			ret.WriteByte('}')
		case reflect.String, reflect.Bool, reflect.Int, reflect.Int32, reflect.Int64, reflect.Uint32, reflect.Uint64:
			ret.WriteString(fmt.Sprintf("%v", field.Interface()))
		case reflect.Struct:
			ret.WriteByte('{')
			ret.Write((NewCodec("default")).PassFilter(true).Codec(field.Interface()))
			ret.WriteByte('}')
		default:
			ret.WriteString("<unkown type>")
		}
		ret.WriteByte(',')
	}
	data := ret.Bytes()
	if 0 == len(data) {
		return data
	}
	return data[0 : len(data)-1]
}

func (codec *StructCodec) filterTag(tagName string) bool {
	if codec.passFilter {
		return false
	}
	return !printConf.FindTag(tagName)
}

func (codec *StructCodec) PassFilter(value bool) TableCodec {
	codec.passFilter = value
	return codec
}

func (codec *StructCodec) Codec(resource interface{}) []byte {
	if nil == resource {
		return []byte("nil")
	}

	ret := bytes.NewBuffer([]byte{})

	resType, resVal := getReflect(resource)
	filedNum := resType.NumField()

	for i := 0; i < filedNum; i++ {

		typeField := resType.Field(i)
		field := resVal.Field(i)

		tagName := strings.Split(typeField.Tag.Get("json"), ",")[0]
		if tagName == "" {
			fmt.Printf("[polarisctl internal sys err] cannot find json tag in %v\n", resource)
			continue
		}

		isFilter := codec.filterTag(tagName)
		if isFilter {
			continue
		}
		if field.Kind() == reflect.Ptr && (field.IsNil() || field.IsZero()) {
			continue
		}
		if field.Kind() == reflect.Slice && field.Len() == 0 && !isFilter {
			continue
		}
		if field.Kind() == reflect.Map && field.Len() == 0 && !isFilter {
			continue
		}

		ret.WriteString(tagName)
		ret.WriteByte(':')
		ret.WriteString(fmt.Sprintf("%+v", field.Interface()))
		ret.WriteByte(',')
	}
	data := ret.Bytes()
	if 0 == len(data) {
		return data
	}
	return data[0 : len(data)-1]
}

func codeToStr(value uint32) string {
	return strconv.FormatUint(uint64(value), 10)
}

func getAllFieldNames(response interface{}) map[string]string {

	rets := map[string]string{}

	resType := reflect.TypeOf(response)
	resVal := reflect.ValueOf(response)
	if resType.Kind() == reflect.Ptr {
		resType = resType.Elem()
		resVal = resVal.Elem()
	}
	filedNum := resType.NumField()

	for i := 0; i < filedNum; i++ {
		typeField := resType.Field(i)
		protoTag := typeField.Tag.Get("protobuf")
		if len(protoTag) == 0 {
			continue
		}
		tagName := strings.Split(typeField.Tag.Get("json"), ",")[0]
		if tagName == "" {
			fmt.Printf("[polarisctl internal sys err] cannot find tag %s \n", tagName)
			continue
		}

		fieldVal := resVal.Field(i)

		if fieldVal.Kind() == reflect.Slice {
			elemType := fieldVal.Type().Elem()   // 获取 Slice 定义的元素类型
			elem := reflect.New(elemType).Elem() // 构造一个该元素类型的空值
			fieldVal = elem
		}
		msgName := proto.MessageName(fieldVal.Interface().(proto.Message))

		rets[msgName] = resType.Field(i).Name
	}

	return rets
}

func unmarshalAny(any *anypb.Any) (error, string, proto.Message) {
	urlName := any.TypeUrl
	lastIndex := strings.LastIndex(urlName, "/")
	if lastIndex != -1 {
		urlName = urlName[lastIndex+1:]
	}
	resourceType := proto.MessageType(urlName)
	if nil == resourceType {
		fmt.Printf("[polarisctl internal sys err]Error any element MessageType:%s cannot find \n", resourceType)
		return errors.New("MessageType:" + urlName + " cannot find in proto"), "", nil
	}
	resource := reflect.New(resourceType.Elem()).Interface().(proto.Message)

	if err := proto.Unmarshal(any.Value, resource); err != nil {
		fmt.Printf("[polarisctl internal sys err]Error unmarshal any element  %v\n", err)
		return err, "", nil
	}
	return nil, urlName, resource
}

func getReflect(value interface{}) (reflect.Type, reflect.Value) {

	resType := reflect.TypeOf(value)
	resVal := reflect.ValueOf(value)
	if resType.Kind() == reflect.Ptr {
		resType = resType.Elem()
		resVal = resVal.Elem()
	}
	return resType, resVal

}
