package re

import (
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Expected: %s (type %v)\n Actual: %s (type %v)", expected, reflect.TypeOf(expected),
      actual, reflect.TypeOf(actual))
	}
}