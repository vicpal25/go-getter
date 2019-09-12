package main

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	type args struct {
		c Auto
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Generate(tt.args.c)
		})
	}
}
