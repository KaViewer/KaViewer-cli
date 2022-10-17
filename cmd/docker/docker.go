package docker

import (
	"fmt"
	"github.com/spf13/cobra"
)

// Docker Cmd for operations in docker.
var Docker = &cobra.Command{
	Use: "docker",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			fmt.Println(arg)
		}
	},
}

func build() {

	var build = &cobra.Command{
		Use: "build",
		Run: func(cmd *cobra.Command, args []string) {
			buildType, _ := cmd.Flags().GetString("type")
			Handle(buildType)
		},
	}

	build.Flags().StringP("type", "t", "", "Indicate which the image type to build [local/app]")
	Docker.AddCommand(build)
}

func release() {

	var release = &cobra.Command{
		Use: "release",
		Run: func(cmd *cobra.Command, args []string) {
			for _, arg := range args {
				fmt.Println(arg)
			}
		},
	}

	Docker.AddCommand(release)

}

func init() {
	build()
	release()
}

// TODO
func run(params ...string) {

}
