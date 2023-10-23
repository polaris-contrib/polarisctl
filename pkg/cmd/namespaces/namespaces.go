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
package namespaces

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

	"github.com/spf13/cobra"
)

// resourceFile resource:namespaces description file(format:josn) for create/delete/update
var resourceFile string
var resourceFields string

// NewCmdNamespaces build namespaces root cmd
func NewCmdNamespaces() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "namespaces [list|create|delete|update]",
		Short: "namespaces [list|create|delete|update]",
		Long:  "namespaces [list|create|delete|update]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&resourceFile, "file", "f", "", "json file for create/delete/update namespaces")
	cmd.PersistentFlags().StringVar(&resourceFields, "print", "", "namespaces print field,eg:\"jsontag1,jsontag2\"")

	// query command
	cmd.AddCommand(NewCmdNamespacesList())

	// write command
	cmd.AddCommand(NewCmdNamespacesCreate())
	cmd.AddCommand(NewCmdNamespacesDelete())
	cmd.AddCommand(NewCmdNamespacesUpdate())

	return cmd
}

// list param, eg: limit, offset
var listNamespacesParam entity.QueryParam
var listNamespacesQueryParam entity.NamespacesQueryParam

// NewCmdNamespacesList build namespaces list command
func NewCmdNamespacesList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list namespaces",
		Short: "list namespaces",
		Long:  "list namespaces",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_NAMESPACES,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.BatchQueryResponse")),

				repo.WithParam(listNamespacesParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	listNamespacesParam.ResourceParam = &listNamespacesQueryParam
	listNamespacesParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdNamespacesCreate build namespaces create command
func NewCmdNamespacesCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create namespaces",
		Short: "create (-f create_namespaces.json)",
		Long:  "create (-f create_namespaces.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_NAMESPACES,
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

// NewCmdNamespacesDelete build namespaces delete command
func NewCmdNamespacesDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete namespaces",
		Short: "delete (-f delete_namespaces.json)",
		Long:  "delete (-f delete_namespaces.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_NAMESPACES_DEL,
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

// NewCmdNamespacesUpdate build namespaces update command
func NewCmdNamespacesUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update namespaces",
		Short: "update (-f update_namespaces.json)",
		Long:  "update (-f update_namespaces.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_NAMESPACES,
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
