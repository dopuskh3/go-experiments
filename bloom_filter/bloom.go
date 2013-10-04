package goexperiment

import (
  "hash"
  "hash/fnv"
  "math"
)


type BloomFilter struct {
  m uint
  k uint
  h hash.Hash64
  bytes []byte
}


func New(maxSize uint, fpProb float64) *BloomFilter {
  // TODO: error if fpProb > 1
  m, k := getParameters(maxSize, fpProb)
  b := make([]byte, m/8)
  return &BloomFilter{m, k, fnv.New64(), b}
}

func (f *BloomFilter) test(input []byte) bool {
  // to implement

  return true;
}

func (f *BloomFilter) add(input []byte) {

}


func getParameters(maxSize uint, fpProb float64) (m uint, k uint) {
  m = uint(-1 * (float64(maxSize) * math.Log(fpProb)) / math.Pow(math.Log(2), 2))
  k = uint((float64(m) / float64(maxSize)) * math.Log(2))
  return
}


//func (t *BloomFilter) computeHash(input []byte) uint32 {
//  return 10;
//}


