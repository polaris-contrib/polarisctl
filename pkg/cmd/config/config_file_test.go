package config_test

import (
	"bytes"

	"github.com/0226zy/polarisctl/pkg/cmd/config"
	"github.com/glycerine/goconvey/convey"

	"github.com/0226zy/polarisctl/pkg/entity"
	"github.com/0226zy/polarisctl/pkg/repo"

	"testing"
)

func TestFileBygroup(t *testing.T) {

	cmd := config.NewCmdConfig()

	convey.Convey("FileBygroup", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"file", "bygroup"}

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

func TestFileSearch(t *testing.T) {

	cmd := config.NewCmdConfig()

	convey.Convey("FileSearch", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"file", "search"}

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

func TestFileExport(t *testing.T) {

	cmd := config.NewCmdConfig()

	convey.Convey("FileExport", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"file", "export"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/config/file/export.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestFileImport(t *testing.T) {

	cmd := config.NewCmdConfig()

	convey.Convey("FileImport", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"file", "import"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/config/file/import.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestFileCreate(t *testing.T) {

	cmd := config.NewCmdConfig()

	convey.Convey("FileCreate", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"file", "create"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/config/file/create.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestFileCreateandpub(t *testing.T) {

	cmd := config.NewCmdConfig()

	convey.Convey("FileCreateandpub", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"file", "createandpub"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/config/file/createandpub.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestFileDelete(t *testing.T) {

	cmd := config.NewCmdConfig()

	convey.Convey("FileDelete", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"file", "delete"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/config/file/delete.json")

		cmd.SetArgs(args)

		convey.Convey("Should not panic and return nil error", func() {
			convey.So(func() { err = cmd.Execute() }, convey.ShouldNotPanic)
			convey.So(err, convey.ShouldBeNil)
		})
	})
}

func TestFileUpdate(t *testing.T) {

	cmd := config.NewCmdConfig()

	convey.Convey("FileUpdate", t, func() {

		out := bytes.NewBufferString("")
		cmd.SetOut(out)

		args := []string{"file", "update"}

		var err error

		args = append(args, "-f")

		args = append(args, "../../../example/config/file/update.json")

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
