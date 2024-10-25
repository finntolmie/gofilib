package gofilib

import "golang.org/x/exp/constraints"

func Map[I any, O any](arr []I, mapFunc func(I) O) []O {
	res := make([]O, 0, len(arr))
	for _, el := range arr {
		res = append(res, mapFunc(el))
	}
	return res
}

func Filter[I any](arr []I, filterFunc func(I) bool) []I {
	res := make([]I, 0)
	for _, el := range arr {
		if filterFunc(el) {
			res = append(res, el)
		}
	}
	return res
}

func Reduce[I any, O constraints.Ordered](arr []I, reduceFunc func(I) O, acc O) O {
	var res O
	for _, el := range arr {
		res += reduceFunc(el)
	}
	return res
}
