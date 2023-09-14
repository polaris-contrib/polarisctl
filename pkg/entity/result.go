package entity

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/polarismesh/specification/source/go/api/v1/model"
	"github.com/polarismesh/specification/source/go/api/v1/security"
	"github.com/polarismesh/specification/source/go/api/v1/service_manage"
	"github.com/polarismesh/specification/source/go/api/v1/traffic_manage"
)

// Result
type Result struct {
	Code              int                        `json:"code"`
	Info              string                     `json:"info"`
	Amount            int                        `json:"amount"`
	Size              int                        `json:"size"`
	Namespaces        []model.Namespace          `json:"namespaces"`
	Services          []service_manage.Service   `json:"services"`
	Instances         []service_manage.Instance  `json:"instance"`
	Routings          []traffic_manage.Routing   `json:"routings"`
	RateLimits        []traffic_manage.RateLimit `json:"rateLimits"`
	ConfiWithServices []service_manage.Service   `json:"configWithServices"`
	Users             []security.User            `json:"users"`
	UserGroups        []security.UserGroup       `json:"userGroups"`
	AuthStrategies    []security.AuthStrategy    `json:"authStrategies"`
	Clients           []service_manage.Client    `json:"clients"`
	Data              []PolarisData              `json:"data"`
	Summary           model.Summary              `json:"summary"`
}

// PolarisData
type PolarisData struct{}

// BatchResult
type BatchResult struct {
	Code      int    `json:"code"`
	Info      string `json:"info"`
	Size      int    `json:"size"`
	Responses []struct {
		Code          int                      `json:"code"`
		Info          string                   `json:"info"`
		Namespace     model.Namespace          `json:"namespace"`
		Service       service_manage.Service   `json:"service"`
		Instance      service_manage.Instance  `json:"instance"`
		Routing       traffic_manage.Routing   `json:"routing"`
		RateLimit     traffic_manage.RateLimit `json:"rateLimit"`
		User          security.User            `json:"user"`
		UserGroup     security.UserGroup       `json:"userGroup"`
		AuthStrategie security.AuthStrategy    `json:"authStrategie"`
		Data          PolarisData              `json:"data"`
	} `json:responses`
	SuccSize   int `table:"succ_size"`
	FailedSize int `table:"failed_size"`
}

// CheckRes 检查批量执行的结果，统计失败和成功的操作
func (res *BatchResult) CheckRes() {
	res.FailedSize = 0
	res.SuccSize = 0
	for _, r := range res.Responses {
		if model.Code(r.Code) == model.Code_ExecuteSuccess {
			res.SuccSize += 1
		} else {
			res.FailedSize += 1
		}
	}
}

// Dump 向 stdout 输出执行结果
func (res BatchResult) Dump() {
	fmt.Println("====================== result   =========================")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.StripEscape|tabwriter.AlignRight|tabwriter.Debug|tabwriter.TabIndent|tabwriter.DiscardEmptyColumns|tabwriter.TabIndent)
	fmt.Fprintf(w, "code\tinfo\tsize\tsucc_size\tfailed_size\t\n")
	fmt.Fprintf(w, "%d\t%s\t%d\t%d\t%d\t\n", res.Code, res.Info, res.Size, res.SuccSize, res.FailedSize)
	w.Flush()
}

// Dump 向 stdout 输出执行结果
func (result Result) Dump() {
	fmt.Println("====================== result   =========================")
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', tabwriter.StripEscape|tabwriter.AlignRight|tabwriter.Debug|tabwriter.TabIndent|tabwriter.DiscardEmptyColumns|tabwriter.TabIndent)
	fmt.Fprintf(w, "code\tinfo\tsize\tamount\t\n")
	fmt.Fprintf(w, "%d\t%s\t%d\t%d\t\n", result.Code, result.Info, result.Size, result.Amount)
	w.Flush()
}
