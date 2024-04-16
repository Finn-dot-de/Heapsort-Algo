package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	// gibt an wie oft eine zufällige Zahl gewählt werden soll
	numNumbers := 10_0000000

	numbers := make([]int, numNumbers)
	for i := 0; i < numNumbers; i++ {
		numbers[i] = rand.Intn(numNumbers)
	}
	fmt.Printf("Unsortierte liste: %v <<<<<<<<\n", numbers)

	// Aufruf der Sortier Funktion
	numbers = heap_sort(numbers)

	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			log.Fatalf("Fehler konnte nicht sortiert werden!!")
		}
	}
	fmt.Printf("Sortierte liste: %v", numbers)
}

func heapify(heap []int, heap_size int, root_index int) []int {
	largest := root_index
	lef_child := 2*root_index + 1
	right_child := 2*root_index + 2

	if lef_child < heap_size && heap[largest] < heap[lef_child] {
		largest = lef_child
	}
	if right_child < heap_size && heap[largest] < heap[right_child] {
		largest = right_child
	}
	if largest != root_index {
		heap[root_index], heap[largest] = heap[largest], heap[root_index]

		heap = heapify(heap, heap_size, largest)
	}
	return heap
}

func heap_sort(unsorted_list []int) []int {
	if len(unsorted_list) <= 1 {
		return unsorted_list
	}
	heap_size := len(unsorted_list)

	for index := heap_size/2 - 1; index >= 0; index-- {
		unsorted_list = heapify(unsorted_list, heap_size, index)
	}

	for i := heap_size - 1; i > 0; i-- {
		unsorted_list[0], unsorted_list[i] = unsorted_list[i], unsorted_list[0]
		unsorted_list = heapify(unsorted_list, i, 0)
	}
	return unsorted_list
}
