package age

import "fmt"

//go:generate go run github.com/vektra/mockery/v2@v2.28.2 --name=AgeGetter
type AgeGetter interface {
	GetAge(firstName string) (int, error)
}

type API struct{}

type Response struct {
	Text  string `json:"text"`
	Error string `json:"error"`
}

func Age(service AgeGetter, name string) (int, error) {
	data, err := service.GetAge(name)
	if err != nil {
		return -1, fmt.Errorf("error: %w", err)
	}

	return data, nil
}
