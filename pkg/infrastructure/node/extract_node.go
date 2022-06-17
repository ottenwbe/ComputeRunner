package node

import (
	"errors"
	"github.com/sirupsen/logrus"
	"reflect"

	"encoding/json"
	"github.com/robertkrimen/otto"
)

// ExtractNodeStrings converts ("javascript") call parameters to a description. Returns an error when no conversion is possible.
func ExtractNodeStrings(call otto.FunctionCall) (description *Description, err error) {
	if len(call.ArgumentList) < 1 {
		err = errors.New("not enough arguments in node description")
	} else {
		if isJSONString(call) {
			description = &Description{}
			err = json.Unmarshal([]byte(call.Argument(0).String()), description)
			logrus.Debug("Extract Node Description (String): " + call.Argument(0).String())
		} else if isJavaScriptObject(call) {
			description, err = extractFromObject(call)
			logrus.Debug("Extract Node Description (Object): " + description.Name)
		}
	}
	return description, err
}

func isJavaScriptObject(call otto.FunctionCall) bool {
	return call.Argument(0).IsObject()
}

func isJSONString(call otto.FunctionCall) bool {
	return call.Argument(0).IsString()
}

func extractFromObject(call otto.FunctionCall) (*Description, error) {
	var (
		description = &Description{}
		err         error
	)
	descriptionValue := reflect.ValueOf(description).Elem()
	forEachDescriptionValue(
		func(index int) {
			extractDescriptionValueFromJSON(call, descriptionValue, index, description)
		},
		descriptionValue.NumField())
	return description, err
}

func forEachDescriptionValue(execute func(int), max int) {
	for i := 0; i < max; i++ {
		execute(i)
	}
}

func extractDescriptionValueFromJSON(call otto.FunctionCall, descriptionValue reflect.Value, i int, description *Description) {
	JSONKey := getJSONKeyFromDescriptionType(descriptionValue, i)
	val, err := call.Argument(0).Object().Get(JSONKey)
	if err != nil {
		logrus.Error(err)
		logrus.Info(JSONKey)
	} else if valueIsDefined(val) {
		logrus.Info("Setting: " + JSONKey + " " + val.String())
		if val.IsString() {
			var str string
			str, err = val.ToString()
			reflect.ValueOf(description).Elem().Field(i).SetString(str)
		} else if val.IsNumber() {
			var num int64
			num, err = val.ToInteger()
			reflect.ValueOf(description).Elem().Field(i).SetInt(num)
		}
	}
}

func valueIsDefined(val otto.Value) bool {
	return !val.IsUndefined()
}

func getJSONKeyFromDescriptionType(descriptionValue reflect.Value, i int) string {
	return descriptionValue.Type().Field(i).Tag.Get("json")
}
