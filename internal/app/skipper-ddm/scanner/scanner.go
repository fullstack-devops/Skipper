package scanner

import (
	"bufio"
	"os"
	"regexp"

	"github.com/fullstack-devops/skipper/internal/app/skipper-ddm/models"
	"github.com/sirupsen/logrus"
)

const (
	annotaionRegexString = `.*@SKIPPER +([A-Za-z]+) +(https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*))`
	dockerEnvRegexString = `ENV +([A-Za-z0-9_]+)( +|=)"?([0-9\.]+)"?`
)

var (
	annotaionRegexp *regexp.Regexp = regexp.MustCompile(annotaionRegexString)
	dockerEnvRegexp *regexp.Regexp = regexp.MustCompile(dockerEnvRegexString)
)

func ScanSingleFile(fileName string, fileType models.FileType) {
	file, err := os.Open(fileName)
	if err != nil {
		logrus.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	nextLineTarget := false

	for scanner.Scan() {
		if annotaionRegexp.MatchString(scanner.Text()) {
			matches := annotaionRegexp.FindStringSubmatch(scanner.Text())
			nextLineTarget = true
			logrus.Printf("found type %s release at %s", matches[1], matches[2])

		} else if nextLineTarget {
			if dockerEnvRegexp.MatchString(scanner.Text()) {
				matches := dockerEnvRegexp.FindStringSubmatch(scanner.Text())
				logrus.Printf("found set version %s", matches[3])
			} else {
				logrus.Warnln("line should match, but didn't")
			}
			nextLineTarget = false
		}
	}
}
