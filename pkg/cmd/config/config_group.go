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
package config

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

	"github.com/spf13/cobra"
)

// NewCmdGroup build group root cmd
func NewCmdConfigGroup() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "group [list|create|delete|update]",
		Short: "group [list|create|delete|update]",
		Long:  "group [list|create|delete|update]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&resourceFile, "file", "f", "", "json file for create/delete/update group")
	cmd.PersistentFlags().StringVar(&resourceFields, "print", "", "group print field,eg:\"jsontag1,jsontag2\"")

	// query command
	cmd.AddCommand(NewCmdGroupList())

	// write command
	cmd.AddCommand(NewCmdGroupCreate())
	cmd.AddCommand(NewCmdGroupDelete())
	cmd.AddCommand(NewCmdGroupUpdate())

	return cmd
}

// list param, eg: limit, offset
var listGroupParam entity.QueryParam
var listGroupQueryParam entity.ConfigGroupQueryParam

// NewCmdGroupList build group list command
func NewCmdGroupList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list group",
		Short: "list group",
		Long:  "list group",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CONFIGGROUPS,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.ConfigBatchQueryResponse")),

				repo.WithParam(listGroupParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	listGroupParam.ResourceParam = &listGroupQueryParam
	listGroupParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdGroupCreate build group create command
func NewCmdGroupCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create group",
		Short: "create (-f create_group.json)",
		Long:  "create (-f create_group.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CONFIGGROUPS,
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

// NewCmdGroupDelete build group delete command
func NewCmdGroupDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete group",
		Short: "delete (-f delete_group.json)",
		Long:  "delete (-f delete_group.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CONFIGGROUPS_DEL,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.ConfigResponse")),

				repo.WithFile(resourceFile),
				repo.WithMethod("DELETE"))
			rsRepo.Build()
		},
	}

	cmd.MarkFlagRequired("file")
	return cmd
}

// NewCmdGroupUpdate build group update command
func NewCmdGroupUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update group",
		Short: "update (-f update_group.json)",
		Long:  "update (-f update_group.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CONFIGGROUPS,
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
