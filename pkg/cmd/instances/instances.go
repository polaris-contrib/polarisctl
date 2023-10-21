package instances

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

	"github.com/spf13/cobra"
)

// resourceFile resource:instances description file(format:josn) for create/delete/update
var resourceFile string
var resourceFields string

// NewCmdInstances build instances root cmd
func NewCmdInstances() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "instances [list|labels|create|delete|count|update]",
		Short: "instances [list|labels|create|delete|count|update]",
		Long:  "instances [list|labels|create|delete|count|update]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&resourceFile, "file", "f", "", "json file for create/delete/update instances")
	cmd.PersistentFlags().StringVar(&resourceFields, "print", "", "instances print field,eg:\"jsontag1,jsontag2\"")

	// query command
	cmd.AddCommand(NewCmdInstancesList())
	cmd.AddCommand(NewCmdInstancesLabels())
	cmd.AddCommand(NewCmdInstancesCount())

	// write command
	cmd.AddCommand(NewCmdInstancesCreate())
	cmd.AddCommand(NewCmdInstancesDelete())
	cmd.AddCommand(NewCmdInstancesUpdate())
	cmd.AddCommand(NewCmdInstancesHost())

	return cmd
}

// list param, eg: limit, offset
var listInstancesParam entity.QueryParam
var listInstancesQueryParam entity.InstancesQueryParam

// NewCmdInstancesList build instances list command
func NewCmdInstancesList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list instances",
		Short: "list instances",
		Long:  "list instances",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_INSTANCES,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.BatchQueryResponse")),

				repo.WithParam(listInstancesParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	listInstancesParam.ResourceParam = &listInstancesQueryParam
	listInstancesParam.RegisterFlag(cmd)
	return cmd
}

// list param, eg: limit, offset
var labelsInstancesParam entity.QueryParam
var labelsInstancesQueryParam entity.InstancesQueryParam

// NewCmdInstancesLabels build instances labels command
func NewCmdInstancesLabels() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "labels instances",
		Short: "labels instances",
		Long:  "labels instances",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_INSTANCES_LABELS,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.Response")),

				repo.WithParam(labelsInstancesParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	labelsInstancesParam.ResourceParam = &labelsInstancesQueryParam
	labelsInstancesParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdInstancesCreate build instances create command
func NewCmdInstancesCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create instances",
		Short: "create (-f create_instances.json)",
		Long:  "create (-f create_instances.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_INSTANCES,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.BatchWriteResponse")),

				repo.WithFile(resourceFile),
				repo.WithMethod("POST"))
			rsRepo.Build()
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
			rsRepo := repo.NewResourceRepo(
				repo.API_INSTANCES_DEL,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.BatchWriteResponse")),

				repo.WithFile(resourceFile),
				repo.WithMethod("POST"))
			rsRepo.Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// list param, eg: limit, offset
var countInstancesParam entity.QueryParam
var countInstancesQueryParam entity.InstancesQueryParam

// NewCmdInstancesCount build instances count command
func NewCmdInstancesCount() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "count instances",
		Short: "count instances",
		Long:  "count instances",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_INSTANCES_COUNT,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.BatchQueryResponse")),

				repo.WithParam(countInstancesParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	countInstancesParam.ResourceParam = &countInstancesQueryParam
	countInstancesParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdInstancesUpdate build instances update command
func NewCmdInstancesUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update instances",
		Short: "update (-f update_instances.json)",
		Long:  "update (-f update_instances.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_INSTANCES,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.BatchWriteResponse")),

				repo.WithFile(resourceFile),
				repo.WithMethod("PUT"))
			rsRepo.Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}
