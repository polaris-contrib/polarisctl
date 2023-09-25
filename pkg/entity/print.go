package entity

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/golang/protobuf/proto"
	"github.com/polarismesh/specification/source/go/api/v1/model"
	"github.com/polarismesh/specification/source/go/api/v1/service_manage"
	_ "github.com/polarismesh/specification/source/go/api/v1/traffic_manage"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var resourcesTags []string
var resourceTags []string

func init() {
	resourceTags = []string{
		"namespace", "service", "instance",
		"routing", "alias", "rateLimit",
		"user", "userGroup", "authStrategy",
		"client", "circuitBreaker", "relation",
		"loginResponse", "modifyAuthStrategy",
		"modifyUserGroup", "resources",
		"optionSwitch", "instanceLabels",
	}

	resourcesTags = []string{
		"namespaces", "services", "instances",
		"routings", "aliases", "rateLimits",
		"configWithServices", "users", "userGroups",
		"authStrategies", "clients", "summary"}
}

// TODO(karlyzhang): 拆分 print，支持按资源指定输出插件

// PolarisPrint 输出执行结果
type PolarisPrint struct {
	result interface{}
	writer *tabwriter.Writer

	// tag -- > field :BatchQueryResponse print field conf
	queryConf map[string]string
	// tag -- > field :BatchWriteResponse print field conf
	writeConf map[string]string

	// resourceName -> []string{tags,tags} resource field print conf
	resourceConf map[string][]string

	filterTags string
	v2Opt      *v2PrintOpt
}

type v2PrintOpt struct {
	resourceName string
}

// NewPolarisPrint 构造
func NewPolarisPrint() *PolarisPrint {
	return &PolarisPrint{
		//writer:       tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.StripEscape|tabwriter.AlignRight|tabwriter.Debug|tabwriter.TabIndent|tabwriter.DiscardEmptyColumns|tabwriter.TabIndent),
		writer:       tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug),
		resourceConf: map[string][]string{},
		filterTags:   "|remove_user_ids|remove_group_ids|",
		v2Opt:        nil,
	}
}

// V2Api set v2 api resource print
func (p *PolarisPrint) V2Api(resourceName string) *PolarisPrint {
	p.v2Opt = &v2PrintOpt{resourceName: resourceName}
	return p
}

// Response set result
func (p *PolarisPrint) Response(response interface{}) *PolarisPrint {
	p.result = response
	return p
}

// BatchQuery build batch query common tags
func (p *PolarisPrint) BatchQuery() *PolarisPrint {
	response := &service_manage.BatchQueryResponse{}
	p.queryConf = findFieldNames("BatchQueryResponse", response, resourcesTags)
	return p
}

// Resource set resource print conf
func (p *PolarisPrint) ResourceConf(resource, value string) *PolarisPrint {
	if len(value) == 0 {
		return p
	}
	arr := strings.Split(value, ",")
	for i, s := range arr {
		arr[i] = strings.TrimSpace(s)
	}
	p.resourceConf[resource] = arr
	return p
}

// Query build  query common tags
func (p *PolarisPrint) Query() *PolarisPrint {
	return p.BatchWrite()
}

// Write build  write common tags
func (p *PolarisPrint) Write() *PolarisPrint {
	return p.BatchWrite()
}

// BatchQuery build batch write common tags
func (p *PolarisPrint) BatchWrite() *PolarisPrint {
	response := &service_manage.Response{}
	p.writeConf = findFieldNames("Response", response, resourceTags)
	return p
}

// Print 输出结果
func (p PolarisPrint) Print() {
	switch v := p.result.(type) {
	case service_manage.Response:
		if response, ok := p.result.(service_manage.Response); ok {
			p.printResponse(response)
		} else {
			fmt.Printf("[polarisctl internal err] print failed, covert to service_manage.Response failed\n")
		}
	case service_manage.BatchWriteResponse:
		if response, ok := p.result.(service_manage.BatchWriteResponse); ok {
			p.printBatchWriteResponse(response)
		} else {
			fmt.Printf("[polarisctl internal err] print failed, covert to service_manage.BatchWriteResponse failed\n")
		}
	case service_manage.BatchQueryResponse:
		if response, ok := p.result.(service_manage.BatchQueryResponse); ok {
			p.printBatchQueryResponse(response)
		} else {
			fmt.Printf("[polarisctl internal err] print failed, covert to service_manage.Response failed\n")
		}
	default:
		fmt.Printf("[polarisctl internal err] print failed,unkown response type:%T\n", v)
	}
}

