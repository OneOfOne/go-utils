package hash

const prime = 2862933555777941757

// JumpConsistentHash a fast, minimal memory, consistent hash algorithm based on the
// [paper](http://arxiv.org/pdf/1406.2294v1.pdf) by John Lamping and Eric Veach.
func JumpConsistentHash(key uint64, buckets int) (b int) {
	j := 0
	for j < buckets {
		b = j
		key = key*prime + 1
		j = int(float64(b+1) * (float64(1<<31) / float64(key>>33+1)))
	}
	return b
}
