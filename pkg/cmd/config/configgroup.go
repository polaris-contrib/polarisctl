package config

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// NewCmdgroup build group root cmd
func NewCmdConfiggroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group [list|create|delete|update]",
		Short: "group [list|create|delete|update]",
		Long:  "group [list|create|delete|update]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	// query command
	cmd.AddCommand(NewCmdgroupList())

	// write command
	cmd.AddCommand(NewCmdgroupCreate())
	cmd.AddCommand(NewCmdgroupDelete())
	cmd.AddCommand(NewCmdgroupUpdate())

	return cmd
}

// list param, eg: limit, offset
var listgroupParam entity.QueryParam
var listgroupQueryParam entity.ConfiggroupQueryParam

// NewCmdgroupList build group list command
func NewCmdgroupList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list group",
		Short: "list group",
		Long:  "list group",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("group", configFields)
			rsRepo := repo.NewResourceRepo(repo.RS_CONFIGGROUP, repo.API_CONFIGGROUPS)
			rsRepo.Method("GET").Param(listgroupParam.Encode()).Print(print).Build()
		},
	}

	listgroupParam.ResourceParam = &listgroupQueryParam
	listgroupParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdgroupCreate build group create command
func NewCmdgroupCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create group",
		Short: "create (-f create_group.json)",
		Long:  "create (-f create_group.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_CONFIGGROUP, repo.API_CONFIGGROUPS)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdgroupDelete build group delete command
func NewCmdgroupDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete group",
		Short: "delete (-f delete_group.json)",
		Long:  "delete (-f delete_group.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_CONFIGGROUP, repo.API_CONFIGGROUPS_DEL)
			rsRepo.Method("DEL").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdgroupUpdate build group update command
func NewCmdgroupUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update group",
		Short: "update (-f update_group.json)",
		Long:  "update (-f update_group.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_CONFIGGROUP, repo.API_CONFIGGROUPS)
			rsRepo.Method("PUT").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}
