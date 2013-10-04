package goexperiment

import (
  "testing"
)

func TestBloomFiterParameters (t *testing.T) {
  t.Log("This is a test.")
  filter := New(1000, 0.1)
  t.Log(filter.bytes)
}

func TestBloomFilterFilter(t *testing.T) {
  m, k := GetParameters(1000, 0.001)
  t.Logf("Parameters are m=%d, k=%d", m, k)
}

