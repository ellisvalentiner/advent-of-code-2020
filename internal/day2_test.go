package internal_test

import (
  "testing"
  "github.com/ellisvalentiner/advent-of-code-2020/internal"
)

func TestDay2(t *testing.T) {
  part1, part2 := internal.Day2()
  if part1 != 636 {
    t.Errorf("Expected 636, got %q", part1)
  }
  if part2 != 588 {
    t.Errorf("Expected 588, got %q", part2)
  }
}
