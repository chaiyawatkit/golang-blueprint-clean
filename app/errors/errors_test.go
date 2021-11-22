package errors_test

import (
	"golang-blueprint-clean/app/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInternalError_Error(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		messageError := "internal error"
		e := errors.InternalError{
			Message: messageError,
		}
		assert.Equal(t, messageError, e.Error())
	})

	t.Run("Success", func(t *testing.T) {
		messageError := "record not found error"
		e := errors.RecordNotFoundError{
			Message: messageError,
		}
		assert.Equal(t, messageError, e.Error())
	})

	t.Run("Success", func(t *testing.T) {
		messageError := "unprocessable error"
		e := errors.UnprocessableEntity{
			Message: messageError,
		}
		assert.Equal(t, messageError, e.Error())
	})

	t.Run("Success", func(t *testing.T) {
		messageError := "parameter error"
		e := errors.ParameterError{
			Message: messageError,
		}
		assert.Equal(t, messageError, e.Error())
	})

	t.Run("Success", func(t *testing.T) {
		messageError := "violates foreign key constraint"
		e := errors.ForeignKeyConstraintError{
			Message: messageError,
		}
		assert.Equal(t, messageError, e.Error())
	})

	t.Run("Success", func(t *testing.T) {
		messageError := "duplicate key value violates unique constraint"
		e := errors.UniqueKeyConstraintError{
			Message: messageError,
		}
		assert.Equal(t, messageError, e.Error())
	})

}
