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
package ratelimits_test

import (
	"bytes"

	"github.com/glycerine/goconvey/convey"
	"github.com/polaris-contrilb/polarisctl/pkg/cmd/ratelimits"

	"github.com/polaris-contrilb/polarisctl/pkg/entity"
	"github.com/polaris-contrilb/polarisctl/pkg/repo"

	"testing"
)

func TestRatelimitsList(t *testing.T) {

	cmd := ratelimits.NewCmdRatelimits()

	convey.Convey("RatelimitsList", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"list"}

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

func TestRatelimitsCreate(t *testing.T) {

	cmd := ratelimits.NewCmdRatelimits()

	convey.Convey("RatelimitsCreate", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"create"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/ratelimits/create.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestRatelimitsUpdate(t *testing.T) {

	cmd := ratelimits.NewCmdRatelimits()

	convey.Convey("RatelimitsUpdate", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"update"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/ratelimits/update.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestRatelimitsEnable(t *testing.T) {

	cmd := ratelimits.NewCmdRatelimits()

	convey.Convey("RatelimitsEnable", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"enable"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/ratelimits/enable.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestRatelimitsDelete(t *testing.T) {

	cmd := ratelimits.NewCmdRatelimits()

	convey.Convey("RatelimitsDelete", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"delete"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/ratelimits/delete.json")

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
