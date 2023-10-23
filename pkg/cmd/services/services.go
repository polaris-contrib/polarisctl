/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */
package services

import (
	"github.com/polaris-contrilb/polarisctl/pkg/entity"
	"github.com/polaris-contrilb/polarisctl/pkg/repo"

	"github.com/spf13/cobra"
)

// resourceFile resource:services description file(format:josn) for create/delete/update
var resourceFile string
var resourceFields string

// NewCmdServices build services root cmd
func NewCmdServices() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "services [list|all|create|update|delete]",
		Short: "services [list|all|create|update|delete]",
		Long:  "services [list|all|create|update|delete]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&resourceFile, "file", "f", "", "json file for create/update/delete services")
	cmd.PersistentFlags().StringVar(&resourceFields, "print", "", "services print field,eg:\"jsontag1,jsontag2\"")

	// query command
	cmd.AddCommand(NewCmdServicesList())
	cmd.AddCommand(NewCmdServicesAll())

	// write command
	cmd.AddCommand(NewCmdServicesCreate())
	cmd.AddCommand(NewCmdServicesUpdate())
	cmd.AddCommand(NewCmdServicesDelete())
	cmd.AddCommand(NewCmdServicesAlias())

	return cmd
}

// list param, eg: limit, offset
var listServicesParam entity.QueryParam
var listServicesQueryParam entity.ServicesQueryParam

// NewCmdServicesList build services list command
func NewCmdServicesList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list services",
		Short: "list services",
		Long:  "list services",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_SERVICES,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.BatchQueryResponse")),

				repo.WithParam(listServicesParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	listServicesParam.ResourceParam = &listServicesQueryParam
	listServicesParam.RegisterFlag(cmd)
	return cmd
}

// list param, eg: limit, offset
var allServicesParam entity.QueryParam
var allServicesQueryParam entity.ServicesQueryParam

// NewCmdServicesAll build services all command
func NewCmdServicesAll() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "all services",
		Short: "all services",
		Long:  "all services",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_SERVICES_ALL,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.BatchQueryResponse")),

				repo.WithParam(allServicesParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	allServicesParam.ResourceParam = &allServicesQueryParam
	allServicesParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdServicesCreate build services create command
func NewCmdServicesCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create services",
		Short: "create (-f create_services.json)",
		Long:  "create (-f create_services.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_SERVICES,
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

// NewCmdServicesUpdate build services update command
func NewCmdServicesUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update services",
		Short: "update (-f update_services.json)",
		Long:  "update (-f update_services.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_SERVICES,
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

// NewCmdServicesDelete build services delete command
func NewCmdServicesDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete services",
		Short: "delete (-f delete_services.json)",
		Long:  "delete (-f delete_services.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_SERVICES_DEL,
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
