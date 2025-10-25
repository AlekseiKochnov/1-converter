package main

import (
	"errors"
	"time"
)

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name string
}

type BinList []Bin

func NewBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {
	if id == "" || name == "" {
		return nil, errors.New("empty string (id or name)")
	}

	b := Bin{
		Id: id,
		Private: private,
		CreatedAt: createdAt, 
		Name: name,
	}
	return &b, nil
}

func NewBinList() *BinList {
	return &BinList{}
}

func main() {


}