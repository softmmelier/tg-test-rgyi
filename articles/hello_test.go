package articles

import (
	"context"
	"testing"
)

func TestWorld(t *testing.T) {
	resp, err := World(context.Background(), "Jane Doe")
	if err != nil {
		t.Fatal(err)
	}
	want := "Articles there, Jane Doe!"
	if got := resp.Message; got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
