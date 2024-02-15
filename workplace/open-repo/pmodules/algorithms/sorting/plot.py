import matplotlib.pyplot as plt
import timeit

def selection_sort(A):
    n = len(A)
    for i in range(n - 1):
        min_index = i
        for j in range(i + 1, n):
            if A[j] < A[min_index]:
                min_index = j

        # Swap the found minimum element with the first element
        A[i], A[min_index] = A[min_index], A[i]

# Function to measure the running time of selection_sort
def measure_running_time(arr):
    start_time = timeit.default_timer()
    selection_sort(arr)
    end_time = timeit.default_timer()
    return end_time - start_time


# Vary the input size and measure the running time
input_sizes = list(range(1, 101))
running_times = []

for size in input_sizes:
    arr = list(range(size, 0, -1))  # Reverse sorted array for worst-case time complexity
    running_time = measure_running_time(arr)
    running_times.append(running_time)

# Plotting the running time
plt.plot(input_sizes, running_times, marker='o')
plt.title('Running Time of Selection Sort')
plt.xlabel('Input Size')
plt.ylabel('Running Time (seconds)')
plt.grid(True)
plt.show()
