package main

import "fmt"

func checkError(err error){
    if err != nil {
        fmt.Println(err)
    }
}

func main() {
    
    var x1 int    
    _, err := fmt.Scanf("%d", &x1)
    checkError(err)
    
    var v1 int
    _, err = fmt.Scanf("%d", &v1)
    checkError(err)
    
    var x2 int
    _, err = fmt.Scanf("%d", &x2)
    checkError(err)

    var v2 int
    _, err = fmt.Scanf("%d", &v2)
    checkError(err)
    
    if x1 == x2 && v1 == v2 {
        fmt.Println("YES")
    } else if x1 == x2 && v1 != v2{
        fmt.Println("NO")
    } else {
        if v1 > v2 && ((x2-x1)%(v1-v2)) == 0 {
            fmt.Println("YES")
        } else {
            fmt.Println("NO")
        }
    }
}
