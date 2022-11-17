package cmd

import "testing"

func TestStack(t *testing.T) {
	t.Run("Test isEmpty throws true if Stack has no elements", func(t *testing.T) {
		s := Stack{}
		if !s.isEmpty() {
			t.Errorf("Empty Stack returned false for isEmpty check")
		}
	})

	t.Run("Test pushing one element onto stack returns isEmpty false", func(t *testing.T) {
		s := Stack{}
		s.push("apple")
		if s.isEmpty() {
			t.Errorf("Stack has one element pushed onto it, should return false for isEmpty")
		}
	})

	t.Run("Test pushing and popping off a stack", func(t *testing.T) {
		s := Stack{}
		s.push("apple")
		s.push("orange")
		s.push("banana")

		firstElem, _ := s.pop()

		if firstElem != "banana" {
			t.Errorf("Expected banana to be popped off first from stack instead got %v", firstElem)
		}
	})
}
