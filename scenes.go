package domoto

import (
	"encoding/json"
	"strconv"
)

// Scene is the description of a scene
type Scene struct {
	Favorite   int    `json:"Favorite"`
	HardwareID int    `json:"HardwareID"`
	LastUpdate string `json:"LastUpdate"`
	Name       string `json:"Name"`
	Status     string `json:"Status"`
	Timers     string `json:"Timers"`
	Type       string `json:"Type"`
	Idx        string `json:"idx"`
}

// Scenes is a list of all the scenes
type Scenes struct {
	Scene  []Scene `json:"result"`
	Status string  `json:"status"`
	Title  string  `json:"title"`
}

// SceneTimer is a timer for a scene
type SceneTimer struct {
	Active     string `json:"Active"`
	Cmd        int    `json:"Cmd"`
	Date       string `json:"Date"`
	Days       int    `json:"Days"`
	Hue        int    `json:"Hue"`
	Level      int    `json:"Level"`
	Randomness bool   `json:"Randomness"`
	Time       string `json:"Time"`
	Type       int    `json:"Type"`
	Idx        string `json:"idx"`
}

// SceneTimers is a list of all the scene timers
type SceneTimers struct {
	Timers []SceneTimer `json:"result"`
	Status string       `json:"status"`
	Title  string       `json:"title"`
}

// SceneDevices is the list of all devices in a scene/group
type SceneDevices struct {
	Result []struct {
		Command    string `json:"Command"`
		DevID      string `json:"DevID"`
		DevRealIdx string `json:"DevRealIdx"`
		Hue        int    `json:"Hue"`
		ID         string `json:"ID"`
		Level      int    `json:"Level"`
		Name       string `json:"Name"`
		OffDelay   int    `json:"OffDelay"`
		OnDelay    int    `json:"OnDelay"`
		Order      int    `json:"Order"`
		SubType    string `json:"SubType"`
		Type       string `json:"Type"`
	} `json:"result"`
	Status string `json:"status"`
	Title  string `json:"title"`
}

// AllScenes returns all scenes and groups
func (c *Config) AllScenes() (s Scenes, err error) {
	// prepare params for API call
	qp := make(map[string]string)
	qp["type"] = "scenes"
	// execute API call
	resp, err := c.Call(&qp)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Body(), &s)
	return
}

// SceneDevices returns all devices in scenes and groups
func (c *Config) SceneDevices(id int) (s SceneDevices, err error) {
	// /json.htm?type=command&param=getscenedevices&idx=1&isscene=true
	// prepare params for API call
	qp := make(map[string]string)
	qp["type"] = "command"
	qp["param"] = "getscenedevices"
	qp["idx"] = strconv.Itoa(id)
	qp["isscene"] = "true"
	resp, err := c.Call(&qp)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Body(), &s)
	return
}

// SceneSwitch turns a scene or group on/off
// cmd is either "On" or "Off", scenes can only be turned on
func (c *Config) SceneSwitch(id int, cmd string) (res Result, err error) {
	// example: /json.htm?type=command&param=switchscene&idx=&switchcmd=
	// prepare params for API call
	qp := make(map[string]string)
	qp["type"] = "command"
	qp["param"] = "switchscene"
	qp["idx"] = strconv.Itoa(id)
	qp["switchcmd"] = cmd
	// execute API call
	resp, err := c.Call(&qp)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Body(), &res)
	return
}

// SceneTimers list all timers for a scene or group
func (c *Config) SceneTimers(id int) (res SceneTimers, err error) {
	// /json.htm?type=scenetimers&idx=number
	// prepare params for API call
	qp := make(map[string]string)
	qp["type"] = "scenetimers"
	qp["idx"] = strconv.Itoa(id)
	// execute API call
	resp, err := c.Call(&qp)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Body(), &res)
	return
}
