package util

import (
    "fmt"
)

func Check(err error) {
    if err != nil {
        fmt.Println(err)
    }
}
