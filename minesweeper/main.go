package main

import (
    "fmt"
    "game/game"
    "github.com/inancgumus/screen"
    "github.com/mattn/go-tty"
)

func main() {
    result := _PlayConsole(10, 20)
    if (result) {
        fmt.Println("you win")
    } else {
        fmt.Println("you lose")
    }
}

func _PlayConsole(mapSize int, mineCount int) bool {
    tty, _ := tty.Open()
    defer tty.Close()

    game := minesweeper.CreateGame(mapSize, mineCount)

    screen.Clear()
    screen.MoveTopLeft()
    game.ShowGameMap(false)

    var isWin bool = false

    for {
        var end bool = false 
        r, err := tty.ReadRune()
        if err != nil {
            break
        }
        switch string(r) {
        case "z": 
            _, gameover := game.OpenSquare()
            if (gameover) {
                fmt.Println("oops!")
                isWin = false
                end = true
            }
        case "x": 
            game.ToggleMark()
        case "h": 
            game.MoveLeft()
        case "j": 
            game.MoveDown()
        case "k": 
            game.MoveUp()
        case "l": 
            game.MoveRight()
        case "e": 
            fmt.Println("give up")
            end = true
            break
        }
        isWin = game.IsClear()
        var showAnswer =  isWin || end
        if (showAnswer) {
            game.ShowGameMap(true)
            break
        }

        screen.Clear()
        screen.MoveTopLeft()
        game.ShowGameMap(false)
    }
    return isWin
}
