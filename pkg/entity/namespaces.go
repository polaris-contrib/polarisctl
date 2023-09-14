package entity

import (
	"fmt"
	"os"
	"reflect"
	"text/tabwriter"

	"github.com/polarismesh/specification/source/go/api/v1/model"
)

var nsTags []string
var nsTagIndex map[string]int
var tableHeader string

func init() {
	ns := model.Namespace{}
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

// NamespaceDumps 将 namespace 按 table 输出到 stdout
func NamespaceDump(data []model.Namespace) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.StripEscape|tabwriter.AlignRight|tabwriter.Debug|tabwriter.TabIndent|tabwriter.DiscardEmptyColumns|tabwriter.TabIndent)
	fmt.Fprintf(w, "%s\n", tableHeader)

	for _, ns := range data {
		v := reflect.ValueOf(ns)
		for _, tag := range nsTags {
			str := fmt.Sprintf("%v\t", v.Field(nsTagIndex[tag]))
			fmt.Fprintf(w, "%s\t", str)
		}
		fmt.Fprintf(w, "\n")
	}
	w.Flush()
}
