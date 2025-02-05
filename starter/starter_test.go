package starter

import (
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestRun(t *testing.T) {
	done := make(chan error, 1)
	godotenv.Load("../.env") //Adding relative path to pickup correct env values.

	go func() {
		done <- Run()
	}()

	select {
	case err := <-done:
		if err != nil {
			t.Fatalf("run() failed: %v", err)
		}
	case <-time.After(10 * time.Second):
		t.Fatal("test timed out")
	}
}
