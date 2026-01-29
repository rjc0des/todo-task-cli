package store

import (
	"encoding/json"
	"os"
	"task-cli/internal/model"
)

var FileName = "tasks.json"

type Data struct {
	LastId int          `json:"last_id"`
	Tasks  []model.Task `json:"tasks"`
}

func Save(db Data) error {
	data, err := json.MarshalIndent(db, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(FileName, data, 0644)
}

func Load() (Data, error) {
	if _, err := os.Stat(FileName); os.IsNotExist(err) {
		return Data{LastId: 0, Tasks: []model.Task{}}, nil
	}

	data, err := os.ReadFile(FileName)

	if err != nil {
		return Data{}, err
	}

	var db Data
	err = json.Unmarshal(data, &db)

	if err != nil {
		return Data{}, err
	}

	return db, nil
}
