class Node:
    def __init__(self, data=None):
        self.data = data
        self.prev = None
        self.next = None

class DoublyLinkedList:
    def __init__(self):
        self.head = None

    def delete_at_start(self):
        if not self.head:
            print("List is empty. Cannot delete.")
            return
        self.head = self.head.next
        if self.head:
            self.head.prev = None

    def delete_at_end(self):
        if not self.head:
            print("List is empty. Cannot delete.")
            return
        if not self.head.next:
            self.head = None
            return
        current = self.head
        while current.next.next:
            current = current.next
        current.next = None

    def insert_at_end(self, data):
        new_node = Node(data)
        if not self.head:
            self.head = new_node
        else:
            current = self.head
            while current.next:
                current = current.next
            current.next = new_node
            new_node.prev = current

    def insert_at_start(self, data):
            new_node = Node(data)
            new_node.next = self.head
            if self.head:
                self.head.prev = new_node
            self.head = new_node

    def display(self):
        current = self.head
        while current:
            print(current.data, end=" <-> ")
            current = current.next
        print("None")
    

# Test the doubly linked list
dll = DoublyLinkedList()

# Test 1: Append elements
dll.insert_at_end(1)
dll.insert_at_end(2)
dll.insert_at_end(3)
dll.display()

# Expected output: 1 <-> 2 <-> 3 <-> None

# Test 2: Append more elements
dll.insert_at_end(4)
dll.insert_at_end(5)
dll.display()

# Expected output: 1 <-> 2 <-> 3 <-> 4 <-> 5 <-> None

# Test 3: Insert element at the end
dll.insert_at_end(6)
dll.display()

# Expected output: 1 <-> 2 <-> 3 <-> 4 <-> 5 <-> 6 <-> None

# Test 4: Delete at start
dll.delete_at_start()
dll.display()

# Expected output: 2 <-> 3 <-> 4 <-> 5 <-> 6 <-> None

# Test 5: Delete at end
dll.delete_at_end()
dll.display()

# Expected output: 2 <-> 3 <-> 4 <-> 5 <-> None

# Test 6: Insert at start
dll.insert_at_start(11)
dll.display()

# Expected output: 111 <-> 2 <-> 3 <-> 4 <-> 5 <-> None
