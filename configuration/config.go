package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configuration configuration setup
type Configuration struct {
	Host                string `json:"host"`
	Port                string `json:"port"`
	JWTKey              string `json:"jwtKey"`
	StreamServerLiveURL string `json:"streamServerLiveUrl"`
	StreamThumbnailURL  string `json:"streamThumbnailUrl"`
	StreamServerRTMPURL string `json:"streamServerRtmpUrl"`
	StreamVideoURL      string `json:"streamVideoUrl"`
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
