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
package instances

import (
	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

	"github.com/spf13/cobra"
)

// NewCmdHost build host root cmd
func NewCmdInstancesHost() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "host [delete|isolate]",
		Short: "host [delete|isolate]",
		Long:  "host [delete|isolate]",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&resourceFile, "file", "f", "", "json file for delete/isolate host")
	cmd.PersistentFlags().StringVar(&resourceFields, "print", "", "host print field,eg:\"jsontag1,jsontag2\"")

	// query command

	// write command
	cmd.AddCommand(NewCmdHostDelete())
	cmd.AddCommand(NewCmdHostIsolate())

	return cmd
}

// NewCmdHostDelete build host delete command
func NewCmdHostDelete() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete host",
		Short: "delete (-f delete_host.json)",
		Long:  "delete (-f delete_host.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_INSTANCES_HOST_DEL,
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

// NewCmdHostIsolate build host isolate command
func NewCmdHostIsolate() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "isolate host",
		Short: "isolate (-f isolate_host.json)",
		Long:  "isolate (-f isolate_host.json)",
		Run: func(cmd *cobra.Command, args []string) {
			rsRepo := repo.NewResourceRepo(
				repo.API_INSTANCES_HOST_ISOLATE,
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
