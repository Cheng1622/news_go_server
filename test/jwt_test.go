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
	a := `eyJhbGciOiJQUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjM3ODAzODI4NTI0MDUwNDMyMCwiZXhwIjoxNjk5Njc2OTUwLCJpYXQiOjE2OTk1OTA1NTAsImlzcyI6ImNjLW5ld3MiLCJzdWIiOiJ1c2VydG9rZW4ifQ.fveWtfHNcYWi5SsySh9gfNzx8jc3552TW6lB6excYg6no5LvXLHjxz2L0eGh4WkFrE3_Y-cLnQNtG5X6-UErwCB3XFK28Q5FEyVImQxqzGQwFGYHjDARwrzHQqxy2AKXvGh-LOTw7R0pW_T3pHB9pDCZ48wPi29Uti47ZkIoHIe6t8q1AT2Y-Caa5BFJ4-NUrUHykjmWb4-Kj3QytBWLCq7OecoU6jvr3iU4QbDEoyDN3A3PygLAaM25SDoWrZ21oAtpmKAFvz1y0gJoinANWKFAjn8pGC3wCqhw462Nua0NEUrd0aSpoc7sLCEd30Obg3y8w5wC0r0bt2dwOVS9Hg`
	b, err := jwt.ParseToken(a)
	if err != nil {
		t.Fatal("err:", err)
	}
	t.Fatal("ok", b.UserId)
}
