package goexperiment

import (
  "testing"
)

func TestBloomFiterParameters (t *testing.T) {
  filter := New(1000, 0.1)
  t.Logf("Filter parameters m=%d k=%d" , filter.m, filter.k)
}

func TestBloomFilterFilter(t *testing.T) {
  m, k := GetParameters(1000, 0.001)
  t.Logf("Parameters are m=%d, k=%d", m, k)
}

func TestBloomFilterKeyIsPresent(t *testing.T) {
  filter := New(1000, 0.001)
  filter.Add([]byte("FOO"))
  if filter.Test([]byte("BAR")) {
    t.Error("BAR is in set and should not be in filter.")
  }
  if filter.Test([]byte("FOO")) {
    t.Logf("You're lucky, FOO is in filter")
  }
}
