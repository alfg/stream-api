package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Settings configuration setup
type Settings struct {
	Host                string   `json:"host"`
	Port                string   `json:"port"`
	JWTKey              string   `json:"jwtKey"`
	Database            Database `json:"db"`
	StreamServerLiveURL string   `json:"streamServerLiveUrl"`
	StreamThumbnailURL  string   `json:"streamThumbnailUrl"`
	StreamServerRTMPURL string   `json:"streamServerRtmpUrl"`
	StreamVideoURL      string   `json:"streamVideoUrl"`
	RtmpHost            string   `json:"rtmpHost"`
}

// Database Database configuration setup.
type Database struct {
	Server   string `json:"server"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

var settings = Settings{}

// Init configuration setup
func init() {
	file, err := ioutil.ReadFile("settings/defaults.json")
	if err != nil {
		fmt.Println("error reading settings file.")
	}
	settings = Settings{}
	err = json.Unmarshal(file, &settings)
	if err != nil {
		fmt.Println("error parsing settings file.")
	}
	fmt.Println("getting settings")
	GetSettingsFromEnv(&settings)
}

// Get Gets settings.
func Get() Settings {
	return settings
}

// GetSettingsFromEnv Gets settings from environment variable with default.
func GetSettingsFromEnv(settings *Settings) {
	streamThumbnailURL := os.Getenv("STREAM_THUMBNAIL_URL")
	if streamThumbnailURL != "" {
		settings.StreamThumbnailURL = streamThumbnailURL
	}

	streamServerRtmpURL := os.Getenv("STREAM_SERVER_RTMP_URL")
	if streamServerRtmpURL != "" {
		settings.StreamServerRTMPURL = streamServerRtmpURL
	}

	streamServerLiveURL := os.Getenv("STREAM_SERVER_LIVE_URL")
	if streamThumbnailURL != "" {
		settings.StreamServerLiveURL = streamServerLiveURL
	}

	streamVideoURL := os.Getenv("STREAM_VIDEO_URL")
	if streamVideoURL != "" {
		settings.StreamVideoURL = streamVideoURL
	}

	rtmpHost := os.Getenv("RTMP_HOST")
	if rtmpHost != "" {
		settings.RtmpHost = rtmpHost
	}
}
