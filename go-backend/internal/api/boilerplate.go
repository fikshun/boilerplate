package api

import (
  "crypto/md5"
  "encoding/hex"
  "fmt"
  "github.com/fikshun/go-backend/internal/sec"
  "github.com/gin-gonic/gin"
  "github.com/google/uuid"
  "gorm.io/gorm"
  "net/http"
  "sync"
)

type Boilerplate struct {
  db   *gorm.DB
  Lock sync.Mutex
}

func NewBoilerplate(d *gorm.DB) *Boilerplate {
  return &Boilerplate{db: d}
}

func (b *Boilerplate) Authenticate(c *gin.Context) {
  var authRequestUser User
  err := c.BindJSON(&authRequestUser)

  if err != nil {
    c.JSON(400, fmt.Sprintf("bad request: %s", err))
    return
  }

  var user User
  b.db.First(&user, "username = ?", authRequestUser.Username)

  if user.Id == nil {
    c.JSON(http.StatusUnauthorized, "unauthorized")
    return
  }

  hash := md5.Sum([]byte(*authRequestUser.Password))
  hashedPassword := hex.EncodeToString(hash[:])

  if hashedPassword != *user.Password {
    c.JSON(http.StatusUnauthorized, "unauthorized")
    return
  }

  var auth Auth

  auth.Username = user.Username
  token, err := sec.CreateToken(*user.Username)

  if err != nil {
    c.JSON(http.StatusInternalServerError, fmt.Sprintf("error creating token: %s", err))
    return
  }

  auth.Jwt = &token

  c.JSON(http.StatusOK, auth)
}

func (b *Boilerplate) CreateUser(c *gin.Context) {
  var user User
  err := c.BindJSON(&user)
  newUuid := uuid.New()
  user.Id = &newUuid
  hash := md5.Sum([]byte(*user.Password))
  hashedPassword := hex.EncodeToString(hash[:])
  user.Password = &hashedPassword

  if err != nil {
    c.JSON(400, fmt.Sprintf("bad request: %s", err))
    return
  }

  result := b.db.Create(&user)

  if result.Error != nil {
    c.JSON(http.StatusInternalServerError, fmt.Sprintf("error creating user: %s", result.Error))
    return
  }

  // obscure fields - find a better way to do this
  user.Password = nil
  c.JSON(http.StatusCreated, user)
}

func (b *Boilerplate) GetUser(c *gin.Context, uuid string) {
  var user User
  result := b.db.First(&user, uuid)
  if result.Error != nil {
    c.JSON(http.StatusInternalServerError, fmt.Sprintf("error getting user: %s", result.Error))
    return
  }

  c.JSON(http.StatusOK, user)
}

func (b *Boilerplate) UpdateUser(c *gin.Context, uuid string) {
  user := b.db.First(&User{}, uuid)

  if user.Error != nil {
    c.JSON(http.StatusInternalServerError, fmt.Sprintf("error updating user: %s", user.Error))
    return
  }

  var updatedUser User

  if err := c.ShouldBindJSON(&updatedUser); err != nil {
    c.JSON(400, fmt.Sprintf("bad request: %s", err))
    return
  }

  c.JSON(http.StatusOK, user)
}

func (b *Boilerplate) DeleteUser(c *gin.Context, uuid string) {
  result := b.db.Delete(&User{}, uuid)

  if result.Error != nil {
    c.JSON(http.StatusInternalServerError, fmt.Sprintf("error deleting user: %s", result.Error))
    return
  }
  c.JSON(http.StatusNoContent, "delete user")
}

func (b *Boilerplate) ListUsers(c *gin.Context) {
  var users []User
  result := b.db.Find(&users)

  if result.Error != nil {
    c.JSON(http.StatusInternalServerError, fmt.Sprintf("error listing users: %s", result.Error))
    return
  }

  c.JSON(http.StatusOK, users)
}
