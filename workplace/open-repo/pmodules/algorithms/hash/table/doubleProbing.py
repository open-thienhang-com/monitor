class QuadraticProbingHashTable:
    def __init__(self, size, c1, c2):
        self.size = size
        self.table = [None] * size
        self.c1 = c1
        self.c2 = c2

    def hash_function(self, key):
        print("=> ", key, key % self.size)
        return key % self.size

    def insert(self, key):
        index = self.hash_function(key)
        i = 1

        while self.table[index] is not None:
            index = (index + self.c1 * i + self.c2 * i**2) % self.size
            print("---> ", key, i, index, "| ", self.c1*i, self.c2, i**2)
            i += 1
           

        self.table[index] = key

    def display(self):
        print("Hash Table:")
        for i, key in enumerate(self.table):
            print(f"Index {i}: {key}")


class DoubleHashingHashTable:
    def __init__(self, size):
        self.size = size
        self.table = [None] * size

    def hash_function1(self, key):
        print("=> ", key, key % self.size)
        return key % self.size

    def hash_function2(self, key):
        return 1 + (key % (self.size - 1))

    def insert(self, key):
        index = self.hash_function1(key)
        step = self.hash_function2(key)

        while self.table[index] is not None:
            index = (index + step) % self.size

        self.table[index] = key

    def display(self):
        print("Hash Table:")
        for i, key in enumerate(self.table):
            print(f"Index {i}: {key}")

keys = [10, 22, 31, 4, 15, 28, 17, 88, 59]

# Example usage for quadratic probing
quadratic_hash_table = QuadraticProbingHashTable(11, 1, 3)

for key in keys:
    quadratic_hash_table.insert(key)

quadratic_hash_table.display()

# # Example usage for double hashing
# double_hash_table = DoubleHashingHashTable(11)

# for key in keys:
#     double_hash_table.insert(key)

# double_hash_table.display()
