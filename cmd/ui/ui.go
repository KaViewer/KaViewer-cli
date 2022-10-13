package ui

import (
	"github.com/kaviewer/kaviewer-cli/cmd/common"
	"github.com/spf13/cobra"
)

var UI = &cobra.Command{
	Use: "ui",
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		run("build", path)

	},
}

func init() {
	UI.Flags().StringP("path", "p", "build path", "where the front end folder located")
}

func run(params ...string) {
	common.ExecCmd("yarn", params...)
}
