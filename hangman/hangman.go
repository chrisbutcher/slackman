package hangman

import (
  "encoding/json"
  "fmt"
  "strconv"
  "strings"
)

type GameState struct {
  WordToGuess      []string `json:"word_to_guess"`
  WordProgress     []string `json:"word_progress"`
  LettersGuessed   []string `json:"leters_guessed"`
  GuessesRemaining int      `json:"guesses_remaining"`
  GameOver         bool     `json:"game_over"`
  GameWon          bool     `json:"game_won"`
}

func (g GameState) String() string {
  bytes, _ := json.Marshal(g)
  return fmt.Sprintf("%s", bytes)
}

func (g *GameState) GameStatusLine() string {
  wordProgress := strings.Join(g.WordProgress, ",") + " "
  lettersGuessed := "Guesses: [" + strings.Join(g.LettersGuessed, ",") + "] "
  guessesRemaining := "Guesses left: " + strconv.Itoa(g.GuessesRemaining)

  return wordProgress + lettersGuessed + guessesRemaining
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

      if g.checkGameWon() {
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

func (g *GameState) checkGameWon() bool {
  gameWon := true
  for _, letter := range g.WordProgress {
    if letter == "_" {
      gameWon = false
    }
  }

  return gameWon
}
