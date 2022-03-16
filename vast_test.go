package vast_akamai

import (
	"os"
	"testing"
)

func TestVast_GetSensor(t *testing.T) {
	vastKey := os.Getenv("VAST_KEY")
	if vastKey == "" {
		t.Skip("VAST_KEY environment variable not set")
	}

	vastUrl := os.Getenv("VAST_URL")
	if vastUrl == "" {
		t.Skip("VAST_URL environment variable not set")
	}
	vast := NewVast(vastKey, vastUrl)

	payload := &Payload{
		Abck:      "D1FD201C3035001E680880F88F59AFBC~-1~YAAQOgUXAmGAbQt/AQAAZCT/IQfrkiav1UEPnbfbDG7hozVfAVfzsekZFdp1EK8gR7sjYrERlBbQov/7roLqMPtfWqiciLmLSYjRgjuZkwBv3AvnBS1Elbcr1P8/o3t7gr+6+/3IeOcEo+68Jl4QKdmMG1Y71I+XY6STlILfP2Q4f+0IfeW2m4UHL0AMuaQKjDty7yf8Krh37oc2vtYoSf9ZLtjUQNuT7LskIgNkeKv9nM3DSv70q6rHZ+MWYLp4TrB+E1/fbkEfWtbUSjkWYt/SGohOLECh5O2zDPqHO0Mttjg6nmznI0rGmuxXIG3wewLfxUKLNQGZM7KKoItQVLRJrq2T78ACCeQXiJwyRYAOHkmHzYHfPWfBDEVPNm64HXj007OaCmE962GDTkwgq5bQK6qD7V4XHCohYT82sN0cCMblUlzpGTxxTp1cJhWXKoddBFCKk/bWMf6BL+eAmtDtC63qxvQH1idE2ZbIRG75tQAlA+DGc9V+Ab4YJMpG503jNw0x/XcPAJkfrUyY1EcDt2CsGFMj774moA==~-1~-1~-1",
		Url:       "https://www.nike.com/gb/",
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36",
	}

	sensor, err := vast.GetSensor(payload)
	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Logf("%+v", sensor)
}

func BenchmarkVast_GetSensor(b *testing.B) {
	vastKey := os.Getenv("VAST_KEY")
	if vastKey == "" {
		b.Skip("VAST_KEY environment variable not set")
	}

	vastUrl := os.Getenv("VAST_URL")
	if vastUrl == "" {
		b.Skip("VAST_URL environment variable not set")
	}
	vast := NewVast(vastKey, vastUrl)

	payload := &Payload{
		Abck:      "D1FD201C3035001E680880F88F59AFBC~-1~YAAQOgUXAmGAbQt/AQAAZCT/IQfrkiav1UEPnbfbDG7hozVfAVfzsekZFdp1EK8gR7sjYrERlBbQov/7roLqMPtfWqiciLmLSYjRgjuZkwBv3AvnBS1Elbcr1P8/o3t7gr+6+/3IeOcEo+68Jl4QKdmMG1Y71I+XY6STlILfP2Q4f+0IfeW2m4UHL0AMuaQKjDty7yf8Krh37oc2vtYoSf9ZLtjUQNuT7LskIgNkeKv9nM3DSv70q6rHZ+MWYLp4TrB+E1/fbkEfWtbUSjkWYt/SGohOLECh5O2zDPqHO0Mttjg6nmznI0rGmuxXIG3wewLfxUKLNQGZM7KKoItQVLRJrq2T78ACCeQXiJwyRYAOHkmHzYHfPWfBDEVPNm64HXj007OaCmE962GDTkwgq5bQK6qD7V4XHCohYT82sN0cCMblUlzpGTxxTp1cJhWXKoddBFCKk/bWMf6BL+eAmtDtC63qxvQH1idE2ZbIRG75tQAlA+DGc9V+Ab4YJMpG503jNw0x/XcPAJkfrUyY1EcDt2CsGFMj774moA==~-1~-1~-1",
		Url:       "https://www.nike.com/gb/",
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/99.0.4844.51 Safari/537.36",
	}

	for i := 0; i < b.N; i++ {
		sensor, err := vast.GetSensor(payload)
		if err != nil {
			b.Error(err)
		}

		b.Logf("Sensor Len: %d", len(sensor.Sensor))
	}

}
