package scan

import (
	"errors"

	"github.com/fullstack-devops/skipper/internal/app/skipper-ddm/models"
	"github.com/fullstack-devops/skipper/internal/app/skipper-ddm/scanner"
	"github.com/spf13/cobra"
)

var (
	FileType   string
	Server     string
	Repository string
)

var Cmd = &cobra.Command{
	Use:   "scan [FILE]",
	Short: "scan searches after updates in a specific file",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		scanner.ScanSingleFile(args[0], models.FileType(FileType))
	},
}

func init() {
	// commands
	// Cmd.AddCommand(...Cmd)

	// Flags
	Cmd.Flags().StringVarP(&FileType, "file-type", "t", "", "predetermain the file type, eg.: Dockerfile")

	// Cmd.MarkFlagRequired("repository")
	// Cmd.MarkFlagRequired("token")

	// exclusive Flags
}
