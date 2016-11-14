package services

import "fmt"

const baseURL = "http://192.168.99.100:8080/"

// RTMPStats model.
type RTMPStats struct {
	NginxVersion     string         `json:"nginx_version" xml:"nginx_version"`
	NginxRTMPVersion string         `json:"nginx_rtmp_version" xml:"nginx_rtmp_version"`
	Compiler         string         `json:"compiler" xml:"compiler"`
	Built            string         `json:"built" xml:"built"`
	PID              int            `json:"pid" xml:"pid"`
	Uptime           int            `json:"uptime" xml:"uptime"`
	NAccepted        int            `json:"naccepted" xml:"naccepted"`
	BWIn             int            `json:"bw_in" xml:"bw_in"`
	BWOut            int            `json:"bw_out" xml:"bw_out"`
	BytesIn          int            `json:"bytes_in" xml:"bytes_in"`
	BytesOut         int            `json:"bytes_out" xml:"bytes_out"`
	Server           []Applications `json:"server" xml:"server"`
}

// Applications model.
type Applications struct {
	Application []Application `json:"application" xml:"application"`
}

// Application model.
type Application struct {
	Name string `json:"name" xml:"name"`
	Live Live   `json:"live" xml:"live"`
}

// Live model.
type Live struct {
	Stream   Stream `json:"stream" xml:"stream"`
	NClients int    `json:"nclients" xml:"nclients"`
}

type Stream struct {
	Name          string         `json:"name" xml:"name"`
	Time          int            `json:"time" xml:"time"`
	BWIn          int            `json:"bw_in" xml:"bw_in"`
	BWOut         int            `json:"bw_out" xml:"bw_out"`
	BytesIn       int            `json:"bytes_in" xml:"bytes_in"`
	BytesOut      int            `json:"bytes_out" xml:"bytes_out"`
	BWAudio       int            `json:"bw_audio" xml:"bw_audio"`
	BWVideo       int            `json:"bw_video" xml:"bw_video"`
	StreamClients []StreamClient `json:"client" xml:"client"`
	Meta          []Meta         `json:"meta" xml:"meta"`
	// Publishing    bool           `json:"publishing,omitempty" xml:"publishing"`
	Active     bool `json:"active"`
	Publishing bool `json:"publishing"`
}

type Clients struct {
	Client []StreamClient `json:"client" xml:"client"`
}

type StreamClient struct {
	ID        int    `json:"id" xml:"id"`
	Address   string `json:"address" xml:"address"`
	Time      int    `json:"time" xml:"time"`
	Dropped   int    `json:"dropped" xml:"dropped"`
	AVSync    int    `json:"avsync" xml:"avsync"`
	Timestamp int    `json:"timestamp" xml:"timestamp"`
}

type Meta struct {
	Video Video `json:"video" xml:"video"`
	Audio Audio `json:"audio" xml:"audio"`
}

type Video struct {
	Width     int    `json:"width" xml:"width"`
	Height    int    `json:"height" xml:"height"`
	FrameRate int    `json:"frame_rate" xml:"frame_rate"`
	Codec     string `json:"codec" xml:"codec"`
	Compat    int    `json:"compat" xml:"compat"`
	Level     string `json:"level" xml:"level"`
}

type Audio struct {
	Codec      string `json:"codec" xml:"codec"`
	Profile    string `json:"profile" xml:"profile"`
	Channels   int    `json:"channels" xml:"channels"`
	SampleRate int    `json:"sample_rate" xml:"sample_rate"`
}

// GetRTMPStats Gets rtmp stats.
func (c *Client) GetRTMPStats() (*RTMPStats, error) {

	var url = baseURL + "stat.xml"

	rsp := &RTMPStats{}
	e := c.loadResponse(url, rsp, "xml")
	if e != nil {
		fmt.Print(e)
	}

	// Set default active and publishing to true.
	// If streams are ingesting, then they are active and publishing.
	for i := range rsp.Server {
		for j := range rsp.Server[i].Application {
			rsp.Server[i].Application[j].Live.Stream.Active = true
			rsp.Server[i].Application[j].Live.Stream.Publishing = true
		}
	}
	return rsp, e
}
