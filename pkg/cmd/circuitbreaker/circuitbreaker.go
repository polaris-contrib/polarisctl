package circuitbreaker

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// fileName resource:circuitbreaker description file(format:josn) for create/delete/update
var fileName string
var circuitbreakerFields string

// NewCmdCircuitbreaker build circuitbreaker root cmd
func NewCmdCircuitbreaker() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "circuitbreaker [list|create|delete|enable|update]",
		Short: "circuitbreaker [list|create|delete|enable|update]",
		Long:  "circuitbreaker [list|create|delete|enable|update]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&fileName, "file", "f", "", "json file for create/delete/enable/update circuitbreaker")
	cmd.PersistentFlags().StringVar(&circuitbreakerFields, "print", "", "circuitbreaker print field")

	// query command
	cmd.AddCommand(NewCmdCircuitbreakerList())

	// write command
	cmd.AddCommand(NewCmdCircuitbreakerCreate())
	cmd.AddCommand(NewCmdCircuitbreakerDelete())
	cmd.AddCommand(NewCmdCircuitbreakerEnable())
	cmd.AddCommand(NewCmdCircuitbreakerUpdate())

	return cmd
}

// list param, eg: limit, offset
var listParam entity.QueryParam
var listCircuitbreakerQueryParam entity.CircuitbreakerQueryParam

// NewCmdCircuitbreakerList build circuitbreaker list command
func NewCmdCircuitbreakerList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list circuitbreaker",
		Short: "list circuitbreaker",
		Long:  "list circuitbreaker",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("v1.CircuitBreakerRule", circuitbreakerFields).V2Api("v1.CircuitBreakerRule")
			rsRepo := repo.NewResourceRepo(repo.RS_CIRCUITBREAKER, repo.API_CIRCUITBREAKER)
			rsRepo.Method("GET").Param(listParam.Encode()).Print(print).Build()
		},
	}

	listParam.ResourceParam = &listCircuitbreakerQueryParam
	listParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdCircuitbreakerCreate build circuitbreaker create command
func NewCmdCircuitbreakerCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create circuitbreaker",
		Short: "create (-f create_circuitbreaker.json)",
		Long:  "create (-f create_circuitbreaker.json)",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("v1.CircuitBreakerRule", circuitbreakerFields).V2Api("v1.CircuitBreakerRule")
			rsRepo := repo.NewResourceRepo(repo.RS_CIRCUITBREAKER, repo.API_CIRCUITBREAKER).Print(print)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdCircuitbreakerDelete build circuitbreaker delete command
func NewCmdCircuitbreakerDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete circuitbreaker",
		Short: "delete (-f delete_circuitbreaker.json)",
		Long:  "delete (-f delete_circuitbreaker.json)",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("v1.CircuitBreakerRule", circuitbreakerFields).V2Api("v1.CircuitBreakerRule")
			rsRepo := repo.NewResourceRepo(repo.RS_CIRCUITBREAKER, repo.API_CIRCUITBREAKER_DEL).Print(print)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdCircuitbreakerEnable build circuitbreaker enable command
func NewCmdCircuitbreakerEnable() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable circuitbreaker",
		Short: "enable (-f enable_circuitbreaker.json)",
		Long:  "enable (-f enable_circuitbreaker.json)",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("v1.CircuitBreakerRule", circuitbreakerFields).V2Api("v1.CircuitBreakerRule")
			rsRepo := repo.NewResourceRepo(repo.RS_CIRCUITBREAKER, repo.API_CIRCUITBREAKER_ENABLE).Print(print)
			rsRepo.Method("PUT").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdCircuitbreakerUpdate build circuitbreaker update command
func NewCmdCircuitbreakerUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update circuitbreaker",
		Short: "update (-f update_circuitbreaker.json)",
		Long:  "update (-f update_circuitbreaker.json)",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("v1.CircuitBreakerRule", circuitbreakerFields).V2Api("v1.CircuitBreakerRule")
			rsRepo := repo.NewResourceRepo(repo.RS_CIRCUITBREAKER, repo.API_CIRCUITBREAKER).Print(print)
			rsRepo.Method("PUT").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}
