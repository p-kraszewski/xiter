package xiter

import (
	"iter"
)

// Seq2ToSeqKeys downgrades Seq2 to Seq of keys
func Seq2ToSeqKeys[K, V any](i iter.Seq2[K, V]) iter.Seq[K] {
	return func(yield func(K) bool) {
		for k, _ := range i {
			if !yield(k) {
				return
			}
		}
	}
}

// Seq2ToSeqValues downgrades Seq2 to Seq of values
func Seq2ToSeqValues[K, V any](i iter.Seq2[K, V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for _, v := range i {
			if !yield(v) {
				return
			}
		}
	}
}
