package storage

import (
	"encoding/json"
	"os"
)

func (s *Storage) Persist() {
	jsonString, _ := json.Marshal(s.Tasks)
	os.WriteFile("data/data.json", jsonString, os.ModePerm)
}
