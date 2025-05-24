package response

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"go-template/internal/exceptions"
)

func Test_newResponse(t *testing.T) {
	var message = http.StatusText(http.StatusNotFound)

	var r = newResponse(http.StatusNotFound, message)

	assert.Equal(t, message, r.Message)
	assert.Equal(t, http.StatusNotFound, r.Code)
	assert.Equal(t, http.StatusOK, r.HeaderCode)

	r = newResponse(http.StatusNotFound, message)

	r.WithPayload("payload")

	assert.Equal(t, message, r.Message)
	assert.Equal(t, http.StatusNotFound, r.Code)
	assert.Equal(t, http.StatusOK, r.HeaderCode)
	assert.Equal(t, "payload", r.Payload)
}

func Test_Ok(t *testing.T) {
	var r = Ok()

	assert.Equal(t, http.StatusOK, r.Code)
	assert.Equal(t, http.StatusOK, r.HeaderCode)
}

func Test_Err(t *testing.T) {
	var message = "InternalErr"

	var r = Err(assert.AnError)

	assert.Equal(t, message, r.Message)
	assert.Equal(t, http.StatusInternalServerError, r.Code)
	assert.Equal(t, http.StatusOK, r.HeaderCode)

	message = "BadRequest"

	r = Err(exceptions.ErrBadRequest)

	assert.Equal(t, message, r.Message)
	assert.Equal(t, http.StatusBadRequest, r.Code)
	assert.Equal(t, http.StatusOK, r.HeaderCode)

	message = "Success"

	r = Err(nil)

	assert.Equal(t, message, r.Message)
	assert.Equal(t, http.StatusOK, r.Code)
	assert.Equal(t, http.StatusOK, r.HeaderCode)
}
