package instances

import (
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// NewCmdInstanceHost build instances host root cmd
func NewCmdInstanceHost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "host [delete|isolate] instances",
		Short: "host [delete|isolate]",
		Long:  "host [delete|isolate]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}

	cmd.AddCommand(NewCmdInstanceHostDelete())
	cmd.AddCommand(NewCmdInstancesHostIsolate())
	return cmd
}

// NewCmdInstanceHostDelete build instance host delete command
func NewCmdInstanceHostDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete instances with host",
		Short: "delete (-f create_instances_host.json)",
		Long:  "delete (-f create_instances_host.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_INSTANCES, repo.API_INSTANCES_HOST_DEL)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdInstanceHostDelete build instance host isolate command
func NewCmdInstancesHostIsolate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "isolate instances with host",
		Short: "isolate (-f isolate_instances.json)",
		Long:  "isolate (-f isolate_instances.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_INSTANCES, repo.API_INSTANCES_HOST_ISOLATE)
			rsRepo.Method("PUT").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}
