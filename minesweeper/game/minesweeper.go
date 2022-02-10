package minesweeper

import (
    "fmt"
    "time"
    "math/rand"
    "github.com/fatih/color"
)

type vector struct {
    x int
    y int
}

type square struct {
    x int
    y int
    isMine bool
    isMark bool
    isZero bool
    mineCount int
    isOpen bool
}

type game struct {
    gameMap [][]square
    mapSize int
    mineCount int
    cursor vector
    remaining int
}

func CreateGame(mapSize int, mineCount int) game {
    rand.Seed(time.Now().UnixNano())

    gm := make([][]square, mapSize)
    for i:=0; i<mapSize; i++{
        gm[i] = make([]square, mapSize)
    }

    mines := make([]vector, mineCount)
    mineSeeds := _createMines(mapSize, mineCount)
    for i:=0; i<mineCount; i++{
        mines[i] = _toVector(mapSize, mineSeeds[i])
    }

    g := game {
        gameMap : gm,
        mapSize : mapSize,
        mineCount : mineCount,
        cursor : vector {x:0, y:0},
    }

    g._addMarker(mines)

    maxSize := mapSize * mapSize
    g.remaining = maxSize - mineCount

    return g
}

func (g game)ShowGameMap(isAll bool) {
    fmt.Print("   ")
    for i:=0; i<g.mapSize; i++{
        fmt.Printf("%d ", i)
    }
    fmt.Print("\n   ")
    for i:=0; i<g.mapSize; i++{
        fmt.Print("--")
    }
    fmt.Print("\n")

    colors := make([]*color.Color, 0)
    colors = append(colors, color.New(color.FgBlack, color.BgHiBlack))
    colors = append(colors, color.New(color.FgHiBlue, color.BgHiBlack))
    colors = append(colors, color.New(color.FgHiGreen, color.BgHiBlack))
    colors = append(colors, color.New(color.FgHiYellow, color.BgHiBlack))
    colors = append(colors, color.New(color.FgHiMagenta, color.BgHiBlack))
    colors = append(colors, color.New(color.FgHiRed, color.BgHiBlack))

    cy := color.New(color.FgCyan, color.BgHiBlack)
    gr := color.New(color.FgGreen, color.BgHiBlack)
    rd := color.New(color.FgRed, color.BgHiBlack)
    wt := color.New(color.FgWhite, color.BgHiBlack)

    min := func(a int, b int) int {
        if ( a > b) {
            return b
        }
        return a
    }

    colorStr := func(x int, y int, str string, colorNo int) {
        if (g.cursor.x == x && g.cursor.y == y) {
           wt.Print(str)
        } else if (g.cursor.x == x || g.cursor.y == y) {
           cy.Print(str)
        } else {
            c := colors[min(colorNo, len(colors) - 1)]
            c.Print(str)
       }
    }

    for y, vy := range g.gameMap {
        fmt.Printf("%d |", y)
        for x, v := range vy {
            if (isAll || v.isOpen || v.isMark) {
                if (v.isMark) {
                    gr.Print("m ")
                } else if (v.isMine){
                    rd.Print("x ")
                } else {
                    colorStr(x, y, fmt.Sprintf("%d ", v.mineCount), v.mineCount)
                }
            } else {
                colorStr(x, y, "- ", 0)
            }
        }
        fmt.Println("")
    }
    fmt.Printf("remaining: %d\n", g.remaining)
}

func (g *game) OpenSquare() (bool, bool) {
    if (!g._inMap(g.cursor.x, g.cursor.y)) {
        return false, false
    }
    s := &g.gameMap[g.cursor.y][g.cursor.x]

    return g.OpenSquareBySquare(s)

}

func (g *game) OpenSquareBySquare(s *square) (bool, bool) {

    if (s.isMark || s.isOpen) {
        return false, false
    }

    if (s.isMine) {
        return false, true
    }

    g.remaining -= 1
    s.isOpen = true
    if (s.isZero) {
        fmt.Println(s.isZero, s.mineCount)
        surrounds := g._getSurrounds(s.x, s.y)
        for _, v := range surrounds {
            g.OpenSquareBySquare(v)
        }
    }

    return true, false
}

func (g *game) ToggleMark() (bool) {
    if (!g._inMap(g.cursor.x, g.cursor.y)) {
        return false
    }
    s := &g.gameMap[g.cursor.y][g.cursor.x]

    if (s.isOpen) {
        return false
    }

    s.isMark = !s.isMark

    return true
}

func (g *game) IsClear() (bool) {
    return g.remaining <= 0
}

func (g *game)MoveUp() (bool) {
    if (!_inRange(g.cursor.y - 1, g.mapSize)) {
        return false
    }
    g.cursor.y -= 1
    return true
}
func (g *game)MoveDown() (bool) {
    if (!_inRange(g.cursor.y + 1, g.mapSize)) {
        return false
    }
    g.cursor.y += 1
    return true
}
func (g *game)MoveLeft() (bool) {
    if (!_inRange(g.cursor.x - 1, g.mapSize)) {
        return false
    }
    g.cursor.x -= 1
    return true
}
func (g *game)MoveRight() (bool) {
    if (!_inRange(g.cursor.x + 1, g.mapSize)) {
        return false
    }
    g.cursor.x += 1
    return true
}

func (g *game)_addMarker(mines []vector) {
    for _, v := range mines {
        g._setMine(&g.gameMap[v.y][v.x])
    }

    for y, vy := range g.gameMap {
        for x, _ := range vy {
            g._setCount(x, y)
        }
    }
}

func _createMines(mapSize int, mineCount int) []int{
    maxSize := mapSize * mapSize
    randomSeeds := make([]int, maxSize)
    for i, _ := range randomSeeds {
        randomSeeds[i] = i
    }
    rand.Shuffle(len(randomSeeds), func(i, j int) { randomSeeds[i], randomSeeds[j] = randomSeeds[j], randomSeeds[i] })
    return randomSeeds
}

func (g *game)_setMine(m *square) {
    m.isMine = true
    m.isZero = false
    m.mineCount = 0
}

func (g *game)_setCount(x int, y int) {
    surrounds := g._getSurrounds(x, y)

    counter := 0
    for _, v := range surrounds {
        if (v.isMine){
            counter += 1
        }
    }
    m := &g.gameMap[y][x]

    m.x = x
    m.y = y
    m.isZero = true
    if (!m.isMine) {
        m.mineCount += counter
        m.isZero = (counter == 0)
    }
}

func (g *game)_getSurrounds(x int, y int) []*square {
    startX := x - 1
    endX := x + 1
    startY := y - 1
    endY := y + 1
    results := make([]*square, 0)
    for i := startX; i <= endX; i++ {
        for l := startY; l <= endY; l++ {
            if (g._inMap(i, l)) {
                results = append(results, &g.gameMap[l][i])
            }
        }
    }
    return results
}

func (g *game)_inMap(x int, y int) bool {
    return _inRange(x, g.mapSize) && _inRange(y, g.mapSize)
}


func _inRange(n int, max int) bool {
    return n >= 0 && n < max
}

func _toVector(mapSize int, no int) vector {
    x := no % mapSize
    y := no / mapSize
    return vector{x:x, y:y}
}

