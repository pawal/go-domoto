package domoto

import (
	"encoding/json"
	"strconv"
)

// Devices is a struct for the device listings
type Devices struct {
	ActTime    int      `json:"ActTime"`
	ServerTime string   `json:"ServerTime"`
	Sunrise    string   `json:"Sunrise"`
	Sunset     string   `json:"Sunset"`
	Devices    []Device `json:"result"`
	Status     string   `json:"status"`
	Title      string   `json:"title"`
}

// Device is the struct for a single device
type Device struct {
	AddjMulti         float64 `json:"AddjMulti,omitempty"`
	AddjMulti2        float64 `json:"AddjMulti2,omitempty"`
	AddjValue         float64 `json:"AddjValue,omitempty"`
	AddjValue2        float64 `json:"AddjValue2,omitempty"`
	BatteryLevel      int     `json:"BatteryLevel,omitempty"`
	CustomImage       int     `json:"CustomImage,omitempty"`
	Data              string  `json:"Data,omitempty"`
	Description       string  `json:"Description,omitempty"`
	Favorite          int     `json:"Favorite,omitempty"`
	HardwareID        int     `json:"HardwareID,omitempty"`
	HardwareName      string  `json:"HardwareName,omitempty"`
	HardwareType      string  `json:"HardwareType,omitempty"`
	HardwareTypeVal   int     `json:"HardwareTypeVal,omitempty"`
	HaveDimmer        bool    `json:"HaveDimmer,omitempty"`
	HaveGroupCmd      bool    `json:"HaveGroupCmd,omitempty"`
	HaveTimeout       bool    `json:"HaveTimeout,omitempty"`
	ID                string  `json:"ID"`
	Image             string  `json:"Image,omitempty"`
	IsSubDevice       bool    `json:"IsSubDevice,omitempty"`
	LastUpdate        string  `json:"LastUpdate,omitempty"`
	Level             int     `json:"Level,omitempty"`
	LevelInt          int     `json:"LevelInt,omitempty"`
	MaxDimLevel       int     `json:"MaxDimLevel,omitempty"`
	Name              string  `json:"Name"`
	Notifications     string  `json:"Notifications,omitempty"`
	PlanID            string  `json:"PlanID,omitempty"`
	PlanIDs           []int   `json:"PlanIDs,omitempty"`
	Protected         bool    `json:"Protected"`
	ShowNotifications bool    `json:"ShowNotifications,omitempty"`
	SignalLevel       string  `json:"SignalLevel,omitempty"`
	Status            string  `json:"Status,omitempty"`
	StrParam1         string  `json:"StrParam1,omitempty"`
	StrParam2         string  `json:"StrParam2,omitempty"`
	SubType           string  `json:"SubType,omitempty"`
	SwitchType        string  `json:"SwitchType,omitempty"`
	SwitchTypeVal     int     `json:"SwitchTypeVal,omitempty"`
	Timers            string  `json:"Timers,omitempty"`
	Type              string  `json:"Type"`
	TypeImg           string  `json:"TypeImg,omitempty"`
	Unit              int     `json:"Unit"`
	Used              int     `json:"Used"`
	UsedByCamera      bool    `json:"UsedByCamera,omitempty"`
	XOffset           string  `json:"XOffset,omitempty"`
	YOffset           string  `json:"YOffset,omitempty"`
	Idx               string  `json:"idx"`
}

// Result is the result of an action:
// {
//    "status" : "OK",
//    "title" : "SwitchLight"
// }
// OR
// {
//    "message" : "WRONG CODE",
//    "status" : "ERROR",
//    "title" : "SwitchLight"
// }
type Result struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Title   string `json:"title"`
}

// AllDevices retrieves all devices, optionally with a filter
// Filter types:
//   light = Get all lights/switches
//   weather = Get all weather devices
//   temp = Get all temperature devices
//   utility = Get all utility devices
// Leave filter with an empty or other string for all devices.
func (c *Config) AllDevices(filter string) (d Devices, err error) {
	// prepare params for API call
	qp := make(map[string]string)
	qp["type"] = "devices"
	qp["used"] = "true"
	qp["order"] = "Name"
	qp["filter"] = filter
	// execute API call
	resp, err := c.Call(&qp)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Body(), &d)
	return
}

// Device retrieves a device with its id
func (c *Config) Device(id int) (d Device, err error) {
	// prepare params for API call
	qp := make(map[string]string)
	qp["type"] = "devices"
	qp["rid"] = strconv.Itoa(id)
	// execute API call
	resp, err := c.Call(&qp)
	if err != nil {
		return
	}

	var ds Devices
	err = json.Unmarshal(resp.Body(), &ds)
	if err != nil {
		return
	}
	// return first device in list
	return ds.Devices[0], nil
}

// DeviceSwitch switches on or off a device switch
// cmd is either "On" or "Off"
func (c *Config) DeviceSwitch(id int, cmd string) (res Result, err error) {
	// example: type=command&param=switchlight&idx=99&switchcmd=On
	// prepare params for API call
	qp := make(map[string]string)
	qp["type"] = "command"
	qp["param"] = "switchlight"
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

// DeviceToggle toggles a device switch, on or off
// Result is the same as for a Device switch
func (c *Config) DeviceToggle(id int) (res Result, err error) {
	// example: /json.htm?type=command&param=switchlight&idx=99&switchcmd=Toggle
	// prepare params for API call
	qp := make(map[string]string)
	qp["type"] = "command"
	qp["param"] = "switchlight"
	qp["idx"] = strconv.Itoa(id)
	qp["switchcmd"] = "Toggle"
	// execute API call
	resp, err := c.Call(&qp)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp.Body(), &res)
	return
}
