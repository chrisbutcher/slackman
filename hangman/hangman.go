package hangman

import (
  "strings"
)

type GameState struct {
  WordToGuess      []string
  WordProgress     []string
  LettersGuessed   []string
  GuessesRemaining int
  GameOver         bool
  GameWon          bool
}

func (g *GameState) Initialize(rawWordToGuess string) {
  if strings.Index(rawWordToGuess, "_") != -1 {
    g.GameOver = true
    return
  }

  g.WordToGuess = strings.Split(rawWordToGuess, "")

  g.WordProgress = make([]string, len(g.WordToGuess))
  for i, _ := range g.WordToGuess {
    g.WordProgress[i] = "_"
  }

  g.LettersGuessed = make([]string, 0)

  g.GuessesRemaining = 6 // head = 5 (remaining), body, arm, arm, leg, leg = 0

  g.GameOver = false
  g.GameWon = false
}

func (g *GameState) GuessLetter(letterGuessed string) {
  letterFound := false

  for i, letter := range g.WordToGuess {
    if letter == letterGuessed {
      letterFound = true
      g.WordProgress[i] = letterGuessed

      if g.CheckGameWon() {
        g.GameWon = true
        return
      }
    }
  }

  if letterFound == false {
    if !g.LetterAlreadyGuessed(letterGuessed) {
      g.LettersGuessed = append(g.LettersGuessed, letterGuessed)
    }
    g.GuessesRemaining -= 1

    if g.GuessesRemaining == 0 {
      g.GameOver = true
    }
  }
}

func (g *GameState) LetterAlreadyGuessed(letterGuessed string) bool {
  for _, letter := range g.LettersGuessed {
    if letter == letterGuessed {
      return true
    }
  }

  return false
}

func (g *GameState) CheckGameWon() bool {
  gameWon := true
  for _, letter := range g.WordProgress {
    if letter == "_" {
      gameWon = false
    }
  }

  return gameWon
}
