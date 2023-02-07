package result

import (
	"fmt"
	"leetcode/utils"
)

const (
	Accepted    = "Accepted"
	WrongAnswer = "Wrong Answer"
)

type LogFunc func(...any)

type Result interface {
	String() string
	Accepted() bool
}

type ValueResult[T comparable] struct {
	expected T
	actual   T
}

func (e ValueResult[T]) Accepted() bool {
	return e.expected == e.actual
}

func (e ValueResult[T]) String() string {
	if e.Accepted() {
		return fmt.Sprintf("Accepted.\nanswer: %v", e.actual)
	}
	return fmt.Sprintf("Wrong Answer.\nexpected: %v\nactual:   %v", e.expected, e.actual)
}

type SliceResult[T comparable] struct {
	expected []T
	actual   []T
}

func (e SliceResult[T]) Accepted() bool {
	return utils.SliceEqual(e.expected, e.actual)
}

func (e SliceResult[T]) String() string {
	if e.Accepted() {
		return fmt.Sprintf("Accepted.\nanswer: %v", e.actual)
	}
	return fmt.Sprintf("Wrong Answer.\nexpected: %v\nactual:   %v", e.expected, e.actual)
}

func New[T comparable](actual, expected T) Result {
	return ValueResult[T]{expected, actual}
}

func NewSlice[T comparable](actual, expected []T) Result {
	return SliceResult[T]{expected, actual}
}
