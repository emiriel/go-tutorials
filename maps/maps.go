package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    fields := strings.Fields(s)
    myMap := make(map[string]int)
    for i := range fields {
        myMap[fields[i]]++
    }
    return myMap
}

func main() {
    wc.Test(WordCount)
}
