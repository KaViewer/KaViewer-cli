package cmd

import (
	"fmt"
	"github.com/kaviewer/kaviewer-cli/cmd/common"
	"github.com/kaviewer/kaviewer-cli/cmd/docker"
	"github.com/kaviewer/kaviewer-cli/cmd/mvn"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "kaviewer",
	Short: "Setup KaViewer asap.",
	Long:  "KaViewer-cli is a tool to setup KaViewer with advantage cmd.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func run() {
	var run = &cobra.Command{
		Use:   "run",
		Short: "Run KaViewer",
		Run: func(cmd *cobra.Command, args []string) {
			runnerJar := "app/target/app-0.0.1.jar"
			jar, _ := cmd.Flags().GetString("jar")

			if jar != "" {
				runnerJar = jar

			}
			common.ExecCmd("java", "-jar", runnerJar)
		},
	}

	run.Flags().StringP("jar", "j", "app/target/app-0.0.1.jar", "the runner jar location")
	rootCmd.AddCommand(run)

}

func build() {
	var build = &cobra.Command{
		Use:   "build",
		Short: "build KaViewer",
		Run: func(cmd *cobra.Command, args []string) {
			// build both front and back end
		},
	}

	rootCmd.AddCommand(build)

}

func clean() {

	var clean = &cobra.Command{
		Use:   "clean",
		Short: "Clean Jar",
		Run: func(cmd *cobra.Command, args []string) {
			common.ExecCmd("mvn", "clean")
		},
	}
	rootCmd.AddCommand(clean)
}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	run()
	clean()
	build()

	rootCmd.AddCommand(docker.Docker)
	rootCmd.AddCommand(mvn.Maven)
}
