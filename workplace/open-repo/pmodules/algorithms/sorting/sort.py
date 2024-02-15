import random
import timeit
import matplotlib.pyplot as plt

# Bubble Sort
def bubble_sort(arr):
    n = len(arr)
    for i in range(n):
        for j in range(0, n - i - 1):
            if arr[j] > arr[j + 1]:
                arr[j], arr[j + 1] = arr[j + 1], arr[j]
    return arr

# Selection Sort
def selection_sort(arr):
    n = len(arr)
    for i in range(n):
        min_index = i
        for j in range(i + 1, n):
            if arr[j] < arr[min_index]:
                min_index = j
        arr[i], arr[min_index] = arr[min_index], arr[i]
    return arr

# Insertion Sort
def insertion_sort(arr):
    n = len(arr)
    for i in range(1, n):
        key = arr[i]
        j = i - 1
        while j >= 0 and key < arr[j]:
            arr[j + 1] = arr[j]
            j -= 1
        arr[j + 1] = key
    return arr

# Merge Sort
def merge_sort(arr):
    if len(arr) > 1:
        mid = len(arr) // 2
        left_half = arr[:mid]
        right_half = arr[mid:]

        merge_sort(left_half)
        merge_sort(right_half)

        i = j = k = 0

        while i < len(left_half) and j < len(right_half):
            if left_half[i] < right_half[j]:
                arr[k] = left_half[i]
                i += 1
            else:
                arr[k] = right_half[j]
                j += 1
            k += 1

        while i < len(left_half):
            arr[k] = left_half[i]
            i += 1
            k += 1

        while j < len(right_half):
            arr[k] = right_half[j]
            j += 1
            k += 1

# Quicksort 
def quick_sort(arr):
    if len(arr) <= 1:
        return arr
    else:
        pivot = arr.pop()
        less_than_pivot = []
        greater_than_pivot = []

        for element in arr:
            if element <= pivot:
                less_than_pivot.append(element)
            else:
                greater_than_pivot.append(element)

        return quick_sort(less_than_pivot) + [pivot] + quick_sort(greater_than_pivot)

# Counting Sort
def counting_sort(arr):
    max_val = max(arr)
    min_val = min(arr)
    count = [0] * (max_val - min_val + 1)

    for num in arr:
        count[num - min_val] += 1

    sorted_arr = []
    for i in range(len(count)):
        sorted_arr += [i + min_val] * count[i]

    return sorted_arr

# Radix Sort
def radix_sort(arr):
    max_val = max(arr)
    exp = 1

    while max_val // exp > 0:
        counting_sort_radix(arr, exp)
        exp *= 10

def counting_sort_radix(arr, exp):
    n = len(arr)
    output = [0] * n
    count = [0] * 10

    for i in range(n):
        index = arr[i] // exp
        count[index % 10] += 1

    for i in range(1, 10):
        count[i] += count[i - 1]

    i = n - 1
    while i >= 0:
        index = arr[i] // exp
        output[count[index % 10] - 1] = arr[i]
        count[index % 10] -= 1
        i -= 1

    for i in range(n):
        arr[i] = output[i]

# Heap Sort
def heap_sort(arr):
    n = len(arr)

    for i in range(n // 2 - 1, -1, -1):
        heapify(arr, n, i)

    for i in range(n - 1, 0, -1):
        arr[i], arr[0] = arr[0], arr[i]
        heapify(arr, i, 0)

def heapify(arr, n, i):
    largest = i
    left_child = 2 * i + 1
    right_child = 2 * i + 2

    if left_child < n and arr[left_child] > arr[largest]:
        largest = left_child

    if right_child < n and arr[right_child] > arr[largest]:
        largest = right_child

    if largest != i:
        arr[i], arr[largest] = arr[largest], arr[i]
        heapify(arr, n, largest)

# Bucket Sort
def bucket_sort(arr):
    buckets = []
    for i in range(10):  # 10 buckets for simplicity, adjust as needed
        buckets.append([])

    for num in arr:
        index = num // 10  # adjust the bucket index as needed
        buckets[index].append(num)

    sorted_arr = []
    for bucket in buckets:
        sorted_arr.extend(sorted(bucket))

    return sorted_arr


# Timsort

def main():
    # Generate random data
    data_size = 1000
    data = [random.randint(1, 10000) for _ in range(data_size)]

    # Bubble Sort
    bubble_sort_time = timeit.timeit(lambda: bubble_sort(data.copy()), number=1)

    # Selection Sort
    selection_sort_time = timeit.timeit(lambda: selection_sort(data.copy()), number=1)

    # Insertion Sort
    insertion_sort_time = timeit.timeit(lambda: insertion_sort(data.copy()), number=1)

    # Merge Sort
    merge_sort_time = timeit.timeit(lambda: merge_sort(data.copy()), number=1)

    # Quicksort
    quick_sort_time = timeit.timeit(lambda: quick_sort(data.copy()), number=1)

    # Counting Sort
    counting_sort_time = timeit.timeit(lambda: counting_sort(data.copy()), number=1)

    # Radix Sort
    radix_sort_time = timeit.timeit(lambda: radix_sort(data.copy()), number=1)

    # Heap Sort
    heap_sort_time = timeit.timeit(lambda: heap_sort(data.copy()), number=1)

    # Bucket Sort
    bucket_sort_time = timeit.timeit(lambda: bucket_sort(data.copy()), number=1)

    # Plotting
    algorithms = ['Bubble', 'Selection', 'Insertion', 'Merge', 'Quicksort',
                  'Counting', 'Radix', 'Heap', 'Bucket']
    times = [bubble_sort_time, selection_sort_time, insertion_sort_time,
             merge_sort_time, quick_sort_time, counting_sort_time,
             radix_sort_time, heap_sort_time, bucket_sort_time]

    plt.barh(algorithms, times, color='skyblue')
    plt.xlabel('Time (seconds)')
    plt.title('Time Consumption of Sorting Algorithms')
    plt.show()

if __name__ == "__main__":
    main()