func (p PolarisPrint) printResponse(response service_manage.Response) {
	defer p.writer.Flush()
	fmt.Fprintln(p.writer, "\n======================  response   =================================\n")
	fmt.Fprintf(p.writer, "code\tinfo\t\n")
	fmt.Fprintf(p.writer, "%d\t%s\t\n", response.Code.Value, response.Info.Value)

	fmt.Fprintln(p.writer, "\n======================  resource   =================================\n")
	fmt.Fprintf(p.writer, "resource type\tresource name\tcode\tinfo\t\n")
	for tagName, fieldName := range p.writeConf {
		if tagName == "instanceLabels" {
			continue
		}
		resValue := reflect.ValueOf(&response).Elem()
		fieldValue := resValue.FieldByName(fieldName)
		if fieldValue.IsNil() {
			continue
		}
		p.resourceWriteResult(response.Code.Value, response.Info.Value, tagName, fieldValue.Interface())
	}

	// TODO(karlyzhang): 将资源的输出插件化，为特殊的资源制定特殊的输出格式
	if fieldName, ok := p.writeConf["instanceLabels"]; ok {
		resValue := reflect.ValueOf(&response).Elem()
		fieldValue := resValue.FieldByName(fieldName)
		if !fieldValue.IsNil() {
			p.instanceLabels(fieldValue.Interface().(*service_manage.InstanceLabels))
		}
	}
}

func (p PolarisPrint) printBatchWriteResponse(response service_manage.BatchWriteResponse) {
	defer p.writer.Flush()

	fmt.Fprintln(p.writer, "====================== batch write result   =========================\n")
	fmt.Fprintf(p.writer, "code\tinfo\tsize\tsucc_size\tfailed_size\t\n")

	succWrite, failedWrite := 0, 0
	for _, res := range response.Responses {
		if res.Code.Value == uint32(model.Code_ExecuteSuccess) {
			succWrite += 1
		} else {
			failedWrite += 1
		}
	}
	fmt.Fprintf(p.writer, "%d\t%s\t%d\t%d\t%d\t\n", response.Code.Value, response.Info.Value, response.Size.Value, succWrite, failedWrite)
	p.writer.Flush()

	// TODO(karlyzhang): 拆分 v2 输出逻辑，插件化处理
	if p.v2Opt != nil {
		fmt.Fprintln(p.writer, "\n======================  resources   =================================\n")
		fmt.Fprintf(p.writer, "resource type\tresource name\tcode\tinfo\t\n")
		p.writer.Flush()
		for _, res := range response.Responses {
			if res.Data == nil {
				continue
			}
			err, urlName, resource := unmarshalAny(res.Data)
			if err != nil {
				continue
			}
			p.resourceWriteResult(res.Code.Value, res.Info.Value, urlName, resource)
		}
		return
	}

	if len(response.Responses) == 0 {
		return
	}

	fmt.Fprintln(p.writer, "\n======================  resources   =================================\n")
	fmt.Fprintf(p.writer, "resource type\tresource name\tcode\tinfo\t\n")

	for tagName, fieldName := range p.writeConf {
		for _, res := range response.Responses {
			resValue := reflect.ValueOf(res).Elem()
			fieldValue := resValue.FieldByName(fieldName)
			if fieldValue.IsNil() {
				continue
			}
			p.resourceWriteResult(res.Code.Value, res.Info.Value, tagName, fieldValue.Interface())
		}
	}

}

