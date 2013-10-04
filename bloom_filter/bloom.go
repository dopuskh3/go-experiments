package goexperiment

import (
  "hash"
  "encoding/binary"
  //"hash/fnv"
  "github.com/spaolacci/murmur3"
  "math"
)

// The bloom filter structure
type BloomFilter struct {
  m uint
  k uint
  h hash.Hash64
  bytes []byte
}

// Returns a new BloomFilter
func New(maxSize uint, fpProb float64) *BloomFilter {
  // TODO: error if fpProb > 1
  m, k := GetParameters(maxSize, fpProb)
  b := make([]byte, m/8)
  return &BloomFilter{m, k, murmur3.New64(), b}
}

// Tests for `input` availability in the bloom filter.
func (f *BloomFilter) test(input []byte) bool {
  return false
}

func (f *BloomFilter) getHash(input []byte) (a uint32, b uint32){
  // to implement
  f.h.Reset()
  f.h.Write(input)
  // Converts hash value to from uint64 to []byte 
  // and to uint32 again.
  vbuff := make([]byte, 8)
  binary.PutUvarint(vbuff, f.h.Sum64())
  upper := vbuff[0:4]
  lower := vbuff[4:8]
  return binary.BigEndian.Uint32(lower), binary.BigEndian.Uint32(upper)
}

// Adds `input` into the bloom filter.
func (f *BloomFilter) add(input []byte) *BloomFilter {
  var loc uint32
  a, b := f.getHash(input)
  for i := uint(0); i <  f.k; i++ {
    // Location for this hash in the filter
    loc = ( a + b * uint32(i) ) % uint32(f.m)
    f.bytes[loc].Set()
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


//func (t *BloomFilter) computeHash(input []byte) uint32 {
//  return 10;
//}


