package go_mobilus_sso

import (
	"testing"
)

var (
	secret = "qwertyuiop"
)

func TestNew(t *testing.T) {
	s := New(secret)
	if s.Secret != secret {
		t.Fatal("invalid secret")
	}
}

var last string

func TestMobilusSSO_Marshal(t *testing.T) {
	s := New(secret)

	t.Run("valid", func(t1 *testing.T) {
		c, err := s.Marshal(User{
			Name:        "テストマン",
			PermitLevel: 0,
			Token:       "thisistesttoken",
			DomainID:    "adm",
			PlusID:      "testman@example.com",
			UserID:      "testman@example.com",
			TenantID:    "mobilus",
		})
		if err != nil {
			t1.Fatal("marshal error", err)
		}
		if c == "" {
			t1.Fatal("marshal error", "empty string")
		}
		last = c
	})
	t.Run("invalid", func(t1 *testing.T) {
		_, err := s.Marshal(func() {})
		if err == nil {
			t1.Fatal("marshal error", err)
		}
	})
}

func TestMobilusSSO_Unmarshal(t *testing.T) {
	s := New(secret)
	t.Run("valid", func(t1 *testing.T) {
		user, err := s.Unmarshal(last)
		if err != nil {
			t1.Fatal("unmarshal error", err)
		}

		switch {
		case user.Name != "テストマン":
			t1.Fatal("unmarshal error", "invalid name", user.Name)
		case user.PermitLevel != 0:
			t1.Fatal("unmarshal error", "invalid permitLevel", user.PermitLevel)
		case user.Token != "thisistesttoken":
			t1.Fatal("unmarshal error", "invalid token", user.Token)
		case user.DomainID != "adm":
			t1.Fatal("unmarshal error", "invalid domainId", user.DomainID)
		case user.PlusID != "testman@example.com":
			t1.Fatal("unmarshal error", "invalid plusId", user.PlusID)
		case user.UserID != "testman@example.com":
			t1.Fatal("unmarshal error", "invalid userId", user.UserID)
		case user.TenantID != "mobilus":
			t1.Fatal("unmarshal error", "invalid tenantId", user.TenantID)
		}
	})
	t.Run("invalid", func(t1 *testing.T) {
		_, err := s.Unmarshal("")
		if err == nil {
			t1.Fatal("marshal error", err)
		}
	})
}
