package xiter

import (
	"iter"
)

// Map transforms values in Seq
func Map[FROM, TO any](seq iter.Seq[FROM], mapper func(FROM) TO) iter.Seq[TO] {
	return func(yield func(TO) bool) {
		for v := range seq {
			if !yield(mapper(v)) {
				return
			}
		}
	}
}

// Map2 transforms values in Seq2 retaining original keys
func Map2[K, VFROM, VTO any](seq iter.Seq2[K, VFROM], mapper func(K, VFROM) VTO) iter.Seq2[K, VTO] {
	return func(yield func(K, VTO) bool) {
		for k, v := range seq {
			if !yield(k, mapper(k, v)) {
				return
			}
		}
	}
}

// Map2K transforms both keys and values in Seq2
func Map2K[KFROM, VFROM, KTO, VTO any](seq iter.Seq2[KFROM, VFROM], mapper func(KFROM, VFROM) (KTO, VTO)) iter.Seq2[KTO, VTO] {
	return func(yield func(KTO, VTO) bool) {
		for k, v := range seq {
			if !yield(mapper(k, v)) {
				return
			}
		}
	}
}
