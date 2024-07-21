package main

import "testing"

func TestQueue(t *testing.T) {
	t.Run("Add to queue", func(t *testing.T) {
		q := &Queue{}
		q.Enque("Alice")
		got, err := q.Front()
		want := "Alice"

		assertEquality(t, got, want)
		assertNoError(t, err)
	})

	t.Run("Get front from an empty queue", func(t *testing.T) {
		q := &Queue{}
		got, err := q.Front()
		want := ""

		assertEquality(t, got, want)
		assertError(t, err, ErrEmptyQueue)
	})

	t.Run("Remove from the queue", func(t *testing.T) {
		q := &Queue{}
		q.Enque("Alice")
		got, err := q.Deque()
		want := "Alice"
		
		assertEquality(t, got, want)
		assertNoError(t, err)
	})

	t.Run("Remove from an empty queue", func(t *testing.T) {
		q := &Queue{}

		got, err := q.Deque()
		want := ""

		assertEquality(t, got, want)
		assertError(t, err, ErrEmptyQueue)
	})

	t.Run("Multiple Operations", func(t *testing.T) {
		q := &Queue{}
		
		q.Enque("Ahmad")
		q.Enque("Amr")
		got, err := q.Deque()
		want := "Ahmad"

		assertEquality(t, got, want)
		assertNoError(t, err)
		q.Enque("Salem")
		got, err = q.Deque()
		want = "Amr"
		assertEquality(t, got, want)
		assertNoError(t, err)
		got, err = q.Deque()
		want = "Salem"
		assertEquality(t, got, want)
		assertNoError(t, err)
		got, err = q.Deque()
		
		assertEquality(t, got, "")
		assertError(t, err, ErrEmptyQueue)
	})

	t.Run("Test Empty", func(t *testing.T) {
		t.Run("Non Empty queue", func(t *testing.T) {
			q := &Queue{}

			q.Enque("Ahmad")
			got := q.Empty()
			
			assertBooleanEquality(t, got, false)
		})	

		t.Run("Empty queue", func(t *testing.T) {
			q := &Queue{}
			assertBooleanEquality(t, q.Empty(), true)
		})
	})
}

func assertBooleanEquality(t testing.TB, got, want bool) {
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}

func assertEquality(t testing.TB, got, want string) {
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	if got == nil {
		t.Errorf("got \"nil\" want %q", want.Error())
	} else if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	if got != nil {
		t.Errorf("got %v want nil", got.Error())
	}
}