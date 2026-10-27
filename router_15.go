package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.3" )

func nbsmes2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B4muSky3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wV4pz6fEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 251
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WejXa6RHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func meyrLA7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GiJrlKSfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1tlPyPxOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func En9HVtrIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lav0MeMmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zgJGYwhrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YUBjB94LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7f6fOyytWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eBEBN62tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xAbg3ZCvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5J4uqWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CNoIQRDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YLXylpbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EfjLMKPrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6I9814IFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WjMxmNzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MbLmfhULWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRl0fJ3RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OmET5FDbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B8paOSaGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N6Q9lVztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhCT9i30Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BjG4qvNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uO07EqBZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bLdTERs7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SYGAEv3UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ApoqWeFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mH0CWGD4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VHBJUlj0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1SEkLWSJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J1kNiLVcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kqhBSDBKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pdGFB334Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func seEgUCKsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wQChQkwOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SOxvQB6hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZOZsLYOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBsruHkSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ex20HJG2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CklzHgYtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oZObgQYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o8zorhJ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H4nqPpQkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqpIBXeTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func riZ2if7kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EiU0HlbQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cgzT0ncBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oAoSiQl3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e4Z9X5VrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cObZ2cZTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pMh6miqWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qIVUFxQOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ppnnvMnDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 22p37gWyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BXX99TgsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lrqksrC3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dzP2WoczWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q3oOjupqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WUM6LHmCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 37
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R6irisW4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ppJlvHgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8ONj03vfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bOMzB9CvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oVz7zYUtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R02fDfOFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KrmgG3jPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LnX3VqXtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gDZE5NkbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yhUNQ9DdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8C0vhYVeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ctyX0iZWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5S09lkc1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MwfN3t8aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3oMcYhHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ax8Wtc05Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VZ7vaOVJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3z1a54NXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hkljCSa1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQ953JPWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFUZnZSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yLN1vo7VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m7AI2CX6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RFY15RJFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1YB8RbBIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DTDwMIfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E3vfYBF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8Zdk1zuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KJe5lBnHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2TF1kI4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KhbYLs10Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vkHqjSe9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t0z1gEgdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kiM3CwkWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fckpCEJ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SiMZd9N2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R3lH9fMqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a6B2j0qWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G4LByCEGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tW40tMzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GRnSrcGBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8PEPtypOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fl5qE8oiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGoTr2h5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dnRO46aMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BqfBQM7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TYnzdETgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GYc4NOLyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5rw3tJZRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GSwNLqL1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QF9qnabDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8neHvZvCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IhX5cMfZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cqVSQDjhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bJsPs5MWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MmEbARO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YmPaIb3CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6tOn9h4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ODH1QgPSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dKnC5YI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K6v4vWrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yVlsHgPgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GzNfJ39xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCC1lYShWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bFbHsSh1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xhekKe4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tKt0vlJsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D5P1E4McWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pY3Vvq1nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xcgSl0o0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 18yvrgK1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yvOseLHbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1gyHjE1WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xxLpZtoiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEhIadqjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RPZHeVzRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EFjsoA9hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UYDFw19CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l5XmLH7SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pa2MlIdGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NKRWlKIBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VuqnagBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iyCDwBdkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JT3AuAX7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 72vJBUPVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXfCU3bIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z1rPbe87Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WpfivOR0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CepOFuQ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LW3nb0pPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zmQcr1fEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j5Azu3wAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fMtvPs5RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fKXq2i3HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M1lru5xnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQlhO7QIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oteD5hQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJ7ZIdJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mKobVOhyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NYSGYUVxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DKEa8peBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qvpy6JeBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KCrWMqCEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DsijGjgYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZybVAz2fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lhos5GHBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xJY1A3TMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGDAmoyJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EnI2StRTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1IPUbFvaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LRoAbEMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xkbe2mtFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func usYwWOa3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQfBjpmOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jQ8mFjRVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func urWpw2KrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aponA5jvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T9CpVrXCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uRMeGHf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4m7PqLlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GvJXwWKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uPnoVNRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t9WLzGPoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kw8Vfg0hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0LGoVkMWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NmzTzTorWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0NuefDYjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zGVjOemeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R39nMGmoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2w4SHAIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X8Cm6ivVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 18vCfs5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NThfFUgDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EnpjvltSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3Qyxkd9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CSijVRDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 54YkFsweWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KXDdYAo4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CU8v0zjOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H00MiJT1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CVXLKaE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hVvf3uXUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BdEZqJq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IbiAmG8NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rC3QeEEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w9ZuLve9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S64KLEOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YoYOPHWSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func prdl1c4PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LkfuVyApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCgL3Ss2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kmFmLmdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0EDGntv1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HEA0mH95Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fxyAarlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9qHIzfFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gvTEkv5jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dLD2gt2bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RxXxEugVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4GMNjtdrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DvzvxfETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cK3g1qmpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7AIoUkzKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func asNvmBK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Oi9uC1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nD9YubLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ySs9csUnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XUmw2lwzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fIC6ore7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3DDLCMMLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PgHxTDyoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wOGNIZapWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Bl0voLNuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RglRrB1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2VrQQs04Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 59Ris38EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g8jFtD4eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6te2e256Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eCCzovblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pscn9l7YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jqjJKPdDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5axufkGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mtdwMkk3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

