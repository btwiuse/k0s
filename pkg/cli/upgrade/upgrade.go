package upgrade

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/creativeprojects/go-selfupdate"
	"github.com/btwiuse/version"
)

func Run(args []string) error {
	currentVersion := version.Info.GitVersion

	log.Println(fmt.Sprintf("current version: %s", currentVersion))

	err := upgrade(currentVersion)

	return err
}

func upgrade(version string) error {
	latest, found, err := selfupdate.DetectLatest(context.TODO(), selfupdate.ParseSlug("btwiuse/k0s"))
	if err != nil {
		return fmt.Errorf("error occurred while detecting version: %v", err)
	}
	if !found {
		return fmt.Errorf("latest version for %s/%s could not be found from github repository", runtime.GOOS, runtime.GOARCH)
	}

	log.Println(fmt.Sprintf("found latest: %s", latest.AssetURL))

	if latest.LessOrEqual(version) {
		log.Printf("Current version (%s) is the latest", version)
		return nil
	}

	exe, err := os.Executable()
	if err != nil {
		return errors.New("could not locate executable path")
	}
	if err := selfupdate.UpdateTo(context.TODO(), latest.AssetURL, latest.AssetName, exe); err != nil {
		return fmt.Errorf("error occurred while updating binary: %v", err)
	}
	log.Printf("Successfully updated to version %s", latest.Version())
	return nil
}
