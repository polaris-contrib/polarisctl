package entity

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
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
	table        *table.Table
	codecFactory *RsCodecFactory
}

// NewTableWriter
func NewTableWriter(options ...PrintOption) Writer {

	for _, option := range options {
		option(printConf)
	}
	return &TableWriter{
		codecFactory: NewRsCodeFactory(),
	}
}

// print
func (p *TableWriter) Write(result interface{}) {
	switch v := result.(type) {
	case *service_manage.Response, *config_manage.ConfigResponse:
		if nil != result {
			resValue := reflect.ValueOf(result).Elem()
			p.table, _ = gotable.Create("code", "info")
			p.table.AddRow([]string{
				fmt.Sprintf("%v", resValue.FieldByName("Code").Elem().FieldByName("Value").Interface()),
				fmt.Sprintf("%v", resValue.FieldByName("Info").Elem().FieldByName("Value").Interface()),
			},
			)
			fmt.Println(p.table)
		}

		p.table, _ = gotable.Create("type", "resource", "code", "info")
		p.response(result)
		fmt.Println(p.table)
	case *service_manage.BatchWriteResponse:
		if response, ok := result.(*service_manage.BatchWriteResponse); ok {
			p.batchWriteResponse(*response)
		} else {
			fmt.Printf("[polarisctl internal err] print failed, covert to service_manage.BatchWriteResponse failed\n")
		}
	case *service_manage.BatchQueryResponse:
		if response, ok := result.(*service_manage.BatchQueryResponse); ok {
			p.batchQueryResponse(*response)
		} else {
			fmt.Printf("[polarisctl internal err] print failed, covert to service_manage.BatchQueryResponse failed\n")
		}
	case *config_manage.ConfigBatchQueryResponse:
		if response, ok := result.(*config_manage.ConfigBatchQueryResponse); ok {
			p.configBatchQueryResponse(*response)
		} else {
			fmt.Printf("[polarisctl internal err] print failed, covert to service_manage.Response failed\n")
		}
	case *config_manage.ConfigBatchWriteResponse:
		if response, ok := result.(*config_manage.ConfigBatchWriteResponse); ok {
			p.configBatchWriteResponse(*response)
		} else {
			fmt.Printf("[polarisctl internal err] print failed, covert to service_manage.Response failed\n")
		}
	case *SDKClient:
		if response, ok := result.(*SDKClient); ok {
			p.sdkClient(*response)
		} else {
			fmt.Printf("[polarisctl internal err] print failed, covert to service_manage.Response failed\n")
		}
	case *[]LeadersResponse:
		if response, ok := result.(*[]LeadersResponse); ok {
			p.leadersResponse(*response)
		} else {
			fmt.Printf("[polarisctl internal err] print failed, covert to service_manage.Response failed\n")
		}
	case *[]CMDBResponse:
		if response, ok := result.(*[]CMDBResponse); ok {
			p.cmdbResponse(*response)
		} else {
			fmt.Printf("[polarisctl internal err] print failed, covert to service_manage.Response failed\n")
		}
	case *[]LogLevelResponse:
		if response, ok := result.(*[]LogLevelResponse); ok {
			p.logLevelResponse(*response)
		} else {
			fmt.Printf("[polarisctl internal err] print failed, covert to service_manage.Response failed\n")
		}
	default:
		fmt.Printf("[polarisctl internal err] print failed,unkown response type:%T\n", v)
	}
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

	codec := p.codecFactory.Get(msgName)
	ret = append(ret, msgName)
	ret = append(ret, string(codec.Codec(res)))
	return ret

}

func (p *TableWriter) response(response interface{}) {
	if response == nil {
		p.table.AddRow([]string{"unkown", "nil", "0", "0"})
		return
	}
	resValue := reflect.ValueOf(response).Elem()
	code := fmt.Sprintf("%v", resValue.FieldByName("Code").Elem().FieldByName("Value").Interface())
	info := fmt.Sprintf("%v", resValue.FieldByName("Info").Elem().FieldByName("Value").Interface())
	for rsName, fieldName := range getAllFieldNames("v1.Response", response) {

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
			}
			continue
		}

		codec := p.codecFactory.Get(rsName)
		p.table.AddRow(
			[]string{rsName, string(codec.Codec(fieldValue.Interface())), code, info},
		)
	}

}

func (p *TableWriter) batchQueryResponse(response service_manage.BatchQueryResponse) {
	p.table, _ = gotable.Create("code", "info", "amount", "size")
	p.table.AddRow([]string{
		codeToStr(response.Code.Value),
		response.Info.Value,
		codeToStr(response.Amount.Value),
		codeToStr(response.Size.Value),
	})

	fmt.Println(p.table)

	p.table, _ = gotable.Create("type", "resource")
	// print resource with config
	resValue := reflect.ValueOf(&response).Elem()
	p.batchQueryResp(getAllFieldNames("v1.BatchQueryResponse", &response), resValue)
	fmt.Println(p.table)
}

