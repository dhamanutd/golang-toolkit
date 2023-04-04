package models_test

import (
	"testing"
	"time"

	. "github.com/dhamanutd/golang-toolkit/models"
	"github.com/stretchr/testify/assert"
)

func TestDatetimeDefaultTimeNow(t *testing.T) {
	now := time.Now().UTC()
	timeNow := NewDateTimeNow()

	t.Log(now.String())
	t.Log(timeNow.String())

	y, m, d := now.Date()
	ty, tm, td := time.Time(timeNow).Date()

	assert.Equal(t, y, ty)
	assert.Equal(t, m, tm)
	assert.Equal(t, d, td)
}