// printBatchQueryResponse print batch query response
func (p PolarisPrint) printBatchQueryResponse(response service_manage.BatchQueryResponse) {
	defer p.writer.Flush()

	// batch query result
	fmt.Println("====================== query response   =========================")
	fmt.Fprintf(p.writer, "code\tinfo\tamount\tsize\t\n")
	fmt.Fprintf(p.writer, "%d\t%s\t%d\t%d\t\n", response.Code.Value, response.Info.Value, response.Amount.Value, response.Size.Value)
	p.writer.Flush()

	// v2 api
	if p.v2Opt != nil {

		if len(response.Data) == 0 {
			return
		}

		printHeader := false
		fieldNames := []string{}
		// same type
		for _, any := range response.Data {

			err, urlName, resource := unmarshalAny(any)
			if err != nil {
				continue
			}
			if !printHeader {
				printHeader = true
				fieldNames = p.resourceTabHeader(urlName, resource)
				fmt.Fprintln(p.writer, strings.Repeat("-", len(fieldNames)*10))
				p.writer.Flush()
			}
			p.resourcev2(resource, fieldNames)

		}
		return
	}

	// print resource with config
	resValue := reflect.ValueOf(&response).Elem()
	for tagName, fieldName := range p.queryConf {
		fieldValue := resValue.FieldByName(fieldName)
		if fieldValue.IsNil() {
			continue
		}

		if fieldValue.Type().Kind() != reflect.Slice {
			continue
		}

		if fieldValue.Len() == 0 {
			continue
		}
		p.resources(tagName, fieldValue)
	}
}

// resourceWriteResult print resource writ result
func (p PolarisPrint) resourceWriteResult(code uint32, info string, rsTypeName string, rs interface{}) {
	// TODO(karlyzhang)	插件化不同资源的输出格式
	fmt.Fprintf(p.writer, "%s\t", rsTypeName)
	value := reflect.ValueOf(rs).Elem()
	var field reflect.Value
	str := "<unkown>"
	if rsTypeName == "Alias" {
		field = value.FieldByName("Alias")
		str = field.Interface().(*wrapperspb.StringValue).GetValue()
	} else if rsTypeName == "instance" {
		host := value.FieldByName("Host")
		port := value.FieldByName("Port")
		portStr := strconv.FormatUint(uint64(port.Interface().(*wrapperspb.UInt32Value).GetValue()), 10)
		hostStr := host.Interface().(*wrapperspb.StringValue).GetValue()
		str = hostStr + ":" + portStr
	} else {
		// 首先使用 name
		str = findFieldValue(rs, []string{"name", "id"})
		//field = value.FieldByName("Name")
		//str = field.Interface().(*wrapperspb.StringValue).GetValue()
	}

	//if field.IsValid() {
	//	str = field.Interface().(*wrapperspb.StringValue).GetValue()
	//}
	fmt.Fprintf(p.writer, "%s\t", str)
	fmt.Fprintf(p.writer, "%d\t", code)
	fmt.Fprintf(p.writer, "%s\t", info)
	fmt.Fprintln(p.writer)
}

// resourceTabHeader print resource tab header
func (p PolarisPrint) resourceTabHeader(name string, resource interface{}) []string {
	fieldNames := []string{}
	// use conf
	if _, ok := p.resourceConf[name]; ok && len(p.resourceConf[name]) != 0 {
		tagFields := findFieldNames(name, resource, p.resourceConf[name])
		for _, tag := range p.resourceConf[name] {
			if fieldName, ok := tagFields[tag]; ok {
				if strings.Contains(p.filterTags, "|"+tag+"|") {
					continue
				}
				fmt.Fprintf(p.writer, "%s\t", tag)
				fieldNames = append(fieldNames, fieldName)
			}
		}
		fmt.Fprintln(p.writer)
		return fieldNames
	}

	resValue := reflect.ValueOf(resource)
	if resValue.Kind() == reflect.Ptr {
		resValue = resValue.Elem()
	}

	resType := resValue.Type()

	filedNum := resValue.NumField()

	// print all field
	for i := 0; i < filedNum; i++ {
		tag := strings.Split(resType.Field(i).Tag.Get("json"), ",")[0]
		name := resType.Field(i).Name
		if tag == "" || tag == "-" {
			continue
		}
		if strings.Contains(p.filterTags, "|"+tag+"|") {
			continue
		}
		fmt.Fprintf(p.writer, "%s\t", tag)
		fieldNames = append(fieldNames, name)
	}
	fmt.Fprintf(p.writer, "\n")

	return fieldNames
}

