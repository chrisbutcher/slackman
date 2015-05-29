package hangman

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestInitialize(t *testing.T) {
  gameState := GameState{}
  gameState.Initialize("hangman")
  expectedWordToGuess := []string{"h", "a", "n", "g", "m", "a", "n"}
  expectedWordProgress := []string{"_", "_", "_", "_", "_", "_", "_"}
  expectedLettersGuessed := []string{}
  expectedGuessesRemaining := 6
  expectedGameOver := false

  assert.Equal(t, expectedWordToGuess, gameState.WordToGuess)
  assert.Equal(t, expectedWordProgress, gameState.WordProgress)
  assert.Equal(t, expectedLettersGuessed, gameState.LettersGuessed)
  assert.Equal(t, expectedGuessesRemaining, gameState.GuessesRemaining)
  assert.Equal(t, expectedGameOver, gameState.GameOver)
}

func TestInitializeInvalidCharacter(t *testing.T) {
  gameState := GameState{}
  gameState.Initialize("han_man")
  expectedGameOver := true

  assert.Equal(t, expectedGameOver, gameState.GameOver)
}

func TestGuessLetterCorrectly(t *testing.T) {
  gameState := GameState{}
  gameState.Initialize("hangman")
  gameState.GuessLetter("a")
  expectedLettersGuessed := []string{}
  expectedWordProgress := []string{"_", "a", "_", "_", "_", "a", "_"}
  expectedGuessesRemaining := 6
  expectedGameOver := false

  assert.Equal(t, expectedWordProgress, gameState.WordProgress)
  assert.Equal(t, expectedLettersGuessed, gameState.LettersGuessed)
  assert.Equal(t, expectedGuessesRemaining, gameState.GuessesRemaining)
  assert.Equal(t, expectedGameOver, gameState.GameOver)
}

func TestGuessLetterCorrectlyGameWon(t *testing.T) {
  gameState := GameState{}
  gameState.Initialize("hangman")
  gameState.GuessLetter("z")
  gameState.GuessLetter("h")
  gameState.GuessLetter("a")
  gameState.GuessLetter("n")
  gameState.GuessLetter("g")
  gameState.GuessLetter("m")
  gameState.GuessLetter("n")

  expectedLettersGuessed := []string{"z"}
  expectedWordProgress := []string{"h", "a", "n", "g", "m", "a", "n"}
  expectedGuessesRemaining := 5
  expectedGameOver := false
  expectedGameWon := true

  assert.Equal(t, expectedWordProgress, gameState.WordProgress)
  assert.Equal(t, expectedLettersGuessed, gameState.LettersGuessed)
  assert.Equal(t, expectedGuessesRemaining, gameState.GuessesRemaining)
  assert.Equal(t, expectedGameOver, gameState.GameOver)
  assert.Equal(t, expectedGameWon, gameState.GameWon)
}

func TestGuessLetterIncorrectly(t *testing.T) {
  gameState := GameState{}
  gameState.Initialize("hangman")
  gameState.GuessLetter("z")
  expectedLettersGuessed := []string{"z"}
  expectedWordProgress := []string{"_", "_", "_", "_", "_", "_", "_"}
  expectedGuessesRemaining := 5
  expectedGameOver := false

  assert.Equal(t, expectedWordProgress, gameState.WordProgress)
  assert.Equal(t, expectedLettersGuessed, gameState.LettersGuessed)
  assert.Equal(t, expectedGuessesRemaining, gameState.GuessesRemaining)
  assert.Equal(t, expectedGameOver, gameState.GameOver)
}

func TestGuessLetterIncorrectlyGameOver(t *testing.T) {
  gameState := GameState{}
  gameState.Initialize("hangman")
  // Correct guesses
  gameState.GuessLetter("a")
  gameState.GuessLetter("n")
  gameState.GuessLetter("g")
  gameState.GuessLetter("m")
  gameState.GuessLetter("n")

  // Wrong guesses
  gameState.GuessLetter("z")
  gameState.GuessLetter("z")
  gameState.GuessLetter("z")
  gameState.GuessLetter("z")
  gameState.GuessLetter("z")
  gameState.GuessLetter("z")

  expectedLettersGuessed := []string{"z"}
  expectedWordProgress := []string{"_", "a", "n", "g", "m", "a", "n"}
  expectedGuessesRemaining := 0
  expectedGameOver := true

  assert.Equal(t, expectedWordProgress, gameState.WordProgress)
  assert.Equal(t, expectedLettersGuessed, gameState.LettersGuessed)
  assert.Equal(t, expectedGuessesRemaining, gameState.GuessesRemaining)
  assert.Equal(t, expectedGameOver, gameState.GameOver)
}
