package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/jenkins-x/jx-logging/pkg/log"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		log.Logger().Fatalf("failed to get current working directory: %v", err)
	}
	err = migrate(dir)
	if err != nil {
		log.Logger().Fatalf("failed to migrate: %v", err)
	}
	log.Logger().Info("migrated")
}

func migrate(dir string) error {
	err := filepath.Walk(dir, visit)
	if err != nil {
		panic(err)
	}
	return nil
}

func visit(path string, fi os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !!fi.IsDir() {
		return nil //
	}

	pattern := "*.go"
	if os.Getenv("TEST_PATTERN") != "" {
		pattern = os.Getenv("TEST_PATTERN")
	}

	matched, err := filepath.Match(pattern, fi.Name())

	if err != nil {
		panic(err)
		return err
	}

	if matched {
		read, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		newContents := strings.Replace(string(read), "github.com/jenkins-x/jx-api/v3", "github.com/jenkins-x/jx-api/v4", -1)
		newContents = strings.Replace(newContents, "github.com/jenkins-x/jx-api/v4/pkg/apis/jenkins.io/v1", "github.com/jenkins-x/jx-api/v4/pkg/apis/core/v4beta1", -1)
		newContents = strings.Replace(newContents, "github.com/jenkins-x/jx-api/v4/pkg/apis/jenkins.io", "github.com/jenkins-x/jx-api/v4/pkg/apis/core", -1)
		newContents = strings.Replace(newContents, "\"github.com/jenkins-x/jx-api/v4/pkg/config\"", "jxcore \"github.com/jenkins-x/jx-api/v4/pkg/apis/core/v4beta1\"", -1)
		newContents = strings.Replace(newContents, "config.RequirementsConfig", "jxcore.RequirementsConfig", -1)
		newContents = strings.Replace(newContents, "jxconfig.LoadRequirementsConfig", "jxcore.LoadRequirementsConfig", -1)
		newContents = strings.Replace(newContents, "config.LoadRequirementsConfig", "jxcore.LoadRequirementsConfig", -1)
		newContents = strings.Replace(newContents, "config.NewRequirementsConfig", "jxcore.NewRequirementsConfig", -1)
		newContents = strings.Replace(newContents, ".JenkinsV1()", ".CoreV4beta1()", -1)
		newContents = strings.Replace(newContents, "github.com/jenkins-x/jx-api/pkg/config", "FIX_ME_CONFIG_HAS_MOVED_TO jxcore github.com/jenkins-x/jx-api/v4/pkg/apis/core/v4beta1", -1)

		err = ioutil.WriteFile(path, []byte(newContents), 0)
		if err != nil {
			panic(err)
		}
	}

	return nil
}
