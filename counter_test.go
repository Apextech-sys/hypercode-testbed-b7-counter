package counter

import "testing"

func TestIncrementAndGet(t *testing.T) {
	s := NewStore()
	s.Increment("a")
	s.Increment("a")
	s.Increment("b")
	if got := s.Get("a"); got != 2 {
		t.Fatalf("Get(a) = %d, want 2", got)
	}
	if got := s.Get("b"); got != 1 {
		t.Fatalf("Get(b) = %d, want 1", got)
	}
}
