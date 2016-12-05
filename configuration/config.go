package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration configuration setup
type Configuration struct {
	Host                string                `json:"host"`
	Port                string                `json:"port"`
	JWTKey              string                `json:"jwtKey"`
	Database            DatabaseConfiguration `json:"db"`
	StreamServerLiveURL string                `json:"streamServerLiveUrl"`
	StreamThumbnailURL  string                `json:"streamThumbnailUrl"`
	StreamServerRTMPURL string                `json:"streamServerRtmpUrl"`
	StreamVideoURL      string                `json:"streamVideoUrl"`
}

// DatabaseConfiguration Database configuration setup.
type DatabaseConfiguration struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

// ConfigurationSetup configuration setup
func ConfigurationSetup() *Configuration {
	file, _ := os.Open("conf.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error: ", err)
	}
	return &configuration
}
