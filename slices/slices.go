package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
    mySlice := make([][]uint8, dy)
    for i := 0; i < dy; i++ {
        mySlice[i] = make ([] uint8, dx)
        for j := 0 ; j < dx ; j++ {
            mySlice[i][j] = uint8(i + j)
        }
    }
    return mySlice
}

func main() {
    pic.Show(Pic)
}
