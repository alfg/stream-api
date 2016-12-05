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
	RtmpHost            string                `json:"rtmpHost"`
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
	file, _ := os.Open("defaults.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error: ", err)
	}
	getConfigFromEnv(&configuration)
	return &configuration
}

func getConfigFromEnv(config *Configuration) {
	streamThumbnailURL := os.Getenv("STREAM_THUMBNAIL_URL")
	if streamThumbnailURL != "" {
		config.StreamThumbnailURL = streamThumbnailURL
	}

	streamServerRtmpURL := os.Getenv("STREAM_SERVER_RTMP_URL")
	if streamServerRtmpURL != "" {
		config.StreamServerRTMPURL = streamServerRtmpURL
	}

	streamServerLiveURL := os.Getenv("STREAM_SERVER_LIVE_URL")
	if streamThumbnailURL != "" {
		config.StreamServerLiveURL = streamServerLiveURL
	}

	streamVideoURL := os.Getenv("STREAM_VIDEO_URL")
	if streamVideoURL != "" {
		config.StreamVideoURL = streamVideoURL
	}

	rtmpHost := os.Getenv("RTMP_HOST")
	if rtmpHost != "" {
		config.RtmpHost = rtmpHost
	}

}
