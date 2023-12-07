package main

import (
	"testing"
	// "database/sql"
	"fmt"
)

func TestNewPostgresStore(t *testing.T) {
	store, err := NewPostgresStore()
	t.Run("Testing store", func(t *testing.T) {
		if err != nil {
			got := "Not nil value for err"
			want := "Nil val for err"
			assertCorrectMsg(t, got, want)
		} else {
			got := "Nil val for err"
			want := "Nil val for err"
			assertCorrectMsg(t, got, want)
		}
	})

	t.Run("Testing db nil", func(t *testing.T) {
		got := store.db
		want := (got != nil)

		if want == true {
			gotMsg := "Db is nil"
			wantMsg := "Db is nil"
			assertCorrectMsg(t, gotMsg, wantMsg)
		} else {
			gotMsg := "Db is not nil"
			wantMsg := "Db is nil"
			assertCorrectMsg(t, gotMsg, wantMsg)
		}
	})

	t.Run("Pinging DB", func(t *testing.T) {
		err = store.db.Ping()
		fmtErrMsg := fmt.Sprintf("Error pinging db: %v", err)
		if err != nil {
			assertCorrectMsg(t, fmtErrMsg, "Pinged db")
		} else {
			assertCorrectMsg(t, "Pinged db", "Pinged db")
		}
	})

	defer func() {
		if err := store.db.Close(); err != nil {
			t.Fatalf("Error closing database connection: %v", err)
		}
	}()
}

func assertCorrectMsg(t testing.TB, got, want string) {
	t.Helper()
	if got != fmt.Sprint(want) {
		t.Errorf("Recebido %q want %q", got, want)
	}
}
