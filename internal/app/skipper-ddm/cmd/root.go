package cmd

import (
	"fmt"

	"github.com/fullstack-devops/skipper/internal/app/build"
	"github.com/fullstack-devops/skipper/internal/app/skipper-ddm/cmd/scan"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Verbose bool

var RootCmd = &cobra.Command{
	Use:              "skddm [OPTIONS] [COMMANDS]",
	TraverseChildren: true,
	Short:            "Skipper DDM helps you identify and update outdated dependencies.",
	Long: `Find more information and examples at https://github.com/fullstack-devops/skipper
		current supported formats:
		- Dockerfile
		- Shell- and Textfile`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Awesome-CI",
	Long:  `All software has versions. This is Awesome-CI's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version:\t", build.Version)
		fmt.Println("Commit: \t", build.CommitHash)
		fmt.Println("Date:   \t", build.BuildDate)
	},
}

func init() {
	// commands
	RootCmd.AddCommand(versionCmd)
	RootCmd.AddCommand(scan.Cmd)
	// RootCmd.AddCommand(update.Cmd)
	// RootCmd.AddCommand(upgrade.Cmd)

	// flags
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")

	// PreRuns
	RootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if Verbose {
			logrus.SetLevel(logrus.TraceLevel)
		}
		return nil
	}
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		logrus.Fatalln(err)
	}
}
