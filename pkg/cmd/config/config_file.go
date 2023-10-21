package config

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

	"github.com/spf13/cobra"
)

// NewCmdFile build file root cmd
func NewCmdConfigFile() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "file [bygroup|search|export|import|create|createandpub|delete|update]",
		Short: "file [bygroup|search|export|import|create|createandpub|delete|update]",
		Long:  "file [bygroup|search|export|import|create|createandpub|delete|update]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&resourceFile, "file", "f", "", "json file for export/import/create/createandpub/delete/update file")
	cmd.PersistentFlags().StringVar(&resourceFields, "print", "", "file print field,eg:\"jsontag1,jsontag2\"")

	// query command
	cmd.AddCommand(NewCmdFileBygroup())
	cmd.AddCommand(NewCmdFileSearch())

	// write command
	cmd.AddCommand(NewCmdFileExport())
	cmd.AddCommand(NewCmdFileImport())
	cmd.AddCommand(NewCmdFileCreate())
	cmd.AddCommand(NewCmdFileCreateandpub())
	cmd.AddCommand(NewCmdFileDelete())
	cmd.AddCommand(NewCmdFileUpdate())

	return cmd
}

// list param, eg: limit, offset
var bygroupFileParam entity.QueryParam
var bygroupFileQueryParam entity.ConfigFileQueryParam

// NewCmdFileBygroup build file bygroup command
func NewCmdFileBygroup() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "bygroup file",
		Short: "bygroup file",
		Long:  "bygroup file",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CONFIGFILES_BYGROUP,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.ConfigBatchQueryResponse")),

				repo.WithParam(bygroupFileParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	bygroupFileParam.ResourceParam = &bygroupFileQueryParam
	bygroupFileParam.RegisterFlag(cmd)
	return cmd
}

// list param, eg: limit, offset
var searchFileParam entity.QueryParam
var searchFileQueryParam entity.ConfigFileQueryParam

// NewCmdFileSearch build file search command
func NewCmdFileSearch() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "search file",
		Short: "search file",
		Long:  "search file",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CONFIGFILES_SEARCH,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.ConfigBatchQueryResponse")),

				repo.WithParam(searchFileParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	searchFileParam.ResourceParam = &searchFileQueryParam
	searchFileParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdFileExport build file export command
func NewCmdFileExport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "export file",
		Short: "export (-f export_file.json)",
		Long:  "export (-f export_file.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CONFIGFILES_EXPORT,
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

// NewCmdFileImport build file import command
func NewCmdFileImport() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "import file",
		Short: "import (-f import_file.json)",
		Long:  "import (-f import_file.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CONFIGFILES_IMPORT,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.Response")),

				repo.WithFile(resourceFile),
				repo.WithMethod("POST"))
			rsRepo.Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdFileCreate build file create command
func NewCmdFileCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create file",
		Short: "create (-f create_file.json)",
		Long:  "create (-f create_file.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CONFIGFILES,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.ConfigResponse")),

				repo.WithFile(resourceFile),
				repo.WithMethod("POST"))
			rsRepo.Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdFileCreateandpub build file createandpub command
func NewCmdFileCreateandpub() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "createandpub file",
		Short: "createandpub (-f createandpub_file.json)",
		Long:  "createandpub (-f createandpub_file.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CONFIGFILES_PUB,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.ConfigResponse")),

				repo.WithFile(resourceFile),
				repo.WithMethod("POST"))
			rsRepo.Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdFileDelete build file delete command
func NewCmdFileDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete file",
		Short: "delete (-f delete_file.json)",
		Long:  "delete (-f delete_file.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CONFIGFILES_DEL,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.ConfigResponse")),

				repo.WithFile(resourceFile),
				repo.WithMethod("POST"))
			rsRepo.Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdFileUpdate build file update command
func NewCmdFileUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update file",
		Short: "update (-f update_file.json)",
		Long:  "update (-f update_file.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CONFIGFILES,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.ConfigResponse")),

				repo.WithFile(resourceFile),
				repo.WithMethod("PUT"))
			rsRepo.Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}
