class LinearProbingHashTable:
    def __init__(self, size):
        self.size = size
        self.table = [None] * size

    def hash_function(self, key):
        return key % self.size

    def insert(self, key):
        index = self.hash_function(key)
        print(f"\nCalculate h({key}) = {index}")
        while self.table[index] is not None:
            print(f"Index {index} is already occupied.")
            index = (index + 1) % self.size
            print(key, " -> Linearly probe to the next index:", index)

        self.table[index] = key

    def display(self):
        print("Hash Table:")
        for i, key in enumerate(self.table):
            print(f"Index {i}: {key}")

# Example usage:
hash_table = LinearProbingHashTable(11)

keys = [10, 22, 31, 4, 15, 28, 17, 88, 59]

for key in keys:
    hash_table.insert(key)

hash_table.display()

