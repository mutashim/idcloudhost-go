package idcloudhost

import (
	"fmt"
	"testing"
)

func TestParseForm(t *testing.T) {

	req := VMInput{
		Backup: false,
		Name:   "ContohVM",
	}

	data, _ := parseForm(req)
	fmt.Println(data)
}
