package file

import (
	"errors"
	"os"
	"strings"
)

func ReadJsonFile(nameFile string) ([]byte, error) {

    if !strings.Contains(nameFile, ".json") {
		return nil, errors.New("расширение файла не json")
    }	

	jsonData, err := os.ReadFile(nameFile)
	if err != nil {
		return nil, err
	}

	return jsonData, nil

}