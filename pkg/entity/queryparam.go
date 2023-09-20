package entity

import (
	"fmt"
	"net/url"
	"os"
	"reflect"
	"strconv"

	"github.com/spf13/cobra"
)

// QueryParam get/list param for http
type QueryParam struct {
	Offset        int `param:"offset" short:"query offset(int)" default:"0"`
	Limit         int `param:"limit" short:"query limit(int)" default:"10"`
	ResourceParam interface{}
}

// ServicesQueryParam service 查询参数
type ServicesQueryParam struct {
	Name           string `param:"name" short:"service name"`
	Namespace      string `param:"namespace" short:"service naemspace"`
	Business       string `param:"business" short:"service business"`
	Department     string `param:"department" short:"service department"`
	Keys           string `param:"keys ,short:"service keys"`
	Values         string `param:"values" short:"service key values"`
	InstanceKeys   string `param:"instance_keys" short:" service instance keys"`
	InstanceValues string `param:"instance_values" short:"service instance key values"`
	Host           string `param:"host" short:"service host"`
	Port           string `param:"port" short:"service port"`
}

// AliasQueryParam namespace 查询参数
type AliasQueryParam struct {
	Alias          string `param:"alais",short:"service alias name"`
	AliasNamespace string `param:"alias_namespace",short:"service alias namespace name"`
	Serice         string `param:"service",short:"service name"`
	Namespace      string `param:"namespace",short:"service namespace name"`
}

// NamespaceQueryParam namespace 查询参数
type NamespacesQueryParam struct {
	Name string `param:"name",short:"namespace name"`
}

// LabelQueryParam namespace 查询参数
type LabelQueryParam struct {
	Namespace string `param:"namespace" short:"service namespace name" request:"true"`
	Service   string `param:"service" short:"service name" request:"true"`
}

// RoutingsQueryParam routings 查询参数
type RoutingsQueryParam struct {
	Id                   string `param:"id",short:"instances id"`
	Namespace            string `param:"namespace",short:"namespace name"`
	Service              string `param:"service",short:"service name"`
	Enable               string `param:"enable",short:"host name"`
	SourceService        string `param:"source_service",short:"host name"`
	SourceNamespace      string `param:"source_namespace",short:"host name"`
	DestinationService   string `param:"destination_service",short:"host name"`
	DestinationNamespace string `param:"destination_namespace",short:"host name"`
	OrderField           string `param:"order_field",short:"host name"`
	OrderType            string `param:"order_type",short:"host name"`
}

// InstancesQueryParam namespace 查询参数
type InstancesQueryParam struct {
	Id        string `param:"id",short:"instances id"`
	Namespace string `param:"namespace",short:"namespace name"`
	Service   string `param:"service",short:"service name"`
	Host      string `param:"host",short:"host name"`
	Keys      string `param:"keys",short:"host name"`
	Values    string `param:"values",short:"host name"`
	Healthy   string `param:"healthy",short:"host name"`
	Isolate   string `param:"isolate",short:"host name"`
	Protocol  string `param:"protocol",short:"host name"`
	Version   string `param:"version",short:"host name"`
}

// RegisterFlag 注册查询参数
func (param *LabelQueryParam) RegisterFlag(cmd *cobra.Command) {
	registerFlag(cmd, param)
}

// RegisterFlag 注册查询参数
func (param *LabelQueryParam) Encode() string {
	values := structToParam(param)
	return values.Encode()
}

// RegisterFlag 注册查询参数
func (param *QueryParam) RegisterFlag(cmd *cobra.Command) {
	registerFlag(cmd, param)
	registerFlag(cmd, param.ResourceParam)
}

func (param QueryParam) Encode() string {
	values := structToParam(&param)
	rsValues := structToParam(param.ResourceParam)
	for key, value := range rsValues {
		values[key] = value
	}
	return values.Encode()
}

func structToParam(param interface{}) url.Values {
	values := url.Values{}
	structType := reflect.TypeOf(param).Elem()
	structValue := reflect.ValueOf(param).Elem()
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		value := structValue.Field(i)

		tag := field.Tag.Get("param")
		if tag == "" {
			continue
		}

		var strValue string
		switch value.Kind() {
		case reflect.String:
			strValue = value.String()
		case reflect.Int:
			strValue = fmt.Sprint(value.Int())
		default:
			continue
		}

		if strValue != "" {
			values.Set(tag, strValue)
		}
	}
	return values
}

// registerFlag 将 struct 的字段注册为命令行参数
func registerFlag(cmd *cobra.Command, param interface{}) {
	structType := reflect.TypeOf(param).Elem()
	structValue := reflect.ValueOf(param).Elem()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		value := structValue.Field(i)

		tag := field.Tag.Get("param")
		short := field.Tag.Get("short")
		defaultVal := field.Tag.Get("default")
		request := field.Tag.Get("request")
		if tag == "" {
			continue
		}

		switch value.Kind() {
		case reflect.String:
			cmd.Flags().StringVar(value.Addr().Interface().(*string), tag, defaultVal, short)
			if len(request) != 0 && request == "true" {
				cmd.MarkFlagRequired(tag)
			}
		case reflect.Int:
			if len(defaultVal) == 0 {
				defaultVal = "0"
			}
			num, err := strconv.Atoi(defaultVal)
			if err != nil {
				fmt.Printf("[polarisctl internal sys err] register flag failed,tag:%s get unkown default value:%s,err:%v\n", tag, defaultVal, err)
				os.Exit(1)
			}
			if tag == "limit" || tag == "offset" {
				cmd.Flags().IntVarP(value.Addr().Interface().(*int), tag, string(tag[0]), num, short)
			} else {
				cmd.Flags().IntVar(value.Addr().Interface().(*int), tag, num, short)
			}

			if len(request) != 0 && request == "true" {
				cmd.MarkFlagRequired(tag)
			}
		default:
			fmt.Printf("[polarisctl internal sys err] register flag catch unkown type:%+v, tag:%s\n", value.Kind(), tag)
			continue
		}
	}
}
