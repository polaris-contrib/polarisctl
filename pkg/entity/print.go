package entity

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/polarismesh/specification/source/go/api/v1/model"
	"github.com/polarismesh/specification/source/go/api/v1/service_manage"
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
		"optionSwitch", "instanceLabels", "data",
	}

	resourcesTags = []string{
		"namespaces", "services", "instances",
		"routings", "aliases", "rateLimits",
		"configWithServices", "users", "userGroups",
		"authStrategies", "clients", "summary"}
}

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
}

// NewPolarisPrint 构造
func NewPolarisPrint() *PolarisPrint {
	return &PolarisPrint{
		writer:       tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.StripEscape|tabwriter.AlignRight|tabwriter.Debug|tabwriter.TabIndent|tabwriter.DiscardEmptyColumns|tabwriter.TabIndent),
		resourceConf: map[string][]string{},
		filterTags:   "|remove_user_ids|remove_group_ids|",
	}
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

	p.writer.Flush()
}

// resourceWriteResult print resource writ result
func (p PolarisPrint) resourceWriteResult(code uint32, info string, rsTypeName string, rs interface{}) {
	// TODO(karlyzhang)	插件化不同资源的输出格式
	fmt.Fprintf(p.writer, "%s\t", rsTypeName)
	value := reflect.ValueOf(rs).Elem()
	var field reflect.Value
	str := "<unkown>"
	if rsTypeName == "Alias" {
		field = value.FieldByName("Name")
		str = field.Interface().(*wrapperspb.StringValue).GetValue()
	} else if rsTypeName == "instance" {
		host := value.FieldByName("Host")
		port := value.FieldByName("Port")
		portStr := strconv.FormatUint(uint64(port.Interface().(*wrapperspb.UInt32Value).GetValue()), 10)
		hostStr := host.Interface().(*wrapperspb.StringValue).GetValue()
		str = hostStr + ":" + portStr
	} else {
		field = value.FieldByName("Alias")
		str = field.Interface().(*wrapperspb.StringValue).GetValue()
	}

	//if field.IsValid() {
	//	str = field.Interface().(*wrapperspb.StringValue).GetValue()
	//}
	fmt.Fprintf(p.writer, "%s\t", str)
	fmt.Fprintf(p.writer, "%d\t", code)
	fmt.Fprintf(p.writer, "%s\t", info)
	fmt.Fprintln(p.writer)
}

// printBatchQueryResponse print batch query response
func (p PolarisPrint) printBatchQueryResponse(response service_manage.BatchQueryResponse) {

	// batch query result
	fmt.Println("====================== query response   =========================")
	fmt.Fprintf(p.writer, "code\tinfo\tamount\tsize\t\n")
	fmt.Fprintf(p.writer, "%d\t%s\t%d\t%d\t\n", response.Code.Value, response.Info.Value, response.Amount.Value, response.Size.Value)
	p.writer.Flush()

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

// resource print resource message
func (p PolarisPrint) resource(rs interface{}, fieldNames []string) {
	value := reflect.ValueOf(rs).Elem()
	for _, name := range fieldNames {
		field := value.FieldByName(name)
		fieldType := field.Type()
		// 输出字段值
		switch fieldType {
		case reflect.TypeOf(&wrapperspb.StringValue{}):
			fmt.Fprintf(p.writer, "%s\t", field.Interface().(*wrapperspb.StringValue).GetValue())
		case reflect.TypeOf(&wrapperspb.BoolValue{}):
			fmt.Fprintf(p.writer, "%t\t", field.Interface().(*wrapperspb.BoolValue).GetValue())
		case reflect.TypeOf(&wrapperspb.UInt32Value{}):
			fmt.Fprintf(p.writer, "%d\t", field.Interface().(*wrapperspb.UInt32Value).GetValue())
		case reflect.TypeOf([]*wrapperspb.StringValue{}):
			stringValues := field.Interface().([]*wrapperspb.StringValue)
			for j, strValue := range stringValues {
				if j > 0 {
					fmt.Fprint(p.writer, ",")
				}
				fmt.Fprint(p.writer, strValue.GetValue())
			}
			fmt.Fprint(p.writer, "\t")
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
		p.resource(resources.Index(i).Interface(), names)
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
