package api

import (
	"testing"
)

func TestGetAccessToken(t *testing.T) {
	client := NewClient("dingkimljpj2mycouknv", "tT73LNnaOaFExsmg2lALiXW9ODBhmvcD8QMcX676t5RB7diXURP_Imw9kPhw4-Dj")
	if client.AccessToken == "" {
		t.Error("Failed to get access_token")
	} else {
		t.Log(client.AccessToken)
	}
}
