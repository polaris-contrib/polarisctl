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
package maintain_test

import (
	"bytes"

	"github.com/glycerine/goconvey/convey"
	"github.com/polaris-contrilb/polarisctl/pkg/cmd/maintain"

	"github.com/polaris-contrilb/polarisctl/pkg/entity"
	"github.com/polaris-contrilb/polarisctl/pkg/repo"

	"testing"
)

func TestMaintainLoglevel(t *testing.T) {

	cmd := maintain.NewCmdMaintain()

	convey.Convey("MaintainLoglevel", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"loglevel"}

		var err error

		args = append(args, "-l")
		args = append(args, "1")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestMaintainSetloglevel(t *testing.T) {

	cmd := maintain.NewCmdMaintain()

	convey.Convey("MaintainSetloglevel", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"setloglevel"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/maintain/setloglevel.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestMaintainLeaders(t *testing.T) {

	cmd := maintain.NewCmdMaintain()

	convey.Convey("MaintainLeaders", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"leaders"}

		var err error

		args = append(args, "-l")
		args = append(args, "1")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestMaintainCmdb(t *testing.T) {

	cmd := maintain.NewCmdMaintain()

	convey.Convey("MaintainCmdb", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"cmdb"}

		var err error

		args = append(args, "-l")
		args = append(args, "1")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestMaintainClients(t *testing.T) {

	cmd := maintain.NewCmdMaintain()

	convey.Convey("MaintainClients", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"clients"}

		var err error

		args = append(args, "-l")
		args = append(args, "1")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func init() {
	repo.RegisterCluster(entity.PolarisClusterConf{
		Name:  "test",
		Host:  "119.91.66.223:8090",
		Token: "GHSNvvpRAKMgGEa8zOKDbEfz3FvgO28yIV01kZRFX1btX1jKFnOZLNxl",
	})
}
