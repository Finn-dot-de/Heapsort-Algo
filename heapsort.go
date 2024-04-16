package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	// gibt an wie oft eine zufällige Zahl gewählt werden soll
	numNumbers := 10_000
	var i int

	numbers := make([]int, numNumbers)
	for i := 0; i < numNumbers; i++ {
		numbers[i] = rand.Intn(numNumbers)
	}
	fmt.Printf("Unsortierte liste: %v <<<<<<<<\n", numbers)

	// Aufruf der Sortier Funktion
	numbers = heap_sort(numbers)

	// Überprüft, ob die Liste sortiert ist
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			log.Fatalf("Fehler bei der Sortierung !!")
		}
	}
	// Gibt die sortierte Liste aus
	fmt.Printf("Sortierte liste: %v <<<<<<< \n", numbers)

	fmt.Println("Bitte eine Entertaste drücken zum beenden..... ")
	fmt.Scanln(&i)
}

func heapify(heap []int, heap_size int, root_index int) []int {
	// Initialisiert den größten Index als Wurzelindex
	largest := root_index
	// Berechnet den Index des linken und rechten Kindes
	lef_child := 2*root_index + 1
	right_child := 2*root_index + 2

	// Überprüft, ob das linke Kind größer ist als die Wurzel
	if lef_child < heap_size && heap[largest] < heap[lef_child] {
		largest = lef_child
	}
	// Überprüft, ob das rechte Kind größer ist als die Wurzel
	if right_child < heap_size && heap[largest] < heap[right_child] {
		largest = right_child
	}
	// Wenn die Wurzel nicht das größte Element ist, tauscht sie mit dem größten Element
	if largest != root_index {
		heap[root_index], heap[largest] = heap[largest], heap[root_index]
		// Heapify den neuen größten Knoten
		heap = heapify(heap, heap_size, largest)
	}
	return heap
}

// Prüft die Länge der unsortierten Liste: Wenn die Liste nur ein Element oder kein Element enthält, wird sie als sortiert betrachtet und sofort zurückgegeben.
// Erstellt einen Heap aus der unsortierten Liste: Ein Heap ist eine spezielle Baumstruktur, in der jedes Elternelement größer oder gleich seinen Kindern ist. 
// Dies wird erreicht, indem die heapify Funktion auf jedes Element der Liste angewendet wird, beginnend in der Mitte der Liste und rückwärts gehend.
// Sortiert den Heap: Die heap_sort Funktion entfernt dann das größte Element (die Wurzel des Heaps) 
// und tauscht es mit dem letzten Element des Heaps. Dieses Element wird nun als sortiert betrachtet und aus dem Heap entfernt.
// Der Heap wird dann erneut “geheapified”, und der Prozess wird wiederholt, bis der gesamte Heap sortiert ist.
func heap_sort(unsorted_list []int) []int {
	// Überprüft, ob die Liste leer oder nur ein Element enthält
	if len(unsorted_list) <= 1 {
		return unsorted_list
	}
	// Berechnet die Länge der Liste
	heap_size := len(unsorted_list)
	// Baut den Max-Heap auf
	for index := heap_size/2 - 1; index >= 0; index-- {
		unsorted_list = heapify(unsorted_list, heap_size, index)
	}
	// Tauscht das erste Element mit dem letzten Element und heapify den reduzierten Heap
	for i := heap_size - 1; i > 0; i-- {
		unsorted_list[0], unsorted_list[i] = unsorted_list[i], unsorted_list[0]
		unsorted_list = heapify(unsorted_list, i, 0)
	}
	return unsorted_list
}