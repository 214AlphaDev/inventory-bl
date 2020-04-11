package utils

import (
	"fmt"
	"reflect"
)

type IInitialized interface {
	Initialized() bool
}

func Initialized(initializedElements ...IInitialized) error {

	for _, element := range initializedElements {

		if element == nil || reflect.ValueOf(element).Kind() == reflect.Ptr && reflect.ValueOf(element).IsNil() {
			continue
		}

		if !element.Initialized() {

			elementType := reflect.TypeOf(element)

			switch elementType.Kind() {
			case reflect.Ptr:
				return fmt.Errorf("element '%s' has not been initialized", elementType.Elem().Name())
			default:
				return fmt.Errorf("element '%s' has not been initialized", elementType.Name())
			}

		}

	}

	return nil

}
