package utils

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

// GetConfig parse the devcontainer.json file and return a viper config object
func GetConfig() (*viper.Viper, error) {
	config := viper.New()
	config.AddConfigPath(".devcontainer/")
	config.SetConfigName("devcontainer")
	config.SetConfigType("json")

	// set aliases
	config.RegisterAlias("dockerfile", "build.dockerfile")
	config.RegisterAlias("context", "build.context")

	// set defaults
	path, _ := os.Getwd()
	dirName := filepath.Base(path)
	config.SetDefault("name", dirName)
	config.SetDefault("build.context", ".")
	config.SetDefault("updateRemoteUserUID", true)
	config.SetDefault("overrideCommand", true)

	err := config.ReadInConfig()

	return config, err
}

// CheckMutuallyExclusiveSettings does what its name says
func CheckMutuallyExclusiveSettings(config *viper.Viper) error {
	// check for mutally exclusive settings
	if config.Get("image") != nil && config.Get("dockerFile") != nil {
		return errors.New("you cannot use both 'image' and 'dockerFile' settings")
	}
	if config.Get("image") != nil && config.Get("build.dockerfile") != nil {
		return errors.New("you cannot use both 'image' and 'build.dockerfile' settings")
	}
	if config.Get("image") != nil && config.Get("dockerComposeFile") != nil {
		return errors.New("you cannot use both 'image' and 'dockerComposeFile' settings")
	}
	if config.Get("dockerFile") != nil && config.Get("dockerComposeFile") != nil {
		return errors.New("you cannot use both 'dockerFile' and 'dockerComposeFile' settings")
	}
	if config.Get("build.dockerfile") != nil && config.Get("dockerComposeFile") != nil {
		return errors.New("you cannot use both 'build.dockerfile' and 'dockerComposeFile' settings")
	}

	return nil
}

// RemoveFromSlice return a slice of strings without the given string
func RemoveFromSlice(sliceIn []string, remove string) (sliceOut []string) {
	for _, item := range sliceIn {
		if item != remove {
			sliceOut = append(sliceOut, item)
		}
	}

	return sliceOut
}
