package history

import (
	"encoding/json"
	"os"
)

type Stats struct {
	TotalGames int `json:"total_games"`
	Wins       int `json:"wins"`
	Losses     int `json:"losses"`
}

const statsFile = "stats.json"

// Load reads stats from the file. If the file does not exist,
// it returns a zero-value Stats.
func Load() (*Stats, error) {
	data, err := os.ReadFile(statsFile)
	if err != nil {
		if os.IsNotExist(err) {
			return &Stats{}, nil
		}

		return nil, err
	}

	var s Stats
	err = json.Unmarshal(data, &s)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

// Save writes the stats to the file.
func (s *Stats) Save() error {
	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(statsFile, data, 0644)
}
