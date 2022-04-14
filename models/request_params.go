package models

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type RequestParams struct {
	Length  int `form:"length", json"length"`
	Digits  int `form:"digits", json:"digits"`
	Symbols int `form:"symbols", json:"symbols"`
}

func (r *RequestParams) Validate() error {
	minLen := 16
	maxLen := 1024

	if r.Length == 0 {
		rand.Seed(time.Now().UnixNano())
		r.Length = minLen + rand.Intn(maxLen-minLen+1)
	} else if r.Length < 16 {
		return errors.New(fmt.Sprintf("Length (%d) is less than the minimum (16)", r.Length))
	}

	if r.Digits == 0 {
		r.Digits = int((r.Length / 3))
	} else if r.Digits < int((r.Length / 3)) {
		return errors.New(fmt.Sprintf("Digits (%d) is less than the minimum (16)", r.Digits))
	}

	if r.Symbols == 0 {
		r.Symbols = int((r.Length / 4))
	} else if r.Symbols < int((r.Length / 4)) {
		return errors.New(fmt.Sprintf("Symbols (%d) is less than the minimum (16)", r.Symbols))
	}
	return nil
}
