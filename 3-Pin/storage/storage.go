package storage

import (
	"3-PIN/bins"
	"encoding/json"
	"os"
)

type Stor struct{}

type StorageServis interface{
	SaveBin([]byte)error
}

func SaveIF(st StorageServis, data []byte) error {
	return st.SaveBin(data)
}

func(s *Stor) SaveBin(data []byte) error{

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