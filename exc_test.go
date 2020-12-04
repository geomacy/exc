package main

import (
	"reflect"
	"testing"
	"time"
)

func TestProcess(t *testing.T) {
	counts := make(map[string]int, 0)
	err := process("data.txt", counts)
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]int{
		"alice":     1,
		"cried":     1,
		"and":       1,
		"curiouser": 2,
	}
	if !reflect.DeepEqual(counts, expected) {
		t.Fatalf("expected %v but got %v", expected, counts)
	}
}

func TestProcessAll(t *testing.T) {
	counts := make(map[string]int, 0)
	err := processAll([]string{"data.txt", "data.txt"}, counts)
	if err != nil {
		t.Fatal(err)
	}
	time.Sleep(2 * time.Second)
	expected := map[string]int{
		"alice":     2,
		"cried":     2,
		"and":       2,
		"curiouser": 4,
	}
	if !reflect.DeepEqual(counts, expected) {
		t.Fatalf("expected %v but got %v", expected, counts)
	}
}
