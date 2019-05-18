package main

import "fmt"

func traverseMap(m map[string]string) {
    fmt.Println("Travesing map...")
    for k := range m {
        fmt.Println(k)
    }
    for k, v := range m {
        fmt.Println(k, v)
    }
    for _, v := range m {
        fmt.Println(v)
    }
}

func main() {
    m := map[string]string{
        "key1": "val1",
        "key2": "val2",
        "key3": "val3",
        "key4": "val4",
    }
    fmt.Println(m)

    m2 := make(map[string]int)  // m3 == empty map
    fmt.Println(m2)
    
    var m3 map[string]int  // m3 == nil
    fmt.Println(m3)

    traverseMap(m)

    val1 := m["key1"]
    fmt.Println(val1)
    val3 := m["key3"]
    fmt.Println(val3)

    val5 := m["key5"]
    fmt.Println(val5)

    if val6, ok := m["key6"]; ok {
        fmt.Println(val6)
    } else {
        fmt.Println("Key val6 does not exists.")
    }

    fmt.Println("Deleting element...")
    val1, ok := m["key1"]
    fmt.Println(val1, ok)
    delete(m, "key1")
    val1, ok = m["key1"]
    fmt.Println(val1, ok)
}
