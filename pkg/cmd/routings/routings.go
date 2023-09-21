package routings

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// fileName resource:routings description file(format:josn) for create/delete/update
var fileName string
var routingsFields string

// NewCmdRoutings build routings root cmd
func NewCmdRoutings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "routings [list|create|update|enable|delete]",
		Short: "routings [list|create|update|enable|delete]",
		Long:  "routings [list|create|update|enable|delete]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&fileName, "file", "f", "", "json file for create/update/enable/delete routings")
	cmd.PersistentFlags().StringVar(&routingsFields, "print", "", "routings print field")

	// query command
	cmd.AddCommand(NewCmdRoutingsList())

	// write command
	cmd.AddCommand(NewCmdRoutingsCreate())
	cmd.AddCommand(NewCmdRoutingsUpdate())
	cmd.AddCommand(NewCmdRoutingsEnable())
	cmd.AddCommand(NewCmdRoutingsDelete())

	return cmd
}

// list param, eg: limit, offset
var listParam entity.QueryParam
var listRoutingsQueryParam entity.RoutingsQueryParam

// NewCmdRoutingsList build routings list command
func NewCmdRoutingsList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list routings",
		Short: "list routings",
		Long:  "list routings",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("v1.RouteRule", routingsFields).V2Api("routings")
			rsRepo := repo.NewResourceRepo(repo.RS_ROUTINGS, repo.API_ROUTINGS)
			rsRepo.Method("GET").Param(listParam.Encode()).Print(print).Build()
		},
	}

	listParam.ResourceParam = &listRoutingsQueryParam
	listParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdRoutingsCreate build routings create command
func NewCmdRoutingsCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create routings",
		Short: "create (-f create_routings.json)",
		Long:  "create (-f create_routings.json)",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("v1.RouteRule", routingsFields).V2Api("v1.RouteRule")
			rsRepo := repo.NewResourceRepo(repo.RS_ROUTINGS, repo.API_ROUTINGS)
			rsRepo.Method("POST").File(fileName).Print(print).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdRoutingsUpdate build routings update command
func NewCmdRoutingsUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update routings",
		Short: "update (-f update_routings.json)",
		Long:  "update (-f update_routings.json)",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("v1.RouteRule", routingsFields).V2Api("v1.RouteRule")
			rsRepo := repo.NewResourceRepo(repo.RS_ROUTINGS, repo.API_ROUTINGS)
			rsRepo.Method("PUT").File(fileName).Print(print).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdRoutingsEnable build routings enable command
func NewCmdRoutingsEnable() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable routings",
		Short: "enable (-f enable_routings.json)",
		Long:  "enable (-f enable_routings.json)",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("v1.RouteRule", routingsFields).V2Api("v1.RouteRule")
			rsRepo := repo.NewResourceRepo(repo.RS_ROUTINGS, repo.API_ROUTINGS_ENABLE)
			rsRepo.Method("PUT").File(fileName).Print(print).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdRoutingsDelete build routings delete command
func NewCmdRoutingsDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete routings",
		Short: "delete (-f delete_routings.json)",
		Long:  "delete (-f delete_routings.json)",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("v1.RouteRule", routingsFields).V2Api("v1.RouteRule")
			rsRepo := repo.NewResourceRepo(repo.RS_ROUTINGS, repo.API_ROUTINGS_DEL)
			rsRepo.Method("POST").File(fileName).Print(print).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}
