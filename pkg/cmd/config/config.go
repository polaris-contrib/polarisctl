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
	"github.com/spf13/cobra"
)

// resourceFile resource:config description file(format:josn) for create/delete/update
var resourceFile string
var resourceFields string

// NewCmdConfig build config root cmd
func NewCmdConfig() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config []",
		Short: "config []",
		Long:  "config []",
		Run:   func(cmd *cobra.Command, args []string) { cmd.Help() },
	}
	cmd.PersistentFlags().StringVarP(&resourceFile, "file", "f", "", "json file for  config")
	cmd.PersistentFlags().StringVar(&resourceFields, "print", "", "config print field,eg:\"jsontag1,jsontag2\"")

	// query command

	// write command
	cmd.AddCommand(NewCmdConfigGroup())
	cmd.AddCommand(NewCmdConfigFile())
	cmd.AddCommand(NewCmdConfigRelease())

	return cmd
}