func (p PolarisPrint) resourcev2(rs interface{}, fieldNames []string) {
	value := reflect.ValueOf(rs).Elem()

	for _, name := range fieldNames {
		field := value.FieldByName(name)
		// 输出字段值
		switch field.Kind() {
		case reflect.Ptr:
			p.printPtr(name, field)
		case reflect.Slice:
			var sb strings.Builder
			for i := 0; i < field.Len(); i++ {
				//elem := field.Index(i).Interface()
				elem := field.Index(i)
				if i > 0 {
					sb.WriteString(",")
				}
				sb.WriteString(p.fetchFieldValue(elem))
			}
			sb.WriteString("\t")
			fmt.Fprint(p.writer, sb.String())
		case reflect.String, reflect.Bool, reflect.Int, reflect.Int32, reflect.Int64, reflect.Uint32, reflect.Uint64:
			fmt.Fprintf(p.writer, "%v\t", field.Interface())
		case reflect.Map:
			fmt.Fprintf(p.writer, "%v\t", field.Interface())
		case reflect.Struct:
			fmt.Fprintf(p.writer, "%v\t", field.Interface())
		default:
			fmt.Fprintf(p.writer, "<unkown type,t:%T k:%v>\t", field.Interface(), field.Kind())
		}
	}
	fmt.Fprintln(p.writer)
}

// func (p *PolarisPrint) fetchFieldValue(elem interface{}) string {
func (p *PolarisPrint) fetchFieldValue(elem reflect.Value) string {

	return fmt.Sprintf("elem %T %v", elem.Interface(), elem.Kind())
	//switch v := elem.(type) {
	//	case *wrapperspb.StringValue:
	//		return v.GetValue()
	//	case *wrapperspb.Int32Value:
	//		return v.GetValue()
	//	case *string:
	//		return *v
	//	case *int:
	//		return strconv.Itoa(*v)
	//	default:
	//		return fmt.Sprintf("%v", v)
	//	}
}

func (p PolarisPrint) printPtr(name string, field reflect.Value) {
	if !field.IsValid() || field.IsZero() {
		fmt.Fprintf(p.writer, "null\t")
		return
	}
	v := field.Interface()
	switch v := v.(type) {
	case *wrapperspb.StringValue:
		fmt.Fprintf(p.writer, "%s\t", v.GetValue())
	case *wrapperspb.BoolValue, *wrapperspb.UInt32Value:
		fmt.Fprintf(p.writer, "%v\t", reflect.ValueOf(v).Elem().FieldByName("Value").Interface())
	case []*wrapperspb.StringValue:
		stringValues := field.Interface().([]*wrapperspb.StringValue)
		for j, strValue := range stringValues {
			if j > 0 {
				fmt.Fprint(p.writer, ",")
			}
			fmt.Fprint(p.writer, strValue.GetValue())
		}
		fmt.Fprint(p.writer, "\t")
	case *anypb.Any:
		err, _, resource := unmarshalAny(v)
		if err == nil {
			fmt.Fprintf(p.writer, "%+v\t", resource)
		} else {
			fmt.Fprintf(p.writer, "%s,%v,<err:%v>\t", v, err)
		}
	default:
		if field.Type().Elem().Kind() == reflect.Struct {
			fmt.Fprintf(p.writer, "%v\t", v)
		} else {
			fmt.Fprintf(p.writer, "<unkown type,t:%T k:%v>\t", field.Type(), field.Type().Kind())
		}
	}
}
func (p PolarisPrint) resource(rs interface{}, fieldNames []string) {
	value := reflect.ValueOf(rs).Elem()
	for _, name := range fieldNames {
		field := value.FieldByName(name)
		// 输出字段值
		switch v := field.Interface().(type) {
		case *wrapperspb.StringValue:
			fmt.Fprintf(p.writer, "%s\t", v.GetValue())
		case *wrapperspb.BoolValue, *wrapperspb.UInt32Value:
			fmt.Fprintf(p.writer, "%v\t", reflect.ValueOf(v).Elem().FieldByName("Value").Interface())
		case []*wrapperspb.StringValue:
			stringValues := field.Interface().([]*wrapperspb.StringValue)
			for j, strValue := range stringValues {
				if j > 0 {
					fmt.Fprint(p.writer, ",")
				}
				fmt.Fprint(p.writer, strValue.GetValue())
			}
			fmt.Fprint(p.writer, "\t")
		case string, bool, uint32, uint64, int, int32, int64:
			fmt.Fprintf(p.writer, "%v\t", v)
		case map[string]string:
			fmt.Fprintf(p.writer, "%v\t", v)
			str := ""
			for key, value := range v {
				str = str + key + ":" + value + ","
			}
			fmt.Fprintf(p.writer, "%s\t", str)
		case *anypb.Any:
			err, _, resource := unmarshalAny(v)
			if err == nil {
				fmt.Fprintf(p.writer, "%+v\t", resource)
			} else {
				fmt.Fprintf(p.writer, "%s,%v,<err:%v>\t", v, err)
			}
		default:
			if field.Type().Kind() == reflect.Int32 {
				// enum
				fmt.Fprintf(p.writer, "%v\t", v)
				continue
			}
			if field.Type().Kind() == reflect.Struct {
				fmt.Fprintf(p.writer, "%v\t", v)
			} else {
				//fmt.Printf("name:%s fieldType:%T unkown\n", name, field.Type())
				fmt.Fprintf(p.writer, "<unkown type,t:%T k:%v>\t", field.Type(), field.Type().Kind())
			}

		}
	}
	fmt.Fprintln(p.writer)
}

