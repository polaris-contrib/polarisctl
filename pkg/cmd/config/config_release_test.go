package config_test

import (
	"bytes"

	"github.com/0226zy/polarisctl/pkg/cmd/config"
	"github.com/glycerine/goconvey/convey"

	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

	"testing"
)

func TestReleaseList(t *testing.T) {

	cmd := config.NewCmdConfig()

	convey.Convey("ReleaseList", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"release", "list"}

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

func TestReleaseHistory(t *testing.T) {

	cmd := config.NewCmdConfig()

	convey.Convey("ReleaseHistory", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"release", "history"}

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

func TestReleaseVersions(t *testing.T) {

	cmd := config.NewCmdConfig()

	convey.Convey("ReleaseVersions", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"release", "versions"}

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

func TestReleaseInfo(t *testing.T) {

	cmd := config.NewCmdConfig()

	convey.Convey("ReleaseInfo", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"release", "info"}

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

func TestReleaseCreate(t *testing.T) {

	cmd := config.NewCmdConfig()

	convey.Convey("ReleaseRelease", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"release", "release"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/config/release/release.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestReleaseDelete(t *testing.T) {

	cmd := config.NewCmdConfig()

	convey.Convey("ReleaseDelete", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"release", "delete"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/config/release/delete.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestReleaseRollback(t *testing.T) {

	cmd := config.NewCmdConfig()

	convey.Convey("ReleaseRollback", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"release", "rollback"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/config/release/rollback.json")

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
