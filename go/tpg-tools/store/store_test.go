package store_test

import (
	"testing"

	"github.com/rwx-yxu/lab/go/tpg-tools/store"
	"golang.org/x/exp/slices"
)

func TestStoreFile(t *testing.T) {
	t.Parallel()
	path := t.TempDir() + "/store.bin"
	output := store.Open(path)
	want := []int{2, 3, 5, 7, 11}
	err := output.Save(want)
	if err != nil {
		t.Fatal(err)
	}
	output.Close()
	input := store.Open(path)
	var got []int
	err = input.Load(&got)
	if err != nil {
		t.Fatal(err)
	}
	if !slices.Equal(want, got) {
		t.Errorf("want %q, got %q", want, got)
	}
}
