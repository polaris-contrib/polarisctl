package services_test

import (
	"bytes"

	"github.com/0226zy/polarisctl/pkg/cmd/services"
	"github.com/glycerine/goconvey/convey"

	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

	"testing"
)

func TestAliasList(t *testing.T) {

	cmd := services.NewCmdServices()

	convey.Convey("AliasList", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"alias", "list"}

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

func TestAliasCreate(t *testing.T) {

	cmd := services.NewCmdServices()

	convey.Convey("AliasCreate", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"alias", "create"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/services/alias/create.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestAliasUpdate(t *testing.T) {

	cmd := services.NewCmdServices()

	convey.Convey("AliasUpdate", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"alias", "update"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/services/alias/update.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestAliasDelete(t *testing.T) {

	cmd := services.NewCmdServices()

	convey.Convey("AliasDelete", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"alias", "delete"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/services/alias/delete.json")

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
