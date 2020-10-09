// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

package configuration

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/arduino/arduino-cli/cli/feedback"
	paths "github.com/arduino/go-paths-helper"
	"github.com/arduino/go-win32-utils"
	"github.com/sirupsen/logrus"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/spf13/viper"
)

// Init initialize defaults and read the configuration file.
// Please note the logging system hasn't been configured yet,
// so logging shouldn't be used here.
func Init(configPath string) {
	// Config file metadata
	jww.SetStdoutThreshold(jww.LevelFatal)

	configDir := paths.New(configPath)
	if configDir != nil && !configDir.IsDir() {
		viper.SetConfigName(strings.TrimSuffix(configDir.Base(), configDir.Ext()))
	} else {
		viper.SetConfigName("arduino-cli")
	}

	// Get default data path if none was provided
	if configPath == "" {
		configPath = getDefaultArduinoDataDir()
		viper.AddConfigPath(configPath)
	} else {
		viper.AddConfigPath(filepath.Dir(configPath))
	}

	// Bind env vars
	viper.SetEnvPrefix("ARDUINO")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	// Bind env aliases to keep backward compatibility
	viper.BindEnv("directories.User", "ARDUINO_SKETCHBOOK_DIR")
	viper.BindEnv("directories.Downloads", "ARDUINO_DOWNLOADS_DIR")
	viper.BindEnv("directories.Data", "ARDUINO_DATA_DIR")

	// Early access directories.Data and directories.User in case
	// those were set through env vars or cli flags
	dataDir := viper.GetString("directories.Data")
	if dataDir == "" {
		dataDir = getDefaultArduinoDataDir()
	}
	userDir := viper.GetString("directories.User")
	if userDir == "" {
		userDir = getDefaultUserDir()
	}

	// Set default values for all the settings
	setDefaults(dataDir, userDir)

	// Attempt to read config file
	if err := viper.ReadInConfig(); err != nil {
		// ConfigFileNotFoundError is acceptable, anything else
		// should be reported to the user
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			feedback.Errorf("Error reading config file: %v", err)
		}
	}
}

// getDefaultArduinoDataDir returns the full path to the default arduino folder
func getDefaultArduinoDataDir() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		feedback.Errorf("Unable to get user home dir: %v", err)
		return "."
	}

	switch runtime.GOOS {
	case "linux":
		return filepath.Join(userHomeDir, ".arduino15")
	case "darwin":
		return filepath.Join(userHomeDir, "Library", "Arduino15")
	case "windows":
		localAppDataPath, err := win32.GetLocalAppDataFolder()
		if err != nil {
			feedback.Errorf("Unable to get Local App Data Folder: %v", err)
			return "."
		}
		return filepath.Join(localAppDataPath, "Arduino15")
	default:
		return "."
	}
}

// getDefaultUserDir returns the full path to the default user folder
func getDefaultUserDir() string {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		feedback.Errorf("Unable to get user home dir: %v", err)
		return "."
	}

	switch runtime.GOOS {
	case "linux":
		return filepath.Join(userHomeDir, "Arduino")
	case "darwin":
		return filepath.Join(userHomeDir, "Documents", "Arduino")
	case "windows":
		documentsPath, err := win32.GetDocumentsFolder()
		if err != nil {
			feedback.Errorf("Unable to get Documents Folder: %v", err)
			return "."
		}
		return filepath.Join(documentsPath, "Arduino")
	default:
		return "."
	}
}

// IsBundledInDesktopIDE returns true if the CLI is bundled with the Arduino IDE.
func IsBundledInDesktopIDE() bool {
	// value is cached the first time we run the check
	if viper.IsSet("IDE.Bundled") {
		return viper.GetBool("IDE.Bundled")
	}

	viper.Set("IDE.Bundled", false)
	viper.Set("IDE.Portable", false)

	logrus.Info("Checking if CLI is Bundled into the IDE")
	executable, err := os.Executable()
	if err != nil {
		feedback.Errorf("Cannot get executable path: %v", err)
		return false
	}

	executablePath, err := filepath.EvalSymlinks(executable)
	if err != nil {
		feedback.Errorf("Cannot get executable path: %v", err)
		return false
	}

	ideDir := paths.New(executablePath).Parent()
	logrus.WithField("dir", ideDir).Trace("Candidate IDE directory")

	// To determine if the CLI is bundled with an IDE, We check an arbitrary
	// number of folders that are part of the IDE install tree
	tests := []string{
		"tools-builder",
		"examples/01.Basics/Blink",
	}

	for _, test := range tests {
		if !ideDir.Join(test).Exist() {
			// the test folder doesn't exist or is not accessible
			return false
		}
	}

	logrus.Info("The CLI is bundled in the Arduino IDE")

	// Persist IDE-related config settings
	viper.Set("IDE.Bundled", true)
	viper.Set("IDE.Directory", ideDir)

	// Check whether this is a portable install
	if ideDir.Join("portable").Exist() {
		logrus.Info("The IDE installation is 'portable'")
		viper.Set("IDE.Portable", true)
	}

	return true
}

// FindConfigFile returns the config file path using the argument '--config-file' if specified or via the current working dir
func FindConfigFile(args []string) string {
	configFile := ""
	for i, arg := range args {
		// 0 --config-file ss
		if arg == "--config-file" {
			if len(args) > i+1 {
				configFile = args[i+1]
			}
		}
	}

	if configFile != "" {
		return configFile
	}

	return searchCwdForConfig()
}

func searchCwdForConfig() string {
	cwd, err := os.Getwd()

	if err != nil {
		return ""
	}

	configFile := searchConfigTree(cwd)
	if configFile == "" {
		return configFile
	}

	return configFile + string(os.PathSeparator) + "arduino-cli.yaml"
}

func searchConfigTree(cwd string) string {

	// go back up to root and search for the config file
	for {
		if _, err := os.Stat(filepath.Join(cwd, "arduino-cli.yaml")); err == nil {
			// config file found
			return cwd
		} else if os.IsNotExist(err) {
			// no config file found
			next := filepath.Dir(cwd)
			if next == cwd {
				return ""
			}
			cwd = next
		} else {
			// some error we can't handle happened
			return ""
		}
	}

}
