package people

import (
	"encoding/json"
	"fmt"
	"net/http"
)

	
type Astros struct {
	People []struct {
		Craft string `json:"craft"`
		Name  string `json:"name"`
	} `json:"people"`
	Number  int    `json:"number"`
	Message string `json:"message"`
}

func FetchData(url string) (*Astros, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error fetching data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error fetching data: %v", resp.StatusCode)
	}

	var result Astros
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("Error decoding data: %v", err)
	}

	return &result, nil
}
