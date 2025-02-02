package cmd

import (
	"os"

	"git.sr.ht/nka/devc/backend/docker"
	"git.sr.ht/nka/devc/backend/dockercompose"
	"github.com/spf13/cobra"
)

var stopDown bool
var stopArgs []string

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop devcontainer services",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		switch rootBackend {
		case "dockerCompose":
			projectName := rootConfig.GetString("name") + "_devcontainer"
			dockerComposeFile := ".devcontainer/" + rootConfig.GetString("dockercomposefile")
			dockercompose.Stop(rootVerbose, projectName, dockerComposeFile, stopArgs...)
			if stopDown {
				dockercompose.Down(rootVerbose, projectName, dockerComposeFile, stopArgs...)
			}
		case "docker":
			path, _ := os.Getwd()
			container, _ := docker.GetContainer(path)
			docker.Stop(rootVerbose, container, stopArgs...)
			if stopDown {
				docker.Remove(rootVerbose, container)
			}
		}
	},
}

func init() {
	stopCmd.PersistentFlags().BoolVarP(&stopDown, "down", "d", false, "remove containers and networks")
	rootCmd.AddCommand(stopCmd)
}
