package entity

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// Result
type Result struct {
	Code              int                     `json:"code"`
	Info              string                  `json:"info"`
	Amount            int                     `json:"amount"`
	Size              int                     `json:"size"`
	Namespaces        []PolarisNamespace      `json:"namespaces"`
	Services          []PolarisService        `json:"services"`
	Instances         []PolarisInstance       `json:"instance"`
	Routings          []PolarisRouting        `json:"routings"`
	RateLimits        []PolarisLimit          `json:"rateLimits"`
	ConfiWithServices []PolarisService        `json:"configWithServices"`
	Users             []PolarisUser           `json:"users"`
	UserGruops        []PolarisUserGroup      `json:"userGroups"`
	AuthStrategies    []PolarisAuthStrategies `json:"authStrategies"`
	Clients           []PolarisClient         `json:"clients"`
	Data              []PolarisData           `json:"data"`
	Summary           PolarisSummary          `json:"summary"`
}

// BatchResult
type BatchResult struct {
	Code      int    `json:"code"`
	Info      string `json:"info"`
	Size      int    `json:"size"`
	Responses []struct {
		Code          int                   `json:"code"`
		Info          string                `json:"info"`
		Namespace     PolarisNamespace      `json:"namespace"`
		Service       PolarisService        `json:"service"`
		Instance      PolarisInstance       `json:"instance"`
		Routing       PolarisRouting        `json:"routing"`
		RateLimit     PolarisLimit          `json:"rateLimit"`
		User          PolarisUser           `json:"user"`
		UserGroup     PolarisUserGroup      `json:"userGroup"`
		AuthStrategie PolarisAuthStrategies `json:"authStrategie"`
		Data          PolarisData           `json:"data"`
	} `json:responses`
	SuccSize   int `table:"succ_size"`
	FailedSize int `table:"failed_size"`
}

// PolarisSummary
type PolarisSummary struct {
	TotalServiceCount        int `json:"total_service_count"`
	TotalHealthInstanceCount int `json:"total_health_instance_count"`
	TotalInstanceCount       int `json:"total_instance_count"`
}

// PolarisService
type PolarisService struct{}

// PolarisInstance
type PolarisInstance struct{}

// PolarisRouting
type PolarisRouting struct{}

// PolarisLimit
type PolarisLimit struct{}

// PolarisUser
type PolarisUser struct{}

// PolarisUserGroup
type PolarisUserGroup struct{}

// PolarisAuthStrategies
type PolarisAuthStrategies struct{}

// PolarisClient
type PolarisClient struct{}

// PolarisData
type PolarisData struct{}

// CheckRes 检查批量执行的结果，统计失败和成功的操作
func (res *BatchResult) CheckRes() {
	res.FailedSize = 0
	res.SuccSize = 0
	for _, r := range res.Responses {
		if r.Code == 200000 {
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
