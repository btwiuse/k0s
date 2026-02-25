package upgrade

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/btwiuse/version"
	"github.com/creativeprojects/go-selfupdate"
)

// ErrExecutablePath is returned when the executable path cannot be determined.
var ErrExecutablePath = errors.New("could not locate executable path")

func Run(args []string) error {
	currentVersion := version.Info.GitVersion

	log.Printf("current version: %s", currentVersion)

	err := upgrade(currentVersion)

	return err
}

func upgrade(version string) error {
	latest, found, err := selfupdate.DetectLatest(context.TODO(), selfupdate.ParseSlug("btwiuse/k0s"))
	if err != nil {
		return fmt.Errorf("error occurred while detecting version: %w", err)
	}
	if !found {
		return fmt.Errorf("latest version for %s/%s could not be found from github repository", runtime.GOOS, runtime.GOARCH)
	}

	log.Printf("found latest: %s", latest.AssetURL)

	if latest.LessOrEqual(version) {
		log.Printf("Current version (%s) is the latest", version)
		return nil
	}

	exe, err := os.Executable()
	if err != nil {
		return ErrExecutablePath
	}
	if err := selfupdate.UpdateTo(context.TODO(), latest.AssetURL, latest.AssetName, exe); err != nil {
		return fmt.Errorf("error occurred while updating binary: %w", err)
	}
	log.Printf("Successfully updated to version %s", latest.Version())
	return nil
}
