package main

import (
	"3-PIN/api"
	"3-PIN/bins"
	"3-PIN/file"
	"3-PIN/storage"
	"encoding/json"
	"fmt"
	"time"
)

func main() {

	api, err := api.NewApi()
	if err != nil {
		fmt.Println(err)
		return
	}
	key := api.GetKye()
	fmt.Println(key)

	bin, err := bins.NewBin("1", false, time.Now(), "one")
	if err != nil {
		fmt.Println(err)
		return
	}
	List := bins.NewBinList()
	binList := append(*List, *bin)
	data, err := json.Marshal(binList)
	if err != nil {
		fmt.Println(err)
		return
	}	

	err = storage.SaveIF(&storage.Stor{}, data)
	if err != nil {
		fmt.Println(err)
		return
	}

	js, err := file.ReadIF(&file.ReadFile{}, "Bin.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(js))

}