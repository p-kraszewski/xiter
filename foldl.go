package xiter

import (
	"iter"
)

// Fold folds Seq aggregating values from the left.
func Fold[V, ACC any](seq iter.Seq[V], acc ACC, folder func(ACC, V) ACC) ACC {
	a := acc
	for v := range seq {
		a = folder(a, v)
	}
	return a
}

// Fold2 folds Seq2 aggregating values from the left. Aggregating function sees both keys and values
func Fold2[K, V, ACC any](seq iter.Seq2[K, V], acc ACC, folder func(ACC, K, V) ACC) ACC {
	a := acc
	for k, v := range seq {
		a = folder(a, k, v)
	}
	return a
}
