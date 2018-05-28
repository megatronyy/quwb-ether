package event

import "testing"

type testEvent int

func TestSubCloseUnsub(t *testing.T) {
	var mux TypeMux
	mux.Stop()
	sub := mux.Subscribe(int(0))
	sub.Unsubscribe()
}

func TestSub(t *testing.T) {
	mux := new(TypeMux)
	defer mux.Stop()

	sub := mux.Subscribe(testEvent(0))
	go func() {
		if err := mux.Post(testEvent(5)); err != nil {
			t.Errorf("Post returned unexpected error: %v", err)
		}
	}()

	ev := <-sub.Chan()

	if ev.Data.(testEvent) != testEvent(5) {
		t.Errorf("Got %v (%T), expected event %v (%T)", ev, ev, testEvent(5), testEvent(5))
	}
}
