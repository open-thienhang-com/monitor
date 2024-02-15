class Node:
    def __init__(self, data=None):
        self.data = data
        self.next = None

class Stack:
    def __init__(self):
        self.head = None

    def push(self, data):
        new_node = Node(data)
        new_node.next = self.head
        self.head = new_node

    def pop(self):
        if self.head is None:
            return None
        popped_data = self.head.data
        self.head = self.head.next
        return popped_data
    
    def print_stack(self):
        current = self.head
        print("Stack [", end="")
        while current:
            print(f" {current.data} ", end="")
            current = current.next
        print("]")
        print("-------------------------------")
# Testing the Stack
stack = Stack()
# Printing all stack elements
stack.print_stack()

# Pushing elements onto the stack
stack.push(10)
stack.push(30)
# Printing all stack elements
stack.print_stack()

# Popping an element from the stack
popped_value = stack.pop()
print("Popped:", popped_value)
# Printing all stack elements
stack.print_stack()

# Pushing another element
stack.push(80)
# Printing all stack elements
stack.print_stack()

# Popping again
popped_value = stack.pop()
print("Popped:", popped_value)
# Printing all stack elements
stack.print_stack()
