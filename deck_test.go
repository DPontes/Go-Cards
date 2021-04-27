package main

import "testing"

func TestNewDeck(t *testing.T) {
  d := newDeck()

  if len(d) != 16 {     // 16 = 4 suits x 4 values
    t.Errorf("Expected deck length of 16, but got %v", len(d))
  }
}
