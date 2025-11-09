package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Konf struct {
	Key string
}

func NewKonf() (*Konf, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Не удалось найти env файл")

	}

	key := os.Getenv("KEY")

	if key == "" {
		return nil, fmt.Errorf("Не передан параметр KEY в переменные окружения")
	}

	return &Konf{Key: key}, nil
}