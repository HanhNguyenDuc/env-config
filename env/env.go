package env

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type EnvParser struct {
	ParserMap map[string]ParseObject
}

func NewEnvParser() *EnvParser {
	return &EnvParser{
		ParserMap: DefaultCallbackMap,
	}
}

// Currently support string, int, float
func (e *EnvParser) ParseFromEnvs(x interface{}) error {
	// Get all FieldNames
	v := reflect.ValueOf(x).Elem()

	data := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		value := reflect.ValueOf(x).Elem().Field(i)
		envName := field.Tag.Get("env")
		jsonType := field.Type.Name()
		envString := os.Getenv(envName)

		parser := e.ParserMap[jsonType]
		if parser == nil {
			continue
		}
		err := parser.Parse(envString, &value)
		if err != nil {
			return err
		}

		// switch jsonType {
		// case "int":
		// 	intVal, err := strconv.Atoi(envString)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	value.SetInt(int64(intVal))
		// case "string":
		// 	value.SetString(envString)
		// case "int64":
		// 	intVal, err := strconv.ParseInt(envString, 10, 64)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	value.SetInt(intVal)
		// case "float32":
		// 	floatVal, err := strconv.ParseFloat(envString, 32)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	value.SetFloat(floatVal)
		// case "float64":
		// 	floatVal, err := strconv.ParseFloat(envString, 64)
		// 	if err != nil {
		// 		return err
		// 	}
		// 	value.SetFloat(floatVal)
		// default:
		// 	continue
		// }
	}

	// map -> string
	jsonString, err := json.Marshal(&data)
	if err != nil {
		return err
	}

	// string -> object
	err = json.Unmarshal(jsonString, x)
	if err != nil {
		return err
	}

	fmt.Printf("x: %v\n", x)

	return nil
}
