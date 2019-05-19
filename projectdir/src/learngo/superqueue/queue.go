package superqueue

type Queue []interface{}

func (q *Queue) Push(n interface{}) {
    *q = append(*q, n)
}

func (q *Queue) Pop() (interface{}, bool) {
    var front interface{}
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
