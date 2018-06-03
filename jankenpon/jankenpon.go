package jankenpon

import "fmt"

var RULES = [3]int{2, 0, 1}

const (
  IDLE = "idle"
  WAITING = "waiting"
  FINISHED = "finished"
)

var (
  state string = IDLE
  choice = [2]int{-1, -1}
  score = [2]int{0, 0}
  tie bool = false
  winner int = -1
)

type status struct {
  State  string `json:"state"`
  Score  [2]int `json:"score"`
  Tie    bool   `json:"tie"`
  Winner int    `json:"winner"`
}

func Status () status {
  return status{state, score, tie, winner}
}

func Reset () {
  state = IDLE
  choice = [2]int{-1, -1}
  tie = false
  winner = -1
}

func Choose (player, shape int) {
  if choice[player] > -1 {
    return
  }

  state = WAITING
  choice[player] = shape


  fmt.Println("Choose", choice[0], choice[1])

  if choice[0] > -1 && choice[1] > -1 {
    play()
  }
}

func play () {
  state = FINISHED
  shape1, shape2 := choice[0], choice[1]

  if shape1 == shape2 {
    tie = true
  } else if RULES[shape1] == shape2 {
    winner = 0
    score[0]++
  } else {
    winner = 1
    score[1]++
  }
}

