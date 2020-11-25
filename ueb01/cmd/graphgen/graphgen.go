package main

import (
    "flag"
    "fmt"
    "github.com/akelsch/vaa/ueb01/internal/errutil"
    "github.com/awalterschulze/gographviz"
    "log"
    "math/rand"
    "strconv"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}

func main() {
    var n, m int
    flag.IntVar(&n, "n", 5, "number of nodes n")
    flag.IntVar(&m, "m", 6, "number of edges m where m > n")
    flag.Parse()

    mMax := (n * (n - 1)) / 2
    if m > mMax {
        log.Fatalf("The number of edges m may not exceed %d", mMax)
    }

    // Effectively requiring n to be >= 4
    if m <= n {
        log.Fatal("The number of edges m must be greater than the number of nodes n")
    }

    // Create graph with nodes
    graph := gographviz.NewGraph()
    for i := 1; i <= n; i++ {
        err := graph.AddNode("", strconv.Itoa(i), map[string]string{})
        errutil.HandleError(err)
    }

    // Connect each node to another node
    matrix := initSquareMatrix(n)
    for i := 2; i <= n; i++ {
        j := getRandomNumber(1, i-1)
        err := graph.AddEdge(strconv.Itoa(i), strconv.Itoa(j), false, map[string]string{})
        errutil.HandleError(err)
        registerEdge(matrix, i, j)
    }

    // Fill with random remaining edges
    for existsRemainingEdge(matrix) && len(graph.Edges.Edges) < m {
        i, j := getRandomRemainingEdge(matrix)
        err := graph.AddEdge(strconv.Itoa(i), strconv.Itoa(j), false, map[string]string{})
        errutil.HandleError(err)
        registerEdge(matrix, i, j)
    }

    fmt.Print(graph)
}

func getRandomNumber(min, max int) int {
    return rand.Intn(max-min+1) + min
}

func initSquareMatrix(n int) *[][]bool {
    var matrix [][]bool
    for i := 0; i < n; i++ {
        matrix = append(matrix, []bool{})
        for j := 0; j < n; j++ {
            b := false
            // Diagonal with loops
            if i == j {
                b = true
            }
            matrix[i] = append(matrix[i], b)
        }
    }
    return &matrix
}

func registerEdge(matrix *[][]bool, i, j int) {
    p := i - 1
    q := j - 1
    // Unidirectional edges so (p,q)=(q,p)
    (*matrix)[p][q] = true
    (*matrix)[q][p] = true
}

func existsRemainingEdge(matrix *[][]bool) bool {
    for _, row := range *matrix {
        for _, elem := range row {
            if elem == false {
                return true
            }
        }
    }

    return false
}

func getRandomRemainingEdge(matrix *[][]bool) (int, int) {
    for _, randRow := range rand.Perm(len(*matrix)) {
        for _, randCol := range rand.Perm(len((*matrix)[randRow])) {
            if (*matrix)[randRow][randCol] == false {
                return randRow + 1, randCol + 1
            }
        }
    }

    return 0, 0
}
