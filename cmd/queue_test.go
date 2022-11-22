package cmd

import "testing"

func TestQueue(t *testing.T) {
	t.Run("Test isEmpty throws true if Queue has no elements", func(t *testing.T) {
		q := &Queue{}
		if !q.isEmpty() {
			t.Errorf("Empty Queue returned false for isEmpty check")
		}
	})

	t.Run("Test enqueuing one element returns isEmpty false", func(t *testing.T) {
		q := &Queue{}
		q.enqueue("blueberry")
		if q.isEmpty() {
			t.Errorf("Queue has one element, should return false for isEmpty check")
		}
	})

	t.Run("Test enqueuing and dequeing", func(t *testing.T) {
		q := &Queue{}
		q.enqueue("blueberry")
		q.enqueue("strawberry")
		q.enqueue("blackberry")

		topOfQueue, _ := q.dequeue()

		if topOfQueue != "blueberry" {
			t.Errorf("Expected blueberry to be removed from queue, instead got %v", topOfQueue)
		}
	})
}
