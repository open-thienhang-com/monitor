
class Node:
    def __init__(self, data=None):
        self.data = data
        self.next = None

class Queue:
    def __init__(self):
        self.front = None
        self.rear = None

    def enqueue(self, data):
        new_node = Node(data)
        if not self.rear:
            self.front = self.rear = new_node
            return
        self.rear.next = new_node
        self.rear = new_node

    def dequeue(self):
        if not self.front:
            return None
        popped_data = self.front.data
        self.front = self.front.next
        if not self.front:
            self.rear = None
        return popped_data
    
    def print_queue(self):
        current = self.front
        print("Queue [", end="")
        while current:
            print(f" {current.data} ", end="")
            current = current.next
        print("]")

# Testing the Queue
queue = Queue()
# Printing the queue
queue.print_queue()

# Enqueueing elements into the queue
queue.enqueue(4)
queue.enqueue(1)
queue.enqueue(3)
# Printing the queue
queue.print_queue()
# Dequeueing an element
dequeued_value = queue.dequeue()
print("Dequeued:", dequeued_value)
# Printing the queue
queue.print_queue()
# Enqueueing another element
queue.enqueue(8)
# Printing the queue
queue.print_queue()
# Dequeueing again
dequeued_value = queue.dequeue()
print("Dequeued:", dequeued_value)
# Printing the queue
queue.print_queue()