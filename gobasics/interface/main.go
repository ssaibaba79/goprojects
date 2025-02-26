package main

import (
	"fmt"
)

type Encoder interface {
	encode()
}

type Base64Encoder struct{}

type URLEncoder struct{}

func (b *Base64Encoder) encode() {
	fmt.Println("Base64 Encoded")
}

func (u *URLEncoder) encode() {
	fmt.Println("URL Encoded")
}

var encoders = map[string]Encoder{
	"base64": &Base64Encoder{},
	"url":    &URLEncoder{},
}

func main() {

	testEncoders := []string{"test", "base64", "url"}

	for _, e := range testEncoders {
		encoder, err := NewEncoder(e)
		if err != nil {
			fmt.Println(err)
		} else {
			encoder.encode()
		}
	}
}

func NewEncoder(name string) (Encoder, error) {
	if v, ok := encoders[name]; ok {
		return v, nil
	}

	return nil, fmt.Errorf("%s is not valid encoder type", name)
}
