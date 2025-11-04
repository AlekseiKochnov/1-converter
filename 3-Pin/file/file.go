package file

import (
	"errors"
	"os"
	"strings"
)

type ReadFile struct{}

type FileServis interface{
	ReadJsonFile(string)([]byte, error)
}

func ReadIF(f FileServis, nameFile string) ([]byte, error) {
	return f.ReadJsonFile(nameFile)
}

func(r *ReadFile) ReadJsonFile(nameFile string) ([]byte, error) {

    if !strings.Contains(nameFile, ".json") {
		return nil, errors.New("расширение файла не json")
    }	

	jsonData, err := os.ReadFile(nameFile)
	if err != nil {
		return nil, err
	}

	return jsonData, nil

}