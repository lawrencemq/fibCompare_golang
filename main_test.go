package main

import (
	"testing"
	"./fibs"
)


const smallBenchmarkTest = 5
const mediumBenchmarkTest = 10
const largeBenchmarkTest = 25


func TestLinkedListFib(t *testing.T){
	if 1 != fibs.FibLinkedList(0){
		t.Fail()
	}
	if 1 != fibs.FibLinkedList(1){
		t.Fail()
	}
	if 2 != fibs.FibLinkedList(2){
		t.Fail()
	}
	if 89 != fibs.FibLinkedList(10){
		t.Fail()
	}

}

func BenchmarkLinkedListFib_SmallNumber(b *testing.B){
	for i := 0; i < b.N; i++ {
		fibs.FibLinkedList(smallBenchmarkTest)
	}
}
func BenchmarkLinkedListFib_MediumNumber(b *testing.B){
	for i := 0; i < b.N; i++ {
		fibs.FibLinkedList(mediumBenchmarkTest)
	}
}
func BenchmarkLinkedListFib_LargeNumber(b *testing.B){
	for i := 0; i < b.N; i++ {
		fibs.FibLinkedList(largeBenchmarkTest)
	}
}
func TestTwoNumberFib(t *testing.T){
	if 1 != fibs.FibTwoNumbers(0){
		t.Fail()
	}
	if 1 != fibs.FibTwoNumbers(1){
		t.Fail()
	}
	if 2 != fibs.FibTwoNumbers(2){
		t.Fail()
	}
	if 89 != fibs.FibTwoNumbers(10){
		t.Fail()
	}

}

func BenchmarkTwoNumberFib_SmallNumber(b *testing.B){
	for i := 0; i < b.N; i++ {
		fibs.FibTwoNumbers(smallBenchmarkTest)
	}
}
func BenchmarkTwoNumberFib_MediumNumber(b *testing.B){
	for i := 0; i < b.N; i++ {
		fibs.FibTwoNumbers(mediumBenchmarkTest)
	}
}
func BenchmarkTwoNumberFib_LargeNumber(b *testing.B){
	for i := 0; i < b.N; i++ {
		fibs.FibTwoNumbers(largeBenchmarkTest)
	}
}

func TestRecursiveFib(t *testing.T){
	if 1 != fibs.FibRecursive(0){
		t.Fail()
	}
	if 1 != fibs.FibRecursive(1){
		t.Fail()
	}
	if 2 != fibs.FibRecursive(2){
		t.Fail()
	}
	if 89 != fibs.FibRecursive(10){
		t.Fail()
	}

}

func BenchmarkRecursiveFib_SmallNumber(b *testing.B){
	for i := 0; i < b.N; i++ {
		fibs.FibRecursive(smallBenchmarkTest)
	}
}
func BenchmarkRecursiveFib_MediumNumber(b *testing.B){
	for i := 0; i < b.N; i++ {
		fibs.FibRecursive(mediumBenchmarkTest)
	}
}
func BenchmarkRecursiveFib_LargeNumber(b *testing.B){
	for i := 0; i < b.N; i++ {
		fibs.FibRecursive(largeBenchmarkTest)
	}
}
