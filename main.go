package main

import (
  "flag"
  "fmt"
  "github.com/chrisbutcher/slackman/hangman"
  "github.com/chrisbutcher/slackman/persistence"
  "github.com/gin-gonic/gin"
  "net/http"
  // "regexp"
  // "strings"
)

// binding:"required"
type SlackCommand struct {
  Token       string `form:"token"`
  TeamID      string `form:"team_id"`
  TeamDomain  string `form:"team_domain"`
  ChannelID   string `form:"channel_id"`
  ChannelName string `form:"channel_name"`
  UserID      string `form:"user_id"`
  UserName    string `form:"user_name"`
  Command     string `form:"command"`
  Text        string `form:"text"`
}

func main() {
  port := flag.String("port", "3000", "HTTP port")
  slack_token := flag.String("slack_token", "debug", "Slack verification token")
  redis_url := flag.String("redis_url", "[localhost]:6379", "Redis address")
  redis_pw := flag.String("redis_pw", "", "Redis password")
  flag.Parse()

  gameState := hangman.GameState{}
  gameState.Initialize("hangman")

  db := persistence.SlackmanDB{}
  db.Initialize(*redis_url, *redis_pw)
  defer db.Close()

  db.SetGameState(gameState)

  r := gin.New()

  // Global middleware
  r.Use(gin.Logger())
  r.Use(gin.Recovery())

  r.GET("/", func(c *gin.Context) {
    c.String(http.StatusOK, "bot active")
  })

  r.POST("/command", func(c *gin.Context) {
    var slackCmd SlackCommand
    c.Bind(&slackCmd)

    if slackCmd.Token != *slack_token {
      c.String(http.StatusUnauthorized, "not authorized")
      return
    }

    redisGameState := db.ReadGameState()

    if slackCmd.Command == "/hang" {
      if slackCmd.Text != "" && len(slackCmd.Text) == 1 {
        redisGameState.GuessLetter(slackCmd.Text)

        if redisGameState.GameWon {
          c.String(http.StatusOK, "Game won! "+redisGameState.GameStatusLine())
          redisGameState.Initialize("hangman")
          db.SetGameState(redisGameState)
        } else if redisGameState.GameOver {
          c.String(http.StatusOK, "Game lost! "+redisGameState.GameStatusLine())
          redisGameState.Initialize("hangman")
          db.SetGameState(redisGameState)
        } else {
          db.SetGameState(redisGameState)
          c.String(http.StatusOK, redisGameState.GameStatusLine())
        }
      } else {
        c.String(http.StatusOK, "Illegal guess")
      }
    } else {
      c.String(http.StatusOK, "Unrecognized command")
    }
  })

  fmt.Println("Listening on port " + *port)
  r.Run(":" + *port)
}
