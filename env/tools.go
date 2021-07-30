package env

import (
	"fmt"
	"reflect"
	"strconv"
)

var ErrTypeNotImplemented error = fmt.Errorf("this type is not implemented")

type ParseObject interface {
	Parse(stringValue string, ref *reflect.Value) error
}

type ParseInt64 struct{}

func (pi *ParseInt64) Parse(stringValue string, ref *reflect.Value) error {
	intVal, err := strconv.ParseInt(stringValue, 10, 64)
	if err != nil {
		return err
	}
	ref.SetInt(intVal)
	return nil
}

type ParseInt struct{}

func (pi *ParseInt) Parse(stringValue string, ref *reflect.Value) error {
	intVal, err := strconv.Atoi(stringValue)
	if err != nil {
		return err
	}
	ref.SetInt(int64(intVal))
	return nil
}

type ParseString struct{}

func (pi *ParseString) Parse(stringValue string, ref *reflect.Value) error {
	ref.SetString(stringValue)
	return nil
}

type ParseFloat32 struct{}

func (pi *ParseFloat32) Parse(stringValue string, ref *reflect.Value) error {
	floatVal, err := strconv.ParseFloat(stringValue, 32)
	if err != nil {
		return err
	}
	ref.SetFloat(floatVal)
	return nil
}

type ParseFloat64 struct{}

func (pi *ParseFloat64) Parse(stringValue string, ref *reflect.Value) error {
	floatVal, err := strconv.ParseFloat(stringValue, 32)
	if err != nil {
		return err
	}
	ref.SetFloat(floatVal)

	return nil
}

var DefaultCallbackMap map[string]ParseObject = map[string]ParseObject{
	"int":     &ParseInt{},
	"int64":   &ParseInt64{},
	"string":  &ParseString{},
	"float32": &ParseFloat32{},
	"float64": &ParseFloat64{},
}
