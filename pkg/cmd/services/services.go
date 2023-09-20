package services

import (
	"fmt"
	"net/url"

	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"
	"github.com/spf13/cobra"
)

// fileName resource:services description file(format:josn) for create/delete/update
var fileName string

// resource print conf
var serviceFields string

// NewCmdServices build services root cmd
func NewCmdServices() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "services [list|all|create|update|delete]",
		Short: "services [list|all|create|update|delete]",
		Long:  "services [list|all|create|update|delete]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&fileName, "file", "f", "", "json file for create/update/delete services")
	cmd.PersistentFlags().StringVar(&serviceFields, "print", "", "set services print field,eg:--print=name,namespaces")

	// query command
	cmd.AddCommand(NewCmdServicesList())
	cmd.AddCommand(NewCmdServicesAll())

	// write command
	cmd.AddCommand(NewCmdServicesCreate())
	cmd.AddCommand(NewCmdServicesUpdate())
	cmd.AddCommand(NewCmdServicesDelete())

	return cmd
}

// list param, eg: limit, offset
var listParam entity.QueryParam
var listServicesQueryParam entity.ServicesQueryParam

// NewCmdServicesList build services list command
func NewCmdServicesList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list services",
		Short: "list services",
		Long:  "list services",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_SERVICES, repo.API_SERVICES)
			print := entity.NewPolarisPrint().ResourceConf("services", serviceFields)
			rsRepo.Method("GET").Param(listParam.Encode()).Print(print).Build()
		},
	}

	listParam.ResourceParam = &listServicesQueryParam
	listParam.RegisterFlag(cmd)
	return cmd
}

var namespace string

// NewCmdServicesAll build services all command
func NewCmdServicesAll() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "all --namespace=xx",
		Short: "all services names",
		Long:  "all services names",
		Run: func(cmd *cobra.Command, args []string) {
			param := url.Values{}
			if 0 != len(namespace) {
				param.Add("namespace", namespace)
			}
			fmt.Println(len(namespace))
			rsRepo := repo.NewResourceRepo(repo.RS_SERVICES, repo.API_SERVICESALL)
			rsRepo.Method("GET").Param(param.Encode()).Build()
		},
	}
	cmd.Flags().StringVar(&namespace, "namespace", "", "get namespace all services name")
	return cmd
}

// NewCmdServicesCreate build services create command
func NewCmdServicesCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create services",
		Short: "create (-f create_services.json)",
		Long:  "create (-f create_services.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_SERVICES, repo.API_SERVICES)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdServicesUpdate build services update command
func NewCmdServicesUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update services",
		Short: "update (-f update_services.json)",
		Long:  "update (-f update_services.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_SERVICES, repo.API_SERVICES)
			rsRepo.Method("PUT").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdServicesDelete build services delete command
func NewCmdServicesDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete services",
		Short: "delete (-f delete_services.json)",
		Long:  "delete (-f delete_services.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(repo.RS_SERVICES, repo.API_SERVICESDEL)
			rsRepo.Method("POST").File(fileName).Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}
