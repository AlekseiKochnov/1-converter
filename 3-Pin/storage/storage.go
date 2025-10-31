package storage

import (
	"3-PIN/bins"
	"encoding/json"
	"os"
)

func SaveBin(data []byte) error{

	file, err := os.Create("Bin.json")
	if err != nil {
		return err
	}
	defer file.Close()	

	_, err = file.Write(data)
	if err != nil {
		return err
	}	

	return nil

}

func ReadBin() (*bins.BinList, error) {
	readBin := bins.NewBinList()

	jsonData, err := os.ReadFile("Bin.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(jsonData, readBin)
	if err != nil {
		return nil, err
	}	

	return readBin, nil
	
}