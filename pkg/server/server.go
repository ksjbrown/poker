package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
)

// This file contains functions and types related to operation of the rest server

// A server start requires:
// 	- Loading of configuration file, generation if it doesn't exist
// 	- Connection to database driver, defined in config
// 	-

type Server struct {
	Config ServerConfig
}

func (s *Server) LoadConfig() error {
	// calculate the os specific config file location
	configPath, err := PlatformConfigPath()
	if err != nil {
		// TODO: log
		return err
	}
	// if the file doesn't exist, load default values and create the file with these values
	var config *ServerConfig
	_, err = os.Stat(configPath)
	if errors.Is(err, fs.ErrNotExist) {
		// TODO: log
		s.Config = *DefaultConfig()
		return s.SaveConfig()
	} else if err != nil {
		return err
	}
	// file exists, open it
	file, err := os.Open(configPath)
	if err != nil {
		// TODO: log
		return err
	}
	defer file.Close()
	// parse the json config contents
	jsonDecoder := json.NewDecoder(file)
	err = jsonDecoder.Decode(config)
	if err != nil {
		// TODO: log
		return err
	}
	// succeeded
	s.Config = *config
	return nil
}

func (s *Server) SaveConfig() error {
	return fmt.Errorf("not yet implemented")
}

func DefaultConfig() *ServerConfig {
	// define default config options and return
	config := ServerConfig{
		Db: DatabaseConfig{
			Type: SQLITE,
			Args: map[DatabaseConfigOption]interface{}{},
		},
	}
	return &config
}

// PlatformConfigPath returns the path to the server config file, depending on the operating system.
func PlatformConfigPath() (string, error) {

	// Get the current user
	currentUser, err := user.Current()
	if err != nil {
		return "", err
	}

	// Determine the configuration directory path based on the operating system
	var configDir string
	switch os := runtime.GOOS; os {
	case "windows":
		configDir = filepath.Join(currentUser.HomeDir, "AppData", "Roaming", "ksjbrown", "poker")
	default:
		configDir = filepath.Join(currentUser.HomeDir, ".config", "ksjbrown", "poker")
	}
	return configDir, nil
}
