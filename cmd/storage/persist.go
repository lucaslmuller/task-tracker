package storage

import (
	"encoding/json"
	"os"
)

func (s *Storage) Persist() error {
	jsonString, err := json.Marshal(s.Tasks)

	if err != nil {
		return err
	}

	return os.WriteFile("data/data.json", jsonString, os.ModePerm)
}
