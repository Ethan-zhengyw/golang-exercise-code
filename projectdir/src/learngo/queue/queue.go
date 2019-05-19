package queue

type Queue []int

func (q *Queue) Push(n int) {
    *q = append(*q, n)
}

func (q *Queue) Pop() (int, bool) {
    var front int
    ok := true
    if len(*q) == 0 {
        ok = false
    } else {
        front = (*q)[0]
        *q = (*q)[1:]
    }
    return front, ok
}

func (q Queue) IsEmpty() bool {
    return len(q) == 0
}
