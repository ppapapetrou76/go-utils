package multierror

import (
	"errors"
	"fmt"
	"testing"

	"github.com/ppapapetrou76/go-testing/assert"
)

func TestPrefixedError_ErrorOrNil(t *testing.T) {
	ft := assert.NewFluentT(t)

	t.Run("should return nil if no errors", func(t *testing.T) {
		errs := NewPrefixed("some prefix")
		ft.AssertThat(errs.ErrorOrNil()).IsNil()
	})

	t.Run("should return the type if errors", func(t *testing.T) {
		err := errors.New("some error")
		errs := NewPrefixed("some prefix", err)
		ft.AssertThat(errs.ErrorOrNil()).
			IsNotNil().
			IsEqualTo(errs)

		ft.AssertThat(errors.As(errs, &err)).IsTrue()
		ft.AssertThat(errors.Is(errs, err)).IsTrue()
	})

	t.Run("should chain multiple errors", func(t *testing.T) {
		err1 := errors.New("error 1")
		err2 := errors.New("error 2")
		err3 := fmt.Errorf("some formatted error")
		errs := NewPrefixed("error example: ", err1, err2).ErrorOrNil()

		ft.AssertThat(errors.Is(errs, err2)).IsTrue()
		ft.AssertThat(errors.Is(errs, err1)).IsTrue()
		ft.AssertThat(errors.Is(errs, err3)).IsFalse()

		ft.AssertThat(errors.As(errs, &err2)).IsTrue()
		ft.AssertThat(errors.As(errs, &err1)).IsTrue()
		ft.AssertThat(errors.As(errs, &err3)).IsTrue()
	})
}

func TestPrefixedError_Error(t *testing.T) {
	ft := assert.NewFluentT(t)
	t.Run("should return empty string if no errors", func(t *testing.T) {
		errs := NewPrefixed("some prefix")
		ft.AssertThatString(errs.Error()).IsEmpty()
	})

	t.Run("should return empty string by ignoring nil errors", func(t *testing.T) {
		errs := NewPrefixed("some prefix", nil).Append(nil)
		ft.AssertThatString(errs.Error()).IsEmpty()
	})

	t.Run("should return expected string when one error", func(t *testing.T) {
		errs := NewPrefixed("some prefix:", errors.New("some error"))
		ft.AssertThat(errs.Error()).IsEqualTo("some prefix: some error")
	})

	t.Run("should return expected string when multiple errors", func(t *testing.T) {
		errs := NewPrefixed("some prefix:", errors.New("some error"))
		errs = errs.Append(errors.New("second error"))
		ft.AssertThat(errs.Error()).IsEqualTo("some prefix: 2 errors:[ some error, second error ]")
	})

	t.Run("should return expected string when wrapped errors", func(t *testing.T) {
		errs := NewPrefixed("some prefix:", fmt.Errorf("some wrapped error %w", errors.New("some error")))
		errs = errs.Append(errors.New("second error"))
		ft.AssertThat(errs.Error()).IsEqualTo("some prefix: 2 errors:[ some wrapped error some error, second error ]")
	})

	t.Run("should return expected string with custom formatter", func(t *testing.T) {
		errs := NewPrefixed("some prefix:", errors.New("some error")).
			WithFormatFunc(func(errs []error) string {
				return fmt.Sprintf("custom formatter %d", len(errs))
			})
		ft.AssertThat(errs.Error()).IsEqualTo("some prefix: custom formatter 1")
	})
}
