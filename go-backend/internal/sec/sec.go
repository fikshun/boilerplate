package sec

import (
  "github.com/gin-gonic/gin"
  "github.com/golang-jwt/jwt/v5"
  "os"
  "time"
)

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func CreateToken(username string) (string, error) {
  token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
    "issuer":   "boilerplate",
    "audience": "boilerplate",
    "username": username,
    "exp":      time.Now().Add(time.Hour * 72).Unix(),
  })

  tokenString, err := token.SignedString(secretKey)

  return tokenString, err
}

func TokenAuthMiddleWare() gin.HandlerFunc {
  return func(c *gin.Context) {
    if c.Request.URL.Path == "/authenticate" {
      c.Next()
      return
    }

    tokenString := c.GetHeader("Authorization")

    if tokenString == "" {
      c.JSON(401, "Unauthorized")
      c.Abort()
      return
    }

    tokenString = tokenString[7:]

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      return secretKey, nil
    })

    if err != nil {
      c.JSON(401, "Unauthorized")
      c.Abort()
      return
    }

    if !token.Valid {
      c.JSON(401, "Unauthorized")
      c.Abort()
      return
    }

    c.Next()
  }
}
