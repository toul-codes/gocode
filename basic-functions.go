package main 

import ( 
        "fmt"
        "math/rand"
        "strconv"
        "time"
)


func dieRoll(size int) int { 
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(size) + 1
}

func rollTwo(size1, size2 int) (int, int) { 
    return dieRoll(size1), dieRoll(size2) 
}

func returnsNamed(input1 string, input2 int) (theResult string, err error) { 
    theResult = "modified" + input1 + ", " + strconv.Itoa(input2)
    return theResult, err
}

type dieRollFunc func(int) int

func fakeDieRoll(size int) int {
    return 42
}

func getDieRolls() []dieRollFunc {
    return []dieRollFunc{
        dieRoll,
        fakeDieRoll,
    }
}




func main() {
    fmt.Printf("Rolled a die of size %d, result: %d\n", 3, dieRoll(3))

    res1, res2 := rollTwo(3,7)
    fmt.Printf("Rolled a pair of dice (%d, %d), results: %d, %d\n", 3, 7, res1, res2)

    named, err := returnsNamed("globule", 42)
    fmt.Printf("Named params returned: '%s', %v\n", named, err)

    var rolls = getDieRolls()
    for index, rollFunc := range rolls { 
        fmt.Printf("Die Roll attempt #%d, result: %d\n", index, rollFunc(10))
    }
}

