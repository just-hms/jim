package actions

import (
	"jim/pkg/io"
	"jim/pkg/models"
	"jim/pkg/rainbow"
	"testing"
	"time"

	"github.com/go-playground/assert"
)

func TestShow(t *testing.T) {

	rainbow.Blank()
	mockResponse := ""

	session := models.Session{
		Start:   time.Now(),
		Elapsed: time.Since(time.Now()),
		Command: "",
	}

	session.Save()

	responseData := io.InterceptStdout(func() {
		Show.Value([]string{})
	})

	// check if the wacth result is in the db
	assert.NotEqual(t, responseData, mockResponse)
}
