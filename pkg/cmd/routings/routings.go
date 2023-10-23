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
package routings

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

	"github.com/spf13/cobra"
)

// resourceFile resource:routings description file(format:josn) for create/delete/update
var resourceFile string
var resourceFields string

// NewCmdRoutings build routings root cmd
func NewCmdRoutings() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "routings [list|create|update|enable|delete]",
		Short: "routings [list|create|update|enable|delete]",
		Long:  "routings [list|create|update|enable|delete]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&resourceFile, "file", "f", "", "json file for create/update/enable/delete routings")
	cmd.PersistentFlags().StringVar(&resourceFields, "print", "", "routings print field,eg:\"jsontag1,jsontag2\"")

	// query command
	cmd.AddCommand(NewCmdRoutingsList())

	// write command
	cmd.AddCommand(NewCmdRoutingsCreate())
	cmd.AddCommand(NewCmdRoutingsUpdate())
	cmd.AddCommand(NewCmdRoutingsEnable())
	cmd.AddCommand(NewCmdRoutingsDelete())

	return cmd
}

// list param, eg: limit, offset
var listRoutingsParam entity.QueryParam
var listRoutingsQueryParam entity.RoutingsQueryParam

// NewCmdRoutingsList build routings list command
func NewCmdRoutingsList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list routings",
		Short: "list routings",
		Long:  "list routings",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_ROUTINGS,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.BatchQueryResponse")),

				repo.WithParam(listRoutingsParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	listRoutingsParam.ResourceParam = &listRoutingsQueryParam
	listRoutingsParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdRoutingsCreate build routings create command
func NewCmdRoutingsCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create routings",
		Short: "create (-f create_routings.json)",
		Long:  "create (-f create_routings.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_ROUTINGS,
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

// NewCmdRoutingsUpdate build routings update command
func NewCmdRoutingsUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update routings",
		Short: "update (-f update_routings.json)",
		Long:  "update (-f update_routings.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_ROUTINGS,
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

// NewCmdRoutingsEnable build routings enable command
func NewCmdRoutingsEnable() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable routings",
		Short: "enable (-f enable_routings.json)",
		Long:  "enable (-f enable_routings.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_ROUTINGS_ENABLE,
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

// NewCmdRoutingsDelete build routings delete command
func NewCmdRoutingsDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete routings",
		Short: "delete (-f delete_routings.json)",
		Long:  "delete (-f delete_routings.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_ROUTINGS_DEL,
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
