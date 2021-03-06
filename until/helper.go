package until

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/rest-layer/schema"
	"gorm.io/gorm"
)

var TIME_FORMAT = "2006-01-02 15:04:05"

// Response ...
type Response struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

// WriteResponse ...
func WriteResponse(c *gin.Context, code int, data gin.H, err error) {
	var msgErr string
	if err != nil {
		msgErr = err.Error()
	}

	response := &Response{
		Code:  code,
		Data:  data,
		Error: msgErr,
	}

	c.JSON(200, response)
	c.Abort()
}

func GetMD5Hash(login, password string) string {
	hash := md5.Sum([]byte(login + password + time.Now().String()))
	return hex.EncodeToString(hash[:])
}

type MultiString []string

func (s MultiString) GormDataType() string {
	result := ""
	for _, v := range s {
		result += v
	}
	return "text"
}

func (s MultiString) GormDBDataType(db *gorm.DB, field *schema.Field) string {
	fmt.Println(db.Dialector.Name())
	switch db.Dialector.Name() {
	case "mysql", "sqlite", "postgres":
		return "text"
	}
	return ""
}
