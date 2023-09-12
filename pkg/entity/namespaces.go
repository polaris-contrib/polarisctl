package entity

import (
	"fmt"
	"os"
	"reflect"
	"text/tabwriter"
)

// PolarisNamespace 查询时返回的namespace
type PolarisNamespace struct {
	Name                    string   `json:"name"`
	Comment                 string   `json:"comment"`
	Owners                  string   `json:"owners"`
	Token                   string   `json:"token"`
	Ctime                   string   `json:"ctime"`
	Mtime                   string   `json:"mtime"`
	TotalServiceCount       int      `json:"total_service_count"`
	TotalHealthInstancCount int      `json:"total_health_instance_count"`
	TotalInstanceCount      int      `json:"total_instance_count"`
	Id                      string   `json:"id"`
	GroupIds                []string `json:"group_ids"`
	UserIds                 []string `json:"user_ids"`
	RemoteGroupIds          []string `json:"remove_group_ids"`
	RemoteUserIds           []string `json:"remote_user_ids"`
	Editable                bool     `json:"editable"`
	ServiceExportTo         []string `json:"service_export_to"`
}

var nsTags []string
var nsTagIndex map[string]int
var tableHeader string

func init() {
	ns := PolarisNamespace{}
	v := reflect.ValueOf(ns)
	filedNum := v.NumField()
	nsTags = []string{}
	nsTagIndex = make(map[string]int, filedNum)
	for i := 0; i < filedNum; i++ {
		tag := v.Type().Field(i).Tag.Get("json")
		if tag == "" {
			continue
		}
		nsTags = append(nsTags, tag)
		nsTagIndex[tag] = i
		tableHeader += tag + "\t"
	}
}

// TabValue 按 tag 将值按 table 格式拼接
func (ns PolarisNamespace) TabValue() string {
	v := reflect.ValueOf(ns)
	str := ""
	for _, tag := range nsTags {
		str += fmt.Sprintf("%v\t", v.Field(nsTagIndex[tag]))
	}
	return str
}

// PrintNamespaces 将 namespace 按 table 输出到 stdout
func PrintNamespaces(data []PolarisNamespace) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.StripEscape|tabwriter.AlignRight|tabwriter.Debug|tabwriter.TabIndent|tabwriter.DiscardEmptyColumns|tabwriter.TabIndent)
	fmt.Fprintf(w, "%s\n", tableHeader)

	for _, ns := range data {
		fmt.Fprintf(w, "%s\n", ns.TabValue())
	}

	w.Flush()
}
