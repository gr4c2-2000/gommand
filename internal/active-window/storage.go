package activewindow

import (
	"encoding/json"
	"log"
	"os"
)

type Storage struct {
	FilePath string
	Items    []StorageItem
}

type StorageItem struct {
	SearchWord string `json:"search_word"`
	LastId     uint32 `json:"last_id"`
}

func InitStorage(filePath string) (*Storage, error) {
	s := Storage{}
	s.FilePath = filePath
	_, err := os.Stat(s.FilePath)
	if os.IsNotExist(err) {
		// File does not exist, create an empty storage
		s.Items = []StorageItem{}
		return &s, nil
	}
	return s.readFromStorage()
}

func (s *Storage) Find(appName string) *StorageItem {

	for i, item := range s.Items {
		if item.SearchWord == appName {
			return &s.Items[i]
		}
	}
	return nil
}

func (s *Storage) Replace(appName string, id uint32) {
	item := s.Find(appName)
	if item == nil {
		s.Items = append(s.Items, StorageItem{appName, id})
		log.Default().Println(s.SaveToFile())
		return
	}
	item.LastId = id
	log.Default().Println(s.SaveToFile())
}

func (s *Storage) readFromStorage() (*Storage, error) {
	// Read file content
	fileContent, err := os.ReadFile(s.FilePath)
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON into Storage
	err = json.Unmarshal(fileContent, &s.Items)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Storage) SaveToFile() error {
	data, err := json.MarshalIndent(s.Items, "", "    ")
	if err != nil {
		return err
	}

	err = os.WriteFile(s.FilePath, data, 0775)
	if err != nil {
		return err
	}

	return nil
}
