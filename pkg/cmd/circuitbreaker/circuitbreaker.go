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
package circuitbreaker

import (
	"github.com/polaris-contrilb/polarisctl/pkg/entity"
	"github.com/polaris-contrilb/polarisctl/pkg/repo"

	"github.com/spf13/cobra"
)

// resourceFile resource:circuitbreaker description file(format:josn) for create/delete/update
var resourceFile string
var resourceFields string

// NewCmdCircuitbreaker build circuitbreaker root cmd
func NewCmdCircuitbreaker() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "circuitbreaker [list|create|delete|enable|update]",
		Short: "circuitbreaker [list|create|delete|enable|update]",
		Long:  "circuitbreaker [list|create|delete|enable|update]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&resourceFile, "file", "f", "", "json file for create/delete/enable/update circuitbreaker")
	cmd.PersistentFlags().StringVar(&resourceFields, "print", "", "circuitbreaker print field,eg:\"jsontag1,jsontag2\"")

	// query command
	cmd.AddCommand(NewCmdCircuitbreakerList())

	// write command
	cmd.AddCommand(NewCmdCircuitbreakerCreate())
	cmd.AddCommand(NewCmdCircuitbreakerDelete())
	cmd.AddCommand(NewCmdCircuitbreakerEnable())
	cmd.AddCommand(NewCmdCircuitbreakerUpdate())

	return cmd
}

// list param, eg: limit, offset
var listCircuitbreakerParam entity.QueryParam
var listCircuitbreakerQueryParam entity.CircuitbreakerQueryParam

// NewCmdCircuitbreakerList build circuitbreaker list command
func NewCmdCircuitbreakerList() *cobra.Command {
	cmd := &cobra.Command{

		Use:   "list circuitbreaker",
		Short: "list circuitbreaker",
		Long:  "list circuitbreaker",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CIRCUITBREAKER,
				repo.WithWriter(entity.NewTableWriter(entity.WithTags(resourceFields))),
				repo.WithParser(entity.NewResponseParse("v1.BatchQueryResponse")),

				repo.WithParam(listCircuitbreakerParam.Encode()),
				repo.WithMethod("GET"))
			rsRepo.Build()
		},
	}

	listCircuitbreakerParam.ResourceParam = &listCircuitbreakerQueryParam
	listCircuitbreakerParam.RegisterFlag(cmd)
	return cmd
}

// NewCmdCircuitbreakerCreate build circuitbreaker create command
func NewCmdCircuitbreakerCreate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create circuitbreaker",
		Short: "create (-f create_circuitbreaker.json)",
		Long:  "create (-f create_circuitbreaker.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CIRCUITBREAKER,
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

// NewCmdCircuitbreakerDelete build circuitbreaker delete command
func NewCmdCircuitbreakerDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete circuitbreaker",
		Short: "delete (-f delete_circuitbreaker.json)",
		Long:  "delete (-f delete_circuitbreaker.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CIRCUITBREAKER_DEL,
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

// NewCmdCircuitbreakerEnable build circuitbreaker enable command
func NewCmdCircuitbreakerEnable() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enable circuitbreaker",
		Short: "enable (-f enable_circuitbreaker.json)",
		Long:  "enable (-f enable_circuitbreaker.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CIRCUITBREAKER_ENABLE,
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

// NewCmdCircuitbreakerUpdate build circuitbreaker update command
func NewCmdCircuitbreakerUpdate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update circuitbreaker",
		Short: "update (-f update_circuitbreaker.json)",
		Long:  "update (-f update_circuitbreaker.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_CIRCUITBREAKER,
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
