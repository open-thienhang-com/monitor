import matplotlib.pyplot as plt
import matplotlib
matplotlib.__version__
import timeit
import random

def heapify(arr, n, i):
    largest = i  # Initialize largest as root
    left_child = 2 * i + 1
    right_child = 2 * i + 2

    # Check if left child of root exists and is greater than the root
    if left_child < n and arr[i] < arr[left_child]:
        largest = left_child

    # Check if right child of root exists and is greater than the largest so far
    if right_child < n and arr[largest] < arr[right_child]:
        largest = right_child

    # Change root, if needed
    if largest != i:
        arr[i], arr[largest] = arr[largest], arr[i]  # Swap
        heapify(arr, n, largest)

def heap_sort(arr):
    n = len(arr)

    # Build a max heap
    for i in range(n // 2 - 1, -1, -1):
       
        heapify(arr, n, i)
        
    # Extract elements one by one
    for i in range(n - 1, 0, -1):
        
        # print("\n --------PROCESS----------\n",)
        arr[i], arr[0] = arr[0], arr[i]  # Swap
        heapify(arr, i, 0)
       

def test_heap_sort():
    # Test 1
    arr1 = [5, 13, 2, 25, 7, 17, 20, 8]
    heap_sort(arr1)
    print("Sorted array (Test 1):", arr1)

    # Test 2
    arr2 = [64, 34, 25, 12, 22, 11, 90]
    heap_sort(arr2)
    print("Sorted array (Test 2):", arr2)

    # Test 3
    arr3 = [38, 27, 43, 3, 9, 82, 10]
    heap_sort(arr3)
    print("Sorted array (Test 3):", arr3)

def plot_heap_sort_time_complexity():
    input_sizes = [10, 100, 500, 1000, 5000]
    time_taken = []

    for size in input_sizes:
        input_array = list(range(size, 0, -1))  # Worst-case scenario

        # Measure the execution time
        execution_time = timeit.timeit(lambda: heap_sort(input_array.copy()), number=1)
        time_taken.append(execution_time)

    # Plotting
    #plt.plot(input_sizes, time_taken, marker='o')
    # plt.title('Heap Sort Time Complexity')
    # plt.xlabel('Input Size')
    # plt.ylabel('Time (seconds)')
    # plt.show()


if __name__ == "__main__":
    #test_heap_sort()
    # Call the plot function
    plot_heap_sort_time_complexity()
