package instances_test

import (
	"bytes"

	"github.com/0226zy/polarisctl/pkg/cmd/instances"
	"github.com/glycerine/goconvey/convey"

	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

	"testing"
)

func TestHostDelete(t *testing.T) {

	cmd := instances.NewCmdInstances()

	convey.Convey("HostDelete", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"host", "delete"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/instances/host/delete.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestHostIsolate(t *testing.T) {

	cmd := instances.NewCmdInstances()

	convey.Convey("HostIsolate", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"host", "isolate"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/instances/host/isolate.json")

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
