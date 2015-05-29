package persistence

import (
  "encoding/json"
  "github.com/chrisbutcher/slackman/hangman"
  "github.com/garyburd/redigo/redis"
)

type SlackmanDB struct {
  Conn redis.Conn
}

func (d *SlackmanDB) Initialize(address, password string) {
  db, err := redis.Dial("tcp", address)
  if err != nil {
    panic(err)
  }

  if password != "" {
    _, err = db.Do("AUTH", password)
    if err != nil {
      panic(err)
    }
  }

  d.Conn = db
}

func (d *SlackmanDB) Close() {
  d.Conn.Close()
}

func (d *SlackmanDB) SetGameState(gameState hangman.GameState) {
  bytes, _ := json.Marshal(gameState)
  d.Conn.Do("SET", "gameState", bytes)
}

func (d *SlackmanDB) ReadGameState() hangman.GameState {
  newGameState := hangman.GameState{}
  result, _ := redis.String(d.Conn.Do("GET", "gameState"))

  if err := json.Unmarshal([]byte(result), &newGameState); err != nil {
    panic(err)
  }

  return newGameState
}
