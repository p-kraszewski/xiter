package xiter

import (
	"iter"
)

// Filter drops values from Seq that resolve filter() to false
func Filter[V any](seq iter.Seq[V], filter func(V) bool) iter.Seq[V] {
	return func(yield func(V) bool) {
		for v := range seq {
			if filter(v) {
				if !yield(v) {
					return
				}
			}
		}
	}
}

// Filter2 drops values from Seq2 that resolve filter() to false
func Filter2[K, V any](seq iter.Seq2[K, V], filter func(K, V) bool) iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		for k, v := range seq {
			if filter(k, v) {
				if !yield(k, v) {
					return
				}
			}
		}
	}
}
