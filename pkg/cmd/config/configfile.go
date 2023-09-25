package config

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// NewCmdfile build file root cmd
func NewCmdConfigfile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "file [bygroup|search|export|import|create|createandpub|delete|update]",
		Short: "file [bygroup|search|export|import|create|createandpub|delete|update]",
		Long:  "file [bygroup|search|export|import|create|createandpub|delete|update]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	// query command
	cmd.AddCommand(NewCmdfileBygroup())
	cmd.AddCommand(NewCmdfileSearch())

	// write command
	cmd.AddCommand(NewCmdfileExport())
	cmd.AddCommand(NewCmdfileImport())
	cmd.AddCommand(NewCmdfileCreate())
	cmd.AddCommand(NewCmdfileCreateandpub())
	cmd.AddCommand(NewCmdfileDelete())
	cmd.AddCommand(NewCmdfileUpdate())

	return cmd
}

// list param, eg: limit, offset
var bygroupParam entity.QueryParam
var bygroupfileQueryParam entity.ConfigfileQueryParam

// NewCmdfileBygroup build file bygroup command
func NewCmdfileBygroup() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "bygroup file",
		Short: "bygroup file",
		Long:  "bygroup file",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("file", configFields)
			rsRepo := repo.NewResourceRepo(repo.RS_CONFIGFILES, repo.API_CONFIGFILES_BYGROUP)
			rsRepo.Method("GET").Param(bygroupParam.Encode()).Print(print).Build()
		},
	}

	bygroupParam.ResourceParam = &bygroupfileQueryParam
	bygroupParam.RegisterFlag(cmd)
	return cmd
}

// list param, eg: limit, offset
var searchParam entity.QueryParam
var searchfileQueryParam entity.ConfigfileSearchParam

// NewCmdfileSearch build file search command
func NewCmdfileSearch() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "search file",
		Short: "search file",
		Long:  "search file",
		Run: func(cmd *cobra.Command, args []string) {
			print := entity.NewPolarisPrint().ResourceConf("file", configFields)
			rsRepo := repo.NewResourceRepo(repo.RS_CONFIGFILES, repo.API_CONFIGFILES_SEARCH)
			rsRepo.Method("GET").Param(searchParam.Encode()).Print(print).Build()
		},
	}

	searchParam.ResourceParam = &searchfileQueryParam
	searchParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdfileExport build file export command
func NewCmdfileExport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export file",
		Short: "export (-f export_file.json)",
		Long:  "export (-f export_file.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_CONFIGFILES, repo.API_CONFIGFILES_EXPORT)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdfileImport build file import command
func NewCmdfileImport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import file",
		Short: "import (-f import_file.json)",
		Long:  "import (-f import_file.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_CONFIGFILES, repo.API_CONFIGFILES_IMPORT)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdfileCreate build file create command
func NewCmdfileCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create file",
		Short: "create (-f create_file.json)",
		Long:  "create (-f create_file.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_CONFIGFILES, repo.API_CONFIGFILES)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdfileCreateandpub build file createandpub command
func NewCmdfileCreateandpub() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "createandpub file",
		Short: "createandpub (-f createandpub_file.json)",
		Long:  "createandpub (-f createandpub_file.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_CONFIGFILES, repo.API_CONFIGFILES_PUB)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdfileDelete build file delete command
func NewCmdfileDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete file",
		Short: "delete (-f delete_file.json)",
		Long:  "delete (-f delete_file.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_CONFIGFILES, repo.API_CONFIGFILES_DEL)
			rsRepo.Method("DEL").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdfileUpdate build file update command
func NewCmdfileUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update file",
		Short: "update (-f update_file.json)",
		Long:  "update (-f update_file.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_CONFIGFILES, repo.API_CONFIGFILES)
			rsRepo.Method("PUT").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}
