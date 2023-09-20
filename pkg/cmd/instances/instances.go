package instances

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// fileName resource:instances description file(format:josn) for create/delete/update
var fileName string
var instancesFields string

// NewCmdInstances build instances root cmd
func NewCmdInstances() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "instances [list|labels|create|delete|count|update]",
		Short: "instances [list|labels|create|delete|count|update]",
		Long:  "instances [list|labels|create|delete|count|update]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&fileName, "file", "f", "", "json file for create/delete/update instances")
	cmd.PersistentFlags().StringVar(&instancesFields, "print", "", "instances print field")

	// query command
	cmd.AddCommand(NewCmdInstancesList())
	cmd.AddCommand(NewCmdInstancesLabels())
	cmd.AddCommand(NewCmdInstancesCount())

	// write command
	cmd.AddCommand(NewCmdInstancesCreate())
	cmd.AddCommand(NewCmdInstancesDelete())
	cmd.AddCommand(NewCmdInstancesUpdate())
	cmd.AddCommand(NewCmdInstanceHost())

	return cmd
}

// list param, eg: limit, offset
var listParam entity.QueryParam
var listInstancesQueryParam entity.InstancesQueryParam

// NewCmdInstancesList build instances list command
func NewCmdInstancesList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list instances",
		Short: "list instances",
		Long:  "list instances",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("instances", instancesFields)
			rsRepo := repo.NewResourceRepo(repo.RS_INSTANCES, repo.API_INSTANCES)
			rsRepo.Method("GET").Param(listParam.Encode()).Print(print).Build()
		},
	}

	listParam.ResourceParam = &listInstancesQueryParam
	listParam.RegisterFlag(cmd)
	return cmd
}

var labelsQueryParam entity.LabelQueryParam

// NewCmdInstancesLabels build instances labels command
func NewCmdInstancesLabels() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "labels instances",
		Short: "labels instances",
		Long:  "labels instances",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("instances", instancesFields)
			rsRepo := repo.NewResourceRepo(repo.RS_INSTANCES, repo.API_INSTANCES_LABELS)
			rsRepo.Method("GET").Param(labelsQueryParam.Encode()).Batch(false).Print(print).Build()
		},
	}

	labelsQueryParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdInstancesCreate build instances create command
func NewCmdInstancesCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create instances",
		Short: "create (-f create_instances.json)",
		Long:  "create (-f create_instances.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_INSTANCES, repo.API_INSTANCES)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdInstancesDelete build instances delete command
func NewCmdInstancesDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete instances",
		Short: "delete (-f delete_instances.json)",
		Long:  "delete (-f delete_instances.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_INSTANCES, repo.API_INSTANCES_DEL)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdInstancesCount build instances count command
func NewCmdInstancesCount() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "count instances",
		Short: "count ",
		Long:  "count ",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_INSTANCES, repo.API_INSTANCES_COUNT)
			rsRepo.Method("GET").Build()
		},
	}

	return cmd
}

// NewCmdInstancesUpdate build instances update command
func NewCmdInstancesUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update instances",
		Short: "update (-f update_instances.json)",
		Long:  "update (-f update_instances.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_INSTANCES, repo.API_INSTANCES)
			rsRepo.Method("PUT").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}
