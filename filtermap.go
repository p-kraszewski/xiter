package xiter

import (
	"iter"
)

// FilterMap transforms values in Seq, optionally dropping ones that resolve filterMapper() to _,false
func FilterMap[FROM, TO any](seq iter.Seq[FROM], filterMapper func(FROM) (TO, bool)) iter.Seq[TO] {
	return func(yield func(TO) bool) {
		for v := range seq {
			nv, ok := filterMapper(v)
			if ok {
				if !yield(nv) {
					return
				}
			}
		}
	}
}

// FilterMap2 transforms values in Seq2 (retaining keys), optionally dropping ones that resolve filterMapper() to _,false.
func FilterMap2[K, VFROM, VTO any](seq iter.Seq2[K, VFROM], filterMapper func(K, VFROM) (VTO, bool)) iter.Seq2[K, VTO] {
	return func(yield func(K, VTO) bool) {
		for k, v := range seq {
			nv, ok := filterMapper(k, v)
			if ok {
				if !yield(k, nv) {
					return
				}
			}
		}
	}
}

// FilterMap2K transforms both keys and values in Seq2, optionally dropping ones that resolve filterMapper() to _,_,false.
func FilterMap2K[KFROM, VFROM, KTO, VTO any](seq iter.Seq2[KFROM, VFROM], filterMapper func(KFROM, VFROM) (KTO, VTO, bool)) iter.Seq2[KTO, VTO] {
	return func(yield func(KTO, VTO) bool) {
		for k, v := range seq {
			nk, nv, ok := filterMapper(k, v)
			if ok {
				if !yield(nk, nv) {
					return
				}
			}
		}
	}
}
