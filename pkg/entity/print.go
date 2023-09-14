package entity

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"

	"github.com/polarismesh/specification/source/go/api/v1/model"
	"github.com/polarismesh/specification/source/go/api/v1/service_manage"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// PolarisPrint 输出执行结果
type PolarisPrint struct {
	result interface{}
	writer *tabwriter.Writer
}

// NewPolarisPrint 构造
func NewPolarisPrint(rs interface{}) *PolarisPrint {
	return &PolarisPrint{
		result: rs,
		writer: tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.StripEscape|tabwriter.AlignRight|tabwriter.Debug|tabwriter.TabIndent|tabwriter.DiscardEmptyColumns|tabwriter.TabIndent),
	}
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
	fmt.Println("print response unimpl")
}

func (p PolarisPrint) printBatchWriteResponse(response service_manage.BatchWriteResponse) {

	fmt.Println("print batch write response unimpl")
}

func (p PolarisPrint) printBatchQueryResponse(response service_manage.BatchQueryResponse) {
	fmt.Println("====================== query response   =========================")
	fmt.Fprintf(p.writer, "\tcode\tinfo\tamount\tsize\t\n")
	fmt.Fprintf(p.writer, "\t%d\t%s\t%d\t%d\t\n", response.Code.Value, response.Info.Value, response.Amount.Value, response.Size.Value)
	p.writer.Flush()

	if 0 != len(response.Namespaces) {
		p.printNamespces(response.Namespaces)
	}
}

func (p PolarisPrint) printRsTabHeader(rs interface{}) []string {

	v := reflect.ValueOf(rs)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	typeOfT := v.Type()

	filedNum := v.NumField()
	fieldNames := []string{}

	fmt.Fprintf(p.writer, "\t")
	for i := 0; i < filedNum; i++ {
		tag := strings.Split(typeOfT.Field(i).Tag.Get("json"), ",")[0]
		name := typeOfT.Field(i).Name
		if tag == "" {
			continue
		}
		fmt.Fprintf(p.writer, "%s\t", tag)
		fieldNames = append(fieldNames, name)
	}
	fmt.Fprintf(p.writer, "\n")

	return fieldNames
}

func (p PolarisPrint) printRs(rs interface{}, fieldNames []string) {
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

func (p PolarisPrint) printNamespces(nas []*model.Namespace) {
	fmt.Fprintln(p.writer, "\n======================    namespaces    =========================")

	names := p.printRsTabHeader(nas[0])
	// 遍历 namespaces 并输出每个消息
	for _, ns := range nas {
		p.printRs(ns, names)
	}
	p.writer.Flush()
}

func (p PolarisPrint) printServices() {

	fmt.Println("\n======================    services    =========================")
}
func (p PolarisPrint) printInstances() {
	fmt.Println("\n======================    instances    =========================")

}
