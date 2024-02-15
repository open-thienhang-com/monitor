def selection_sort(A):
    n = len(A)
    for i in range(n - 1):
        min_index = i
        for j in range(i + 1, n):
            if A[j] < A[min_index]:
                min_index = j

            A[i], A[min_index] = A[min_index], A[i]
            print("Step ", i,  A)

# Example usage:
A = [64, 35, 12, 22, 11]
selection_sort(A)
print("Sorted array:", A)