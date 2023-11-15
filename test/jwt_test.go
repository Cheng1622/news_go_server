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
	a := `eyJhbGciOiJQUzUxMiIsInR5cCI6IkpXVCJ9.eyJVc2VySWQiOjM3OTE5Mzc2NTIxMjA2NTc5MiwiZXhwIjoxNjk5OTc2OTY5LCJpYXQiOjE2OTk4OTA1NjksImlzcyI6ImNjLW5ld3MiLCJzdWIiOiJ1c2VydG9rZW4ifQ.i8emLtwSKcRRcSrRN3D8MZWuLfKOSE9RF73tJbqcgbpLLU8ZM7GL9JHPXU-XAthKzrrc-yrqSqWO9FSc3chsVOUwYM1X-Rxrs5DV5d20V7ZkZWtpVG-mJoH4rVygcsqUsn3EfTpfAuZts5MRY_rc8AjMtgDsidtKVTj2nteP_h4oX-EFgvzXtD8I9LIZNP0q9o8-OPRJWLabj9jHcJJxE680Vrjr7X7IPjtilg_EYA6mgfMiEYBxZ_b39N7EdWpWaF8l5lj2V9SGsGwC2vqfODHxkifFpojn-BkEX6N9G21Kr8K7qeBz2hqU_SndjEEwNnx2onnshctP-YTX0M2f_w`
	b, err := jwt.ParseToken(a)
	if err != nil {
		t.Fatal("err:", err)
	}
	t.Fatal("ok", b.Userid)
}
