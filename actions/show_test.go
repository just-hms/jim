package actions

import (
	"jim/models"
	"jim/rainbow"
	"jim/utils"
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

	responseData := utils.InterceptStdout(func() {
		Show.Value([]string{})
	})

	// check if the wacth result is in the db
	assert.NotEqual(t, responseData, mockResponse)
}
