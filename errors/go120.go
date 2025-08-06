package errors

import stderrors "errors"

// Join returns an error of calling errors.Join from the standard library.
func Join(errs ...error) error {
	return stderrors.Join(errs...)
}
