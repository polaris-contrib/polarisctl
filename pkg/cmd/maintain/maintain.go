package maintain

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

	"github.com/spf13/cobra"
)

// resourceFile resource:maintain description file(format:josn) for create/delete/update
var resourceFile string
var resourceFields string

// NewCmdMaintain build maintain root cmd
func NewCmdMaintain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "maintain [loglevel|setloglevel|leaders|cmdb|clients]",
		Short: "maintain [loglevel|setloglevel|leaders|cmdb|clients]",
		Long:  "maintain [loglevel|setloglevel|leaders|cmdb|clients]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&resourceFile, "file", "f", "", "json file for setloglevel maintain")
	cmd.PersistentFlags().StringVar(&resourceFields, "print", "", "maintain print field,eg:\"jsontag1,jsontag2\"")

	// query command
	cmd.AddCommand(NewCmdMaintainLoglevel())
	cmd.AddCommand(NewCmdMaintainLeaders())
	cmd.AddCommand(NewCmdMaintainCmdb())
	cmd.AddCommand(NewCmdMaintainClients())

	// write command
	cmd.AddCommand(NewCmdMaintainSetloglevel())

	return cmd
}

// list param, eg: limit, offset
var loglevelMaintainParam entity.QueryParam
var loglevelMaintainQueryParam entity.MaintainQueryParam

// NewCmdMaintainLoglevel build maintain loglevel command
func NewCmdMaintainLoglevel() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "loglevel maintain",
		Short: "loglevel maintain",
		Long:  "loglevel maintain",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_MAINTAIN_LOG,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("Maintain.LogLevel")),

				repo.WithParam(loglevelMaintainParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	loglevelMaintainParam.ResourceParam = &loglevelMaintainQueryParam
	loglevelMaintainParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdMaintainSetloglevel build maintain setloglevel command
func NewCmdMaintainSetloglevel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setloglevel maintain",
		Short: "setloglevel (-f setloglevel_maintain.json)",
		Long:  "setloglevel (-f setloglevel_maintain.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_MAINTAIN_LOG,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("Maintain.LogLevel")),

				repo.WithFile(resourceFile),
				repo.WithMethod("POST"))
			rsRepo.Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// list param, eg: limit, offset
var leadersMaintainParam entity.QueryParam
var leadersMaintainQueryParam entity.MaintainQueryParam

// NewCmdMaintainLeaders build maintain leaders command
func NewCmdMaintainLeaders() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "leaders maintain",
		Short: "leaders maintain",
		Long:  "leaders maintain",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_MAINTAIN_LEADER,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("Maintain.Leader")),

				repo.WithParam(leadersMaintainParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	leadersMaintainParam.ResourceParam = &leadersMaintainQueryParam
	leadersMaintainParam.RegisterFlag(cmd)
	return cmd
}

// list param, eg: limit, offset
var cmdbMaintainParam entity.QueryParam
var cmdbMaintainQueryParam entity.MaintainQueryParam

// NewCmdMaintainCmdb build maintain cmdb command
func NewCmdMaintainCmdb() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "cmdb maintain",
		Short: "cmdb maintain",
		Long:  "cmdb maintain",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_MAINTAIN_CMDB,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("Maintain.CMDB")),

				repo.WithParam(cmdbMaintainParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	cmdbMaintainParam.ResourceParam = &cmdbMaintainQueryParam
	cmdbMaintainParam.RegisterFlag(cmd)
	return cmd
}

// list param, eg: limit, offset
var clientsMaintainParam entity.QueryParam
var clientsMaintainQueryParam entity.MaintainQueryParam

// NewCmdMaintainClients build maintain clients command
func NewCmdMaintainClients() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "clients maintain",
		Short: "clients maintain",
		Long:  "clients maintain",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_MAINTAIN_CLIENT,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("Maintain.Clients")),

				repo.WithParam(clientsMaintainParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	clientsMaintainParam.ResourceParam = &clientsMaintainQueryParam
	clientsMaintainParam.RegisterFlag(cmd)
	return cmd
}
