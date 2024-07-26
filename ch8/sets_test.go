package main

import "testing"

func TestSets(t *testing.T) {
	set1 := ToSet([]string{"apple", "orange", "mango", "banana"})
	set2 := ToSet([]string{"grapes", "watermelon", "banana", "orange", "avocado"})
	t.Run("Union", func(t *testing.T) {
		want := ToSet([]string{"apple", "orange", "mango", "banana", "watermelon", "avocado", "grapes"})
		got := Union(set1, set2)

		assertEquality(t, got, want)
	})

	t.Run("Intersection", func(t *testing.T) {
		want := ToSet([]string{"orange", "banana"})
		got := Intersect(set1, set2)

		assertEquality(t, got, want)
	})

	t.Run("Difference", func(t *testing.T) {
		t.Run("set1-set2", func(t *testing.T) {
			want := ToSet([]string{"apple", "mango"})
			got := Difference(set1, set2)
			assertEquality(t, got, want)
		})
		t.Run("set2-set1", func(t *testing.T) {
			want := ToSet([]string{"grapes", "watermelon", "avocado"})
			got := Difference(set2, set1)
			assertEquality(t, got, want)

		})
	})
}

func assertEquality(t testing.TB, got, want set) {
	t.Helper()
	if !Equal(got, want) {
		t.Errorf("\ngot %v \nwant %v", got, want)
	}
}
