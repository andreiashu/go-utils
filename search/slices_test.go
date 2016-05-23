package search

import "testing"

func TestStringSlicesEqual(t *testing.T) {
	a := []string{
		"test",
		"test2",
		"12059125",
	}
	b := []string{
		"test",
		"test2",
		"12059125",
	}
	// test that the order of the slices is not relevant
	if !StringSlicesEqual(a, []string{"test2", "test", "12059125"}) {
		t.Errorf("Slices are equal")
	}
	// test that the comparison is case sensitive
	if StringSlicesEqual(a, []string{"TEST", "test2", "12059125"}) {
		t.Errorf("Slices are equal")
	}

	c := []string{
		"test",
		"test2",
	}
	if StringSlicesEqual(a, c) || StringSlicesEqual(b, c) {
		t.Errorf("Slices are not equal")
	}

	d := []string{}
	if StringSlicesEqual(a, d) || StringSlicesEqual(b, d) {
		t.Errorf("Slices are not equal")
	}
}

func TestStringInSlice(t *testing.T) {
	ss := []string{
		"test",
		"test2",
		"asfd gfh",
		"12059125",
	}

	for _, v := range ss {
		if !StringInSlice(v, ss) {
			t.Errorf("Expected to find %s in slice", v)
		}
	}

	na := []string{
		"non existing",
		"",
		"asfd",
	}
	for _, n := range na {
		if StringInSlice(n, ss) {
			t.Errorf("Didn't expecte to find %s in slice", n)
		}
	}
}
