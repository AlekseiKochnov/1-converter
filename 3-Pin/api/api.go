package api

import "3-PIN/config"

type Api struct {
	Konf config.Konf
}

func NewApi() (*Api, error) {
	konf, err := config.NewKonf()
	if err != nil {
		return nil, err
	}
	return &Api{Konf: *konf}, nil
}

func(a *Api) GetKye() string {
	return a.Konf.Key
}