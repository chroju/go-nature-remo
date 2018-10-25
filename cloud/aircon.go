package natureremocloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Aircon struct {
	Range    Range  `json:"range"`
	TempUnit string `json:"tempUnit"`
}

type CurrentSettings struct {
	Temperature string `json:"temp"`
	Mode        string `json:"mode"`
	Volume      string `json:"vol"`
	Direction   string `json:"dir"`
	Button      string `json:"button"`
}

type AirconSettings struct {
	Temperature string `json:"temperature"`
	Mode        string `json:"operation_mode"`
	Volume      string `json:"air_volume"`
	Direction   string `json:"air_direction"`
	Button      string `json:"button"`
}

type Range struct {
	Modes        Modes    `json:"modes"`
	FixedButtons []string `json:"fixedButtons"`
}

type Modes struct {
	Cool AirconMode `json:"cool"`
	Warm AirconMode `json:"warm"`
	Dry  AirconMode `json:"dry"`
	Blow AirconMode `json:"blow"`
	Auto AirconMode `json:"auto"`
}

type AirconMode struct {
	Temp []string `json:"temp"`
	Vol  []string `json:"vol"`
	Dir  []string `json:"dir"`
}

func (c *Client) PostAirconSettings(applianceID string, settings AirconSettings) error {
	parameters, err := json.Marshal(settings)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST",
		c.urlFor("/appliances/"+applianceID+"/aircon_settings").String(),
		bytes.NewBuffer(parameters))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	resp, err := c.Request(req)
	defer closeResponse(resp)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to update aircon settings")
	}

	return nil
}
