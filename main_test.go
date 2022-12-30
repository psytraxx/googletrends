package main

import (
	"testing"
)

func Test_readGoogleTrends(t *testing.T) {

	t.Run("Load Swiss Trends", func(t *testing.T) {
		if got, _ := readGoogleTrends("CH"); !(len(got) > 1) {
			t.Errorf("readGoogleTrends() = %v", got)
		}
	})

	t.Run("Load Invalid Trends", func(t *testing.T) {
		if _, err := readGoogleTrends("XX"); !(err != nil) {
			t.Errorf("readGoogleTrends() = %v", err)
		}
	})
}
