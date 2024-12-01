package main

import (
	"reflect"
	"testing"
)

func TestParseInputStr(t *testing.T) {
	tests := []struct {
		name     string
		inputStr string
		expected operation
	}{
		{
			name:     "Parse 'turn on'",
			inputStr: "turn on 887,9 through 959,629",
			expected: operation{what: "on", x: struct {
				from int
				to   int
			}{from: 887, to: 959}, y: struct {
				from int
				to   int
			}{from: 9, to: 629}},
		},
		{
			name:     "Parse 'turn off'",
			inputStr: "turn off 539,243 through 559,965",
			expected: operation{what: "off", x: struct {
				from int
				to   int
			}{from: 539, to: 559}, y: struct {
				from int
				to   int
			}{from: 243, to: 965}},
		},
		{
			name:     "Parse 'toggle'",
			inputStr: "toggle 720,196 through 897,994",
			expected: operation{what: "toggle", x: struct {
				from int
				to   int
			}{from: 720, to: 897}, y: struct {
				from int
				to   int
			}{from: 196, to: 994}},
		},
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := parseInputStr(tc.inputStr)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected: %v, got: %v", i, tc.name, tc.expected, actual)
			}
		})
	}

}
