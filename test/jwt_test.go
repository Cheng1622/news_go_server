package test

import (
	"testing"

	"github.com/Cheng1622/news_go_server/pkg/config"
	"github.com/Cheng1622/news_go_server/pkg/jwt"
	"github.com/Cheng1622/news_go_server/pkg/snowflake"
)

func TestGenToken(t *testing.T) {
	config.InitConfig()
	snowflake.InitSnowflake()
	a, _ := snowflake.SF.GenerateID()
	t.Log(a)
	b, err := jwt.GenToken(a)
	if err != nil {
		t.Fatal("err:", err)
	}
	t.Fatal("ok", b)
}

func TestParseToken(t *testing.T) {
	config.InitConfig()
	a := `eyJhbGciOiJQUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjM3NzQ1NTg4ODkwNDA5Nzc5MiwiZXhwIjoxNjk5NTM4MDk2LCJpYXQiOjE2OTk0NTE2OTZ9.HhtgLTkA0-DaHdWYP_9MKvDX_JeG5ac3K0hPtGAY4Q26SG0vQQUZyfcXz4etKTC229dePGVwkOf3jrRhUyf6GzbgO_XFcpqSoymHF5jWKJKKPelITM8LKQ6iBfzJv8dyDwnQxarRDzPSdTKFXdQZLYlcJ6Ogj7oDtsvBC8RHfw39UG9-GmHNd5DrhRLIVm0DoUIlZbawt4livkeyZvPD2legMOf1fxa4je_u46ZUSJykrKUQqSB5KDKLtg9rJ0DJMsFKhgYRRwDDVVDg6xDYlzJRd2tlOiCA-ibcZZ3ARndsOvXLJiUGUEDRt5cMLdpBk8HUTM7aykCHA3wr02tQuA`
	b, err := jwt.ParseToken(a)
	if err != nil {
		t.Fatal("err:", err)
	}
	t.Fatal("ok", b.UserId)
}
