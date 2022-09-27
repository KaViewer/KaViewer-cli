package mvn

import (
	"github.com/kaviewer/kaviewer-cli/cmd/common"
	"github.com/spf13/cobra"
)

// Maven Cmd for build KaViewer operations.
var Maven = &cobra.Command{
	Use: "mvn",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func build() {
	var build = &cobra.Command{
		Use:   "jar",
		Short: "Build KaViewer Exec",
		Run: func(cmd *cobra.Command, args []string) {
			runTest, _ := cmd.Flags().GetBool("run-test")
			if runTest {
				run("clean", "package")
			} else {
				run("clean", "package", "-Dmaven.test.skip=true --file pom.xml")
			}
		},
	}
	build.Flags().BoolP("run-test", "t", false, "skip test or not, default false")
	Maven.AddCommand(build)
}

func init() {
	build()
}

func run(params ...string) {
	common.ExecCmd("mvn", params...)
}
