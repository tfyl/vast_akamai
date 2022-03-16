package vast_akamai

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Vast struct {
	apiKey   string
	endPoint string
	Validate bool
}

// NewVast returns a new Vast instance
// apiKey is the Vast API key
// endPoint is the Vast API endpoint
func NewVast(apiKey, endpoint string) *Vast {
	return &Vast{apiKey: apiKey, endPoint: endpoint, Validate: true}
}

func (v *Vast) GetSensor(p *Payload) (*Response, error) {
	if v.Validate {
		err := p.validate()
		if err != nil {
			return nil, err
		}
	}

	body, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", v.endPoint, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("apikey", v.apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var r Response
	err = json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		return nil, err
	}

	return &r, nil
}
