package queues

import "fmt"

func ExampleNew() {
	q := New[int]()

	fmt.Println(q)

	q.Enqueue(1)
	q.Enqueue(2)

	fmt.Println(q)

	// Output:
	// Queue[]
	// Queue[1 2]
}

func ExampleQueue_IsEmpty() {
	q := New[int]()

	fmt.Println(q.IsEmpty())
	q.Enqueue(1)
	fmt.Println(q.IsEmpty())

	// Output:
	// true
	// false
}

func ExampleQueue_Size() {
	q := New[int]()

	fmt.Println(q.Size())
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	fmt.Println(q.Size())

	// Output:
	// 0
	// 4
}

func ExampleQueue_Clear() {
	q := New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)
	q.Clear()

	fmt.Println(q)

	// Output:
	// Queue[]
}

func ExampleQueue_Enqueue() {
	q := New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	fmt.Println(q)

	// Output:
	// Queue[1 2 3 4]
}

func ExampleQueue_Dequeue() {
	q := New[int]()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	q.Dequeue()
	q.Dequeue()

	fmt.Println(q)

	// Output:
	// Queue[3 4]
}
