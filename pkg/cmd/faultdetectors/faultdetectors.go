package faultdetectors

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// fileName resource:faultdetectors description file(format:josn) for create/delete/update
var fileName string
var faultdetectorsFields string

// NewCmdFaultdetectors build faultdetectors root cmd
func NewCmdFaultdetectors() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "faultdetectors [list|create|delete|update]",
		Short: "faultdetectors [list|create|delete|update]",
		Long:  "faultdetectors [list|create|delete|update]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&fileName, "file", "f", "", "json file for create/delete/update faultdetectors")
	cmd.PersistentFlags().StringVar(&faultdetectorsFields, "print", "", "faultdetectors print field")

	// query command
	cmd.AddCommand(NewCmdFaultdetectorsList())

	// write command
	cmd.AddCommand(NewCmdFaultdetectorsCreate())
	cmd.AddCommand(NewCmdFaultdetectorsDelete())
	cmd.AddCommand(NewCmdFaultdetectorsUpdate())

	return cmd
}

// list param, eg: limit, offset
var listParam entity.QueryParam
var listFaultdetectorsQueryParam entity.FaultdetectorsQueryParam

// NewCmdFaultdetectorsList build faultdetectors list command
func NewCmdFaultdetectorsList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list faultdetectors",
		Short: "list faultdetectors",
		Long:  "list faultdetectors",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("v1.FaultDetectRule", faultdetectorsFields).V2Api("v1.FaultDetectRule")
			rsRepo := repo.NewResourceRepo(repo.RS_FAULTDETECTORS, repo.API_FAULTDETECTORS_DEL).Print(print)
			rsRepo.Method("GET").Param(listParam.Encode()).Print(print).Build()
		},
	}

	listParam.ResourceParam = &listFaultdetectorsQueryParam
	listParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdFaultdetectorsCreate build faultdetectors create command
func NewCmdFaultdetectorsCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create faultdetectors",
		Short: "create (-f create_faultdetectors.json)",
		Long:  "create (-f create_faultdetectors.json)",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("v1.FaultDetectRule", faultdetectorsFields).V2Api("v1.FaultDetectRule")
			rsRepo := repo.NewResourceRepo(repo.RS_FAULTDETECTORS, repo.API_FAULTDETECTORS).Print(print)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdFaultdetectorsDelete build faultdetectors delete command
func NewCmdFaultdetectorsDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete faultdetectors",
		Short: "delete (-f delete_faultdetectors.json)",
		Long:  "delete (-f delete_faultdetectors.json)",
		Run: func(cmd *cobra.Command, args []string) {

			print := entity.NewPolarisPrint().ResourceConf("v1.FaultDetectRule", faultdetectorsFields).V2Api("v1.FaultDetectRule")
			rsRepo := repo.NewResourceRepo(repo.RS_FAULTDETECTORS, repo.API_FAULTDETECTORS_DEL).Print(print)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdFaultdetectorsUpdate build faultdetectors update command
func NewCmdFaultdetectorsUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update faultdetectors",
		Short: "update (-f update_faultdetectors.json)",
		Long:  "update (-f update_faultdetectors.json)",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("v1.FaultDetectRule", faultdetectorsFields).V2Api("v1.FaultDetectRule")
			rsRepo := repo.NewResourceRepo(repo.RS_FAULTDETECTORS, repo.API_FAULTDETECTORS).Print(print)
			rsRepo.Method("PUT").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}
