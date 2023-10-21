package faultdetectors_test

import (
	"bytes"

	"github.com/0226zy/polarisctl/pkg/cmd/faultdetectors"
	"github.com/glycerine/goconvey/convey"

	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

	"testing"
)

func TestFaultdetectorsList(t *testing.T) {

	cmd := faultdetectors.NewCmdFaultdetectors()

	convey.Convey("FaultdetectorsList", t, func() {

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

func TestFaultdetectorsCreate(t *testing.T) {

	cmd := faultdetectors.NewCmdFaultdetectors()

	convey.Convey("FaultdetectorsCreate", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"create"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/faultdetectors/create.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestFaultdetectorsDelete(t *testing.T) {

	cmd := faultdetectors.NewCmdFaultdetectors()

	convey.Convey("FaultdetectorsDelete", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"delete"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/faultdetectors/delete.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestFaultdetectorsUpdate(t *testing.T) {

	cmd := faultdetectors.NewCmdFaultdetectors()

	convey.Convey("FaultdetectorsUpdate", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"update"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/faultdetectors/update.json")

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