// resources print resource slice message
func (p PolarisPrint) resources(resourcesName string, resources reflect.Value) {
	fmt.Fprintf(p.writer, "\n======================    %s    =========================\n", resourcesName)
	names := p.resourceTabHeader(resourcesName, resources.Index(0).Interface())
	// 遍历 namespaces 并输出每个消息
	for i := 0; i < resources.Len(); i++ {
		p.resourcev2(resources.Index(i).Interface(), names)
	}
	p.writer.Flush()
}

// findFieldNames find tag fieldName
func findFieldNames(msg string, response interface{}, tagNames []string) map[string]string {

	fieldNames := map[string]string{}

	resType := reflect.TypeOf(response).Elem()
	filedNum := resType.NumField()

	for _, tagName := range tagNames {
		for i := 0; i < filedNum; i++ {
			tag := strings.Split(resType.Field(i).Tag.Get("json"), ",")[0]
			if tag == "" {
				fmt.Printf("[polarisctl internal sys err] cannot find tag %s in %s\n", tagName, msg)
				continue
			}
			if tag != tagName {
				continue
			}
			fieldNames[tagName] = resType.Field(i).Name
		}
	}

	// check fieldNames
	for _, tagName := range tagNames {
		if _, ok := fieldNames[tagName]; !ok {
			fmt.Printf("[polarisctl internal sys err] cannot find tag %s in %s\n", tagName, msg)
		}
	}
	return fieldNames
}

func (p *PolarisPrint) instanceLabels(labels *service_manage.InstanceLabels) {
	fmt.Fprintln(p.writer, "\n======================  labels   =================================\n")
	fmt.Fprintf(p.writer, "label\tvalue\t\n")
	for label, values := range labels.Labels {
		fmt.Fprintf(p.writer, "%s\t%v\t\n", label, values.Values)
	}
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

func findFieldValue(rs interface{}, tagNames []string) string {
	value := reflect.ValueOf(rs).Elem()
	typ := value.Type()

	for _, tagName := range tagNames {

		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			resType := typ.Field(i)

			tagValue := strings.Split(resType.Tag.Get("json"), ",")[0]
			if tagValue == "" || tagValue != tagName {
				continue
			}

			switch v := field.Interface().(type) {
			case *wrapperspb.StringValue:
				if len(v.GetValue()) == 0 {
					break
				}
				return v.GetValue()
			case *wrapperspb.BoolValue, *wrapperspb.UInt32Value:
				return fmt.Sprintf("%v", reflect.ValueOf(v).Elem().FieldByName("Value").Interface())
			case []*wrapperspb.StringValue:
				stringValues := field.Interface().([]*wrapperspb.StringValue)
				return fmt.Sprintf("%v", stringValues)
			case string:
				if len(v) == 0 {
					break
				}
				return fmt.Sprintf("%v", v)
			case bool, uint32, uint64, int, int32, int64:
				return fmt.Sprintf("%v", v)
			case map[string]string:
				return fmt.Sprintf("%v", v)
			case *anypb.Any:
				err, _, resource := unmarshalAny(v)
				if err == nil {
					return fmt.Sprintf("%+v", resource)
				}
				return fmt.Sprintf("%s,%v,<err:%v>", v, err)
			default:
				if field.Type().Kind() == reflect.Int32 {
					// enum
					return fmt.Sprintf("%v", v)
				}
				if field.Type().Kind() == reflect.Struct {
					return fmt.Sprintf("%v", v)
				} else {
					return fmt.Sprintf("<unkown type:%v>", field.Type())
				}
			}
		}
	}
	return fmt.Sprintf("<undefine field:%v>", tagNames)
}
