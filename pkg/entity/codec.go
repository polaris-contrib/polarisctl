package entity

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"

	"github.com/golang/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

type TableCodec interface {
	Codec(interface{}) []byte
	PassFilter(bool) TableCodec
}

type defaultCodec struct {
	rsName     string
	passFilter bool
}

func NewCodec(rsName string) TableCodec {
	return &defaultCodec{rsName, false}
}

func (codec *defaultCodec) filterTag(tagName string) bool {
	if codec.passFilter {
		return false
	}
	return !printConf.FindTag(tagName)
}

func (codec *defaultCodec) PassFilter(value bool) TableCodec {
	codec.passFilter = value
	return codec
}

func (codec *defaultCodec) Codec(resource interface{}) []byte {
	if nil == resource {
		return []byte("nil")
	}

	ret := bytes.NewBuffer([]byte{})
	resType := reflect.TypeOf(resource).Elem()
	resVal := reflect.ValueOf(resource).Elem()
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
				ret.Write(NewRsCodeFactory().Get(msgName).PassFilter(true).Codec(res))

				continue
			}
			if strings.Contains(field.Type().String(), "wrapperspb") {
				ret.WriteString(fmt.Sprintf("%v", field.Elem().FieldByName("Value").Interface()))
			} else {
				ret.WriteByte('{')
				ret.Write((&defaultCodec{}).PassFilter(true).Codec(field.Interface()))
				ret.WriteByte('}')
			}
		case reflect.Slice:
			ret.WriteByte('[')
			for i := 0; i < field.Len(); i++ {
				if strings.Contains(field.Index(i).Type().String(), "wrapperspb") {
					ret.WriteString(fmt.Sprintf("%v", field.Index(i).Elem().FieldByName("Value").Interface()))
				} else {
					ret.Write((&defaultCodec{}).PassFilter(true).Codec(field.Index(i).Interface()))
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
			ret.Write((&defaultCodec{}).PassFilter(true).Codec(field.Interface()))
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

type RsCodecFactory struct {
	names map[string]TableCodec
}

func NewRsCodeFactory() *RsCodecFactory {
	return &RsCodecFactory{
		names: map[string]TableCodec{
			"v1.Service":             NewCodec("v1.Service"),
			"v1.ServiceAlias":        NewCodec("v1.ServiceAlias"),
			"v1.Instance":            NewCodec("v1.Instance"),
			"v1.Routing":             NewCodec("v1.Routing"),
			"v1.Rule":                NewCodec("v1.Rule"),
			"v1.Namespace":           NewCodec("v1.Namespace"),
			"v1.CircuitBreaker":      NewCodec("v1.CircuitBreaker"),
			"v1.CircuitBreakerRule":  NewCodec("v1.CircuitBreakerRule"),
			"v1.ConfigRelease":       NewCodec("v1.ConfigRelease"),
			"v1.User":                NewCodec("v1.User"),
			"v1.UserGroup":           NewCodec("v1.UserGroup"),
			"v1.AuthStrategy":        NewCodec("v1.AuthStrategy"),
			"v1.UserGroupRelation":   NewCodec("v1.UserGroupRelation"),
			"v1.LoginResponse":       NewCodec("v1.LoginResponse"),
			"v1.ModifyUserGroup":     NewCodec("v1.ModifyUserGroup"),
			"v1.StrategyResources":   NewCodec("v1.StrategyResources"),
			"v1.OptionSwitch":        NewCodec("v1.OptionSwitch"),
			"v1.InstanceLables":      NewCodec("v1.InstanceLables"),
			"v1.Client":              NewCodec("v1.Client"),
			"v1.StatInfo":            NewCodec("v1.StatInfo"),
			"v1.ConfigWithService":   NewCodec("v1.ConfigWithService"),
			"v1.ServiceContract":     NewCodec("v1.ServiceContract"),
			"v1.InterfaceDescriptor": NewCodec("v1.InterfaceDescriptor"),
			"v1.Heartbeatrecord":     NewCodec("v1.Heartbeatrecord"),
			"v1.InstanceHeartbeat":   NewCodec("v1.InstanceHeartbeat"),
			"v1.HeartbeatsRequest":   NewCodec("v1.HeartbeatsRequest"),
			"v1.HeartbeatsResponse":  NewCodec("v1.HeartbeatsResponse"),
		},
	}
}

func (f *RsCodecFactory) Get(name string) TableCodec {
	if codec, ok := f.names[name]; ok {
		return codec
	}

	return NewCodec(name)
}
