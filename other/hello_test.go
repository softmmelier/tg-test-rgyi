package other_test

import (
  "context"
  "testing"
  
  "encore.app/other"
)

func TestWorld(t *testing.T) {
  resp, err := other.World(context.Background(), "Jane Doe")
  if err != nil {
    t.Fatal(err)
  }
  want := "Other there, Jane Doe!"
  if got := resp.Message; got != want {
    t.Errorf("got %q, want %q", got, want)
  }
}
