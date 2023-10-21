package maintain_test

import (
	"bytes"

	"github.com/0226zy/polarisctl/pkg/cmd/maintain"
	"github.com/glycerine/goconvey/convey"

	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

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
