package skipperddm

import (
	"os"

	"github.com/fullstack-devops/skipper/internal/app/skipper-ddm/cmd"
	"github.com/sirupsen/logrus"
)

func init() {
	// log.SetFormatter(&log.JSONFormatter{})
	logrus.SetOutput(os.Stderr)
	logrus.SetLevel(logrus.InfoLevel)
}

func SkipperDDM() {
	cmd.Execute()
}
