package goexperiment

import (
  "hash"
  "github.com/spaolacci/murmur3"
  "github.com/willf/bitset"
  "math"
)

// The bloom filter structure
type BloomFilter struct {
  m uint
  k uint
  h hash.Hash64
  bits *bitset.BitSet
}

// Returns a new BloomFilter
func New(maxSize uint, fpProb float64) *BloomFilter {
  // TODO: error if fpProb > 1
  m, k := GetParameters(maxSize, fpProb)
  b := bitset.New(m)
  return &BloomFilter{m, k, murmur3.New64(), b}
}

// Tests for `input` availability in the bloom filter.
func (f *BloomFilter) Test(input []byte) bool {
  a, b := f.getHash(input)
  for i := uint(0); i <  f.k; i++ {
    // Location for this hash in the filter
    loc := ( a + b * uint32(i) ) % uint32(f.m)
    if !f.bits.Test(uint(loc)){
      return false
    }
  }
  return true
}

func (f *BloomFilter) getHash(input []byte) (a uint32, b uint32){
  f.h.Reset()
  f.h.Write(input)
  hash := f.h.Sum64()
  upper := uint32( (hash>>4) & 0x0000FFFF)
  lower := uint32( hash & 0x0000FFFF )
  return lower, upper
}

// Adds `input` into the bloom filter.
func (f *BloomFilter) Add(input []byte) *BloomFilter {
  var loc uint32
  a, b := f.getHash(input)
  for i := uint(0); i <  f.k; i++ {
    // Location for this hash in the filter
    loc = ( a + b * uint32(i) ) % uint32(f.m)
    f.bits.Set(uint(loc))
  }
  return f
}

// Takes max bloom filter size and false positive probability 
// and compute the following values:
//
// Parameters:
//  maxSize: the maximum bloom filter size
//  fpProb: the false positive probability
// Returns:
//  m: the bloom filter size in bits
//  k: the number of hash function to use
func GetParameters(maxSize uint, fpProb float64) (m uint, k uint) {
  m = uint(-1 * (float64(maxSize) * math.Log(fpProb)) / math.Pow(math.Log(2), 2))
  k = uint((float64(m) / float64(maxSize)) * math.Log(2))
  return
}




