package pkg

import (
	"encoding/json"
	"log"
	"main/pkg/models"
	"os"
)

//var ID []int

func CompareID() (int, error) {
	var notes []models.Note
	bytes, err := os.ReadFile("notes.json")
	if err != nil {
		log.Println(err) //todo убрать логи
		return 0, err
	}
	if len(bytes) == 0 {
		return 1, nil
	}
	err = json.Unmarshal(bytes, &notes)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return notes[len(notes)-1].Id + 1, nil
}
