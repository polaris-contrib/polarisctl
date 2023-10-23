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
package ratelimits

import (
	"github.com/polaris-contrilb/polarisctl/pkg/entity"
	"github.com/polaris-contrilb/polarisctl/pkg/repo"

	"github.com/spf13/cobra"
)

// resourceFile resource:ratelimits description file(format:josn) for create/delete/update
var resourceFile string
var resourceFields string

// NewCmdRatelimits build ratelimits root cmd
func NewCmdRatelimits() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "ratelimits [list|create|update|enable|delete]",
		Short: "ratelimits [list|create|update|enable|delete]",
		Long:  "ratelimits [list|create|update|enable|delete]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&resourceFile, "file", "f", "", "json file for create/update/enable/delete ratelimits")
	cmd.PersistentFlags().StringVar(&resourceFields, "print", "", "ratelimits print field,eg:\"jsontag1,jsontag2\"")

	// query command
	cmd.AddCommand(NewCmdRatelimitsList())

	// write command
	cmd.AddCommand(NewCmdRatelimitsCreate())
	cmd.AddCommand(NewCmdRatelimitsUpdate())
	cmd.AddCommand(NewCmdRatelimitsEnable())
	cmd.AddCommand(NewCmdRatelimitsDelete())

	return cmd
}

// list param, eg: limit, offset
var listRatelimitsParam entity.QueryParam
var listRatelimitsQueryParam entity.RatelimitsQueryParam

// NewCmdRatelimitsList build ratelimits list command
func NewCmdRatelimitsList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list ratelimits",
		Short: "list ratelimits",
		Long:  "list ratelimits",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_RATELIMITS,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.BatchQueryResponse")),

				repo.WithParam(listRatelimitsParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	listRatelimitsParam.ResourceParam = &listRatelimitsQueryParam
	listRatelimitsParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdRatelimitsCreate build ratelimits create command
func NewCmdRatelimitsCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create ratelimits",
		Short: "create (-f create_ratelimits.json)",
		Long:  "create (-f create_ratelimits.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_RATELIMITS,
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

// NewCmdRatelimitsUpdate build ratelimits update command
func NewCmdRatelimitsUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update ratelimits",
		Short: "update (-f update_ratelimits.json)",
		Long:  "update (-f update_ratelimits.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_RATELIMITS,
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

// NewCmdRatelimitsEnable build ratelimits enable command
func NewCmdRatelimitsEnable() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable ratelimits",
		Short: "enable (-f enable_ratelimits.json)",
		Long:  "enable (-f enable_ratelimits.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_RATELIMITS_ENABLE,
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

// NewCmdRatelimitsDelete build ratelimits delete command
func NewCmdRatelimitsDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete ratelimits",
		Short: "delete (-f delete_ratelimits.json)",
		Long:  "delete (-f delete_ratelimits.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_RATELIMITS_DEL,
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
