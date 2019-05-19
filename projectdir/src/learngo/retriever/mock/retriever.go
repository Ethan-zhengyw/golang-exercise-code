package mock

// define a mock class type
type Retriever struct {
    Contents string
}

// implement the Get interface by define a 
// new function whose receiver is the mock class retriever
func (r Retriever) Get(url string) string {
    return r.Contents
}

