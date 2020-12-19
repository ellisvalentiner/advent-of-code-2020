package internal_test

import (
  "testing"
  "github.com/ellisvalentiner/advent-of-code-2020/internal"
)

func TestDay1(t *testing.T) {
  part1, part2 := internal.Day1()
  if part1 != 138379 {
    t.Errorf("Expected 138379, got %q", part1)
  }
  if part2 != 85491920 {
    t.Errorf("Expected 85491920, got %q", part2)
  }
}
