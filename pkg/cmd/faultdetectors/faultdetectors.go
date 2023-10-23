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
package faultdetectors

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

	"github.com/spf13/cobra"
)

// resourceFile resource:faultdetectors description file(format:josn) for create/delete/update
var resourceFile string
var resourceFields string

// NewCmdFaultdetectors build faultdetectors root cmd
func NewCmdFaultdetectors() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "faultdetectors [list|create|delete|update]",
		Short: "faultdetectors [list|create|delete|update]",
		Long:  "faultdetectors [list|create|delete|update]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&resourceFile, "file", "f", "", "json file for create/delete/update faultdetectors")
	cmd.PersistentFlags().StringVar(&resourceFields, "print", "", "faultdetectors print field,eg:\"jsontag1,jsontag2\"")

	// query command
	cmd.AddCommand(NewCmdFaultdetectorsList())

	// write command
	cmd.AddCommand(NewCmdFaultdetectorsCreate())
	cmd.AddCommand(NewCmdFaultdetectorsDelete())
	cmd.AddCommand(NewCmdFaultdetectorsUpdate())

	return cmd
}

// list param, eg: limit, offset
var listFaultdetectorsParam entity.QueryParam
var listFaultdetectorsQueryParam entity.FaultdetectorsQueryParam

// NewCmdFaultdetectorsList build faultdetectors list command
func NewCmdFaultdetectorsList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list faultdetectors",
		Short: "list faultdetectors",
		Long:  "list faultdetectors",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_FAULTDETECTORS,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.BatchQueryResponse")),

				repo.WithParam(listFaultdetectorsParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	listFaultdetectorsParam.ResourceParam = &listFaultdetectorsQueryParam
	listFaultdetectorsParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdFaultdetectorsCreate build faultdetectors create command
func NewCmdFaultdetectorsCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create faultdetectors",
		Short: "create (-f create_faultdetectors.json)",
		Long:  "create (-f create_faultdetectors.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_FAULTDETECTORS,
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

// NewCmdFaultdetectorsDelete build faultdetectors delete command
func NewCmdFaultdetectorsDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete faultdetectors",
		Short: "delete (-f delete_faultdetectors.json)",
		Long:  "delete (-f delete_faultdetectors.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_FAULTDETECTORS_DEL,
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

// NewCmdFaultdetectorsUpdate build faultdetectors update command
func NewCmdFaultdetectorsUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update faultdetectors",
		Short: "update (-f update_faultdetectors.json)",
		Long:  "update (-f update_faultdetectors.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_FAULTDETECTORS,
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
