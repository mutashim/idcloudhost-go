package idcloudhost

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
)

func parseForm(in interface{}) (*url.Values, error) {
	v := reflect.ValueOf(in)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, errors.New("input is not struct")
	}

	typ := v.Type()
	out := url.Values{}
	for i := 0; i < v.NumField(); i++ {
		f := typ.Field(i)
		if tagv := f.Tag.Get("form"); tagv != "" {
			if fieldv := fmt.Sprint(v.Field(i).Interface()); fieldv != "" {
				out.Add(tagv, fieldv)
			}
		}
	}

	return &out, nil
}
