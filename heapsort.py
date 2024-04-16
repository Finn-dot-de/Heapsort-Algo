import random

# Erzeugt eine Liste von 10 zufälligen Zahlen zwischen 0 und 100
numbers = [random.randint(0, 100) for _ in range(10)]

def heapify(heap, heap_size, root_index):
    # Initialisiert den größten Index als Wurzelindex
    largest = root_index
    # Berechnet den Index des linken und rechten Kindes
    left_child = 2 * root_index + 1
    right_child = 2 * root_index + 2

    # Überprüft, ob das linke Kind größer ist als die Wurzel
    if left_child < heap_size and heap[largest] < heap[left_child]:
        largest = left_child
    # Überprüft, ob das rechte Kind größer ist als die Wurzel
    if right_child < heap_size and heap[largest] < heap[right_child]:
        largest = right_child
    # Wenn die Wurzel nicht das größte Element ist, tauscht sie mit dem größten Element
    if largest != root_index:
        heap[root_index], heap[largest] = heap[largest], heap[root_index]
        # Heapify den neuen größten Knoten
        heap = heapify(heap, heap_size, largest)
    return heap

def heap_sort(unsorted_list):
    # Überprüft, ob die Liste leer oder nur ein Element enthält
    if len(unsorted_list) <= 1:
        return unsorted_list
    # Berechnet die Länge der Liste
    heap_size = len(unsorted_list)
    # Baut den Max-Heap auf
    for index in reversed(range(heap_size//2)):
        unsorted_list = heapify(unsorted_list, heap_size, index)
    # Tauscht das erste Element mit dem letzten Element und heapify den reduzierten Heap
    for i in reversed(range(1, heap_size)):
        unsorted_list[i], unsorted_list[0] = unsorted_list[0], unsorted_list[i]
        unsorted_list = heapify(unsorted_list, i, 0)
    return unsorted_list

# Sortiert die Liste mit der Heap-Sort-Methode
numbers = heap_sort(numbers)
# Überprüft, ob die Liste sortiert ist
for i in range(1, len(numbers)):
    assert numbers[i - 1] <= numbers[i]
# Gibt die sortierte Liste aus
print(numbers)
