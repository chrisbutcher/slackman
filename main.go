package main

import (
  // "flag"
  // "fmt"
  // "github.com/gin-gonic/gin"
  // "net/http"
  "fmt"
  "github.com/chrisbutcher/slackman/hangman"
  // "github.com/garyburd/redigo/redis"
  // "strings"
)

func main() {
  gameState := hangman.GameState{}
  gameState.Initialize("hangman")

  for {
    if gameState.GameOver || gameState.GameWon {
      break
    }

    fmt.Println(gameState)

    var input string
    _, _ = fmt.Scanf("%s", &input)

    gameState.GuessLetter(input)
  }

  fmt.Println(gameState)
}

// type SlackCommand struct {
//   Token       string `form:"token" binding:"required"`
//   TeamID      string `form:"team_id" binding:"required"`
//   TeamDomain  string `form:"team_domain" binding:"required"`
//   ChannelID   string `form:"channel_id" binding:"required"`
//   ChannelName string `form:"channel_name" binding:"required"`
//   UserID      string `form:"user_id" binding:"required"`
//   UserName    string `form:"user_name" binding:"required"`
//   Command     string `form:"command" binding:"required"`
//   Text        string `form:"text" binding:"required"`
// }

// func buildReply(slackCmd SlackCommand) string {
//   msg := "Hello " + slackCmd.UserName + ". "
//   msg += "[command: " + slackCmd.Command + ", "
//   msg += "text:" + slackCmd.Text + ", "
//   msg += "channel_name: " + slackCmd.ChannelName + "]"
//   return msg
// }

// func main() {
//   port := flag.String("port", "3000", "HTTP port")
//   slack_token := flag.String("slack_token", "debug", "Slack verification token")
//   flag.Parse()

//   r := gin.New()

//   // Global middleware
//   r.Use(gin.Logger())
//   r.Use(gin.Recovery())

//   r.GET("/", func(c *gin.Context) {
//     c.String(http.StatusOK, "bot active")
//   })

//   r.POST("/bot/command", func(c *gin.Context) {
//     var slackCmd SlackCommand
//     c.Bind(&slackCmd)

//     if slackCmd.Token != *slack_token {
//       c.String(http.StatusUnauthorized, "not authorized")
//       return
//     }

//     c.String(http.StatusOK, buildReply(slackCmd))
//   })

//   fmt.Println("Listening on port " + *port)
//   r.Run(":" + *port)
// }
