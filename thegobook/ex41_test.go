package main

import "testing"
import "reflect"

func TestRemoveAdjacent(t *testing.T) {
	input := []string{"a", "b", "c", "c", "c", "d"}
	expected := []string{"a", "b", "c", "d"}
	actual := RemoveAdjacent(input)
	if !reflect.DeepEqual(actual, expected) {
		t.Error(`RemoveAdjacent() should remove adjacent`, actual, expected)
	}
}