func (p *TableWriter) batchWriteResponse(response service_manage.BatchWriteResponse) {

	succ, failed := uint32(0), uint32(0)
	for _, res := range response.Responses {
		if res.Code.Value == uint32(model.Code_ExecuteSuccess) {
			succ += 1
		} else {
			failed += 1
		}
	}
	p.table, _ = gotable.Create("code", "info", "size", "succ", "failed")
	p.table.AddRow([]string{
		codeToStr(response.Code.Value),
		response.Info.Value,
		codeToStr(response.Size.Value),
		codeToStr(succ),
		codeToStr(failed),
	})
	fmt.Println(p.table)

	p.table, _ = gotable.Create("type", "resource", "code", "info")
	for _, res := range response.Responses {
		p.response(res)
	}
	fmt.Println(p.table)

}

func (p *TableWriter) configBatchQueryResponse(response config_manage.ConfigBatchQueryResponse) {

	p.table, _ = gotable.Create("code", "info", "total")

	total := "nil"
	if nil != response.Total {
		total = codeToStr(response.Total.Value)
	}
	p.table.AddRow([]string{
		codeToStr(response.Code.Value),
		response.Info.Value,
		total,
	})

	fmt.Println(p.table)

	p.table, _ = gotable.Create("type", "resource")
	// print resource with config
	resValue := reflect.ValueOf(&response).Elem()
	p.batchQueryResp(getAllFieldNames("v1.ConfigBatchQueryResponse", &response), resValue)
	fmt.Println(p.table)
}

func (p *TableWriter) configBatchWriteResponse(response config_manage.ConfigBatchWriteResponse) {

	fmt.Printf("response:%+v\n", response)
}

func (p *TableWriter) sdkClient(response SDKClient) {

	fmt.Printf("response:%+v\n", response)
}

func (p *TableWriter) leadersResponse(response []LeadersResponse) {

	fmt.Printf("response:%+v\n", response)
}

func (p *TableWriter) cmdbResponse(response []CMDBResponse) {

	fmt.Printf("response:%+v\n", response)
}

func (p *TableWriter) logLevelResponse(response []LogLevelResponse) {

	fmt.Printf("response:%+v\n", response)
}

func codeToStr(value uint32) string {
	return strconv.FormatUint(uint64(value), 10)
}

func getAllFieldNames(rsName string, response interface{}) map[string]string {

	rets := map[string]string{}

	resType := reflect.TypeOf(response).Elem()
	resVal := reflect.ValueOf(response).Elem()
	filedNum := resType.NumField()

	for i := 0; i < filedNum; i++ {
		typeField := resType.Field(i)
		protoTag := typeField.Tag.Get("protobuf")
		if len(protoTag) == 0 {
			continue
		}
		tagName := strings.Split(typeField.Tag.Get("json"), ",")[0]
		if tagName == "" {
			fmt.Printf("[polarisctl internal sys err] cannot find tag %s in %s\n", tagName, rsName)
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
			codec := p.codecFactory.Get(rsName)
			p.table.AddRow([]string{rsName, string(codec.Codec(fieldValue.Index(i).Interface()))})
		}
	}
}

// getFieldName response must be pointer
func getFieldNames(rsName string, response interface{}, tagNames []string) [][]string {

	rets := make([][]string, 0, 0)

	resType := reflect.TypeOf(response).Elem()
	resVal := reflect.ValueOf(response).Elem()
	filedNum := resType.NumField()
	unkownTags := []string{}

	for _, tagName := range tagNames {
		findField := false
		for i := 0; i < filedNum; i++ {
			tag := strings.Split(resType.Field(i).Tag.Get("json"), ",")[0]
			if tag == "" {
				fmt.Printf("[polarisctl internal sys err] cannot find tag %s in %s\n", tagName, rsName)
				continue
			}

			if tag != tagName {
				continue
			}

			findField = true
			fieldVal := resVal.Field(i)

			if fieldVal.Kind() == reflect.Slice {
				elemType := fieldVal.Type().Elem()   // 获取 Slice 定义的元素类型
				elem := reflect.New(elemType).Elem() // 构造一个该元素类型的空值
				fieldVal = elem
			}
			msgName := proto.MessageName(fieldVal.Interface().(proto.Message))

			ret := []string{msgName, tagName, resType.Field(i).Name}
			rets = append(rets, ret)
		}

		if !findField {
			unkownTags = append(unkownTags, tagName)
		}

	}

	if 0 != len(unkownTags) {
		fmt.Printf("[polarisctl internal sys err] cannot find tag %v in %s\n", unkownTags, rsName)
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
