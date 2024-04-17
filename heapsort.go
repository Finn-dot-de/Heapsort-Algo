package main

import (
	"fmt"
	"log"
	"math/rand"
)

func main() {
	// gibt an wie oft eine zufällige Zahl plus den Bereich gewählt werden soll
	numNumbers := 100_000
	var i int

	// Erstellt eine Zufällige Zahl
	numbers := make([]int, numNumbers)
	for i := 0; i < numNumbers; i++ {
		numbers[i] = rand.Intn(numNumbers)
	}
	fmt.Printf("Unsortierte liste: %v <<<<<<<<\n\n", numbers)

	// Aufruf der Sortier Funktion
	numbers = Heap_sort(numbers)

	// Überprüft, ob die Liste sortiert ist
	for i := 1; i < len(numbers); i++ {
		if numbers[i-1] > numbers[i] {
			log.Fatalf("Fehler bei der Sortierung !!")
		}
	}
	// Gibt die sortierte Liste aus
	fmt.Printf("Sortierte liste: %v <<<<<<< \n", numbers)

	fmt.Println("\nBitte eine Entertaste drücken zum beenden..... ")
	fmt.Scanln(&i)
}

// Initialisierung: Sie initialisiert den größten Index als Wurzelindex und berechnet den Index des linken und rechten Kindes.
//
// Überprüfung der Kinder: Sie überprüft, ob das linke oder rechte Kind größer ist als die Wurzel. Wenn ja, wird der Index des größten Kindes als der neue größte Index festgelegt.
//
// Tauschen der Elemente: Wenn die Wurzel nicht das größte Element ist (d.h., eines der Kinder ist größer), tauscht sie die Wurzel mit dem größten Element.
//
// Rekursiver Aufruf: Nach dem Tauschen könnte der getauschte Knoten die Heap-Eigenschaft verletzen.
// Daher ruft die Funktion sich selbst rekursiv auf, um sicherzustellen, dass der getauschte Knoten und seine Kinder die Heap-Eigenschaft erfüllen.
//
// Rückgabe des Heaps: Schließlich gibt die Funktion den modifizierten Heap zurück.
func Heapify(heap []int, heap_size int, root_index int) []int {
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
		heap = Heapify(heap, heap_size, largest)
	}
	return heap
}

// Prüft die Länge der unsortierten Liste: Wenn die Liste nur ein Element oder kein Element enthält, wird sie als sortiert betrachtet und sofort zurückgegeben.
//
// Erstellt einen Heap aus der unsortierten Liste: Ein Heap ist eine spezielle Baumstruktur, in der jedes Elternelement größer oder gleich seinen Kindern ist.
// Dies wird erreicht, indem die Heapify Funktion auf jedes Element der Liste angewendet wird, beginnend in der Mitte der Liste und rückwärts gehend.
//
// Sortiert den Heap: Die Heap_sort Funktion entfernt dann das größte Element (die Wurzel des Heaps)
// und tauscht es mit dem letzten Element des Heaps.
// Dieses Element wird nun als sortiert betrachtet und aus dem Heap entfernt.
//
// Der Heap wird dann erneut “geheapified”, und der Prozess wird wiederholt, bis der gesamte Heap sortiert ist.
func Heap_sort(unsorted_list []int) []int {
	// Überprüft, ob die Liste leer oder nur ein Element enthält
	if len(unsorted_list) <= 1 {
		return unsorted_list
	}
	// Berechnet die Länge der Liste
	heap_size := len(unsorted_list)
	// Baut den Max-Heap auf
	for index := heap_size/2 - 1; index >= 0; index-- {
		unsorted_list = Heapify(unsorted_list, heap_size, index)
	}
	// Tauscht das erste Element mit dem letzten Element und Heapify den reduzierten Heap
	for i := heap_size - 1; i > 0; i-- {
		unsorted_list[0], unsorted_list[i] = unsorted_list[i], unsorted_list[0]
		unsorted_list = Heapify(unsorted_list, i, 0)
	}
	return unsorted_list
}
