package ch11

import (
	"errors"
	"fmt"
	"testing"
)

var NegativeError = errors.New("v is negative")
var TooBigError = errors.New("v is too big")

func getDataFromDB(v int) (string, error) {
	//if v < 0 || v > 100 {
	//	return "", errors.New("out of range")
	//}
	if v < 0 {
		return "", NegativeError
	}
	if v > 100 {
		return "", TooBigError
	}
	return fmt.Sprintf("%d", v), nil
}

func TestGetData(t *testing.T) {
	if db, err := getDataFromDB(-29); err != nil {
		if errors.Is(err, NegativeError) {
			t.Error("Negative error expected")
		}
		t.Error(err)
	} else {
		t.Log(db)
	}
}
