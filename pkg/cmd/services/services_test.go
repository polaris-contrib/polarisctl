package services_test

import (
	"bytes"

	"github.com/0226zy/polarisctl/pkg/cmd/services"
	"github.com/glycerine/goconvey/convey"

	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

	"testing"
)

func TestServicesList(t *testing.T) {

	cmd := services.NewCmdServices()

	convey.Convey("ServicesList", t, func() {

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

func TestServicesAll(t *testing.T) {

	cmd := services.NewCmdServices()

	convey.Convey("ServicesAll", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"all"}

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

func TestServicesCreate(t *testing.T) {

	cmd := services.NewCmdServices()

	convey.Convey("ServicesCreate", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"create"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/services/create.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestServicesUpdate(t *testing.T) {

	cmd := services.NewCmdServices()

	convey.Convey("ServicesUpdate", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"update"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/services/update.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestServicesDelete(t *testing.T) {

	cmd := services.NewCmdServices()

	convey.Convey("ServicesDelete", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"delete"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/services/delete.json")

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
