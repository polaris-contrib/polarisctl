package maintain

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// fileName resource:maintain description file(format:josn) for create/delete/update
var fileName string
var maintainFields string

// NewCmdMaintain build maintain root cmd
func NewCmdMaintain() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "maintain [loglevel|setloglevel|leaders|cmdb|clients]",
		Short: "maintain [loglevel|setloglevel|leaders|cmdb|clients]",
		Long:  "maintain [loglevel|setloglevel|leaders|cmdb|clients]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&fileName, "file", "f", "", "json file for setloglevel maintain")
	cmd.PersistentFlags().StringVar(&maintainFields, "print", "", "maintain print field")

	// query command
	cmd.AddCommand(NewCmdMaintainLoglevel())
	cmd.AddCommand(NewCmdMaintainLeaders())
	cmd.AddCommand(NewCmdMaintainCmdb())
	cmd.AddCommand(NewCmdMaintainClients())

	// write command
	cmd.AddCommand(NewCmdMaintainSetloglevel())

	return cmd
}

// NewCmdMaintainLoglevel build maintain loglevel command
func NewCmdMaintainLoglevel() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "loglevel maintain",
		Short: "loglevel maintain",
		Long:  "loglevel maintain",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("maintain", maintainFields)
			rsRepo := repo.NewResourceRepo(repo.RS_MAINTAIN, repo.API_MAINTAIN_LOG)
			rsRepo.Method("GET").Print(print).Build()
		},
	}

	return cmd
}

// NewCmdMaintainSetloglevel build maintain setloglevel command
func NewCmdMaintainSetloglevel() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "setloglevel maintain",
		Short: "setloglevel (-f setloglevel_maintain.json)",
		Long:  "setloglevel (-f setloglevel_maintain.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_MAINTAIN, repo.API_MAINTAIN_LOG)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdMaintainLeaders build maintain leaders command
func NewCmdMaintainLeaders() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "leaders maintain",
		Short: "leaders maintain",
		Long:  "leaders maintain",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("maintain", maintainFields)
			rsRepo := repo.NewResourceRepo(repo.RS_MAINTAIN, repo.API_MAINTAIN_LEADER)
			rsRepo.Method("GET").Print(print).Build()
		},
	}

	return cmd
}

// NewCmdMaintainCmdb build maintain cmdb command
func NewCmdMaintainCmdb() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "cmdb maintain",
		Short: "cmdb maintain",
		Long:  "cmdb maintain",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("maintain", maintainFields)
			rsRepo := repo.NewResourceRepo(repo.RS_MAINTAIN, repo.API_MAINTAIN_CMDB)
			rsRepo.Method("GET").Print(print).Build()
		},
	}

	return cmd
}

// NewCmdMaintainClients build maintain clients command
func NewCmdMaintainClients() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "clients maintain",
		Short: "clients maintain",
		Long:  "clients maintain",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("maintain", maintainFields)
			rsRepo := repo.NewResourceRepo(repo.RS_MAINTAIN, repo.API_MAINTAIN_CLIENT)
			rsRepo.Method("GET").Print(print).Build()
		},
	}

	return cmd
}
