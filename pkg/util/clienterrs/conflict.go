package clienterrs

import (
	apierrors "k8s.io/apimachinery/pkg/api/errors"
)

// OnlyConflictErrors returns true when the error only contains conflict errors.
// This means that apierrors.IsConflict returns true, or
// the error implements Unwrap() []error and apierrors.IsConflict returns true for all of them.
func OnlyConflictErrors(err error) bool {
	if ex, ok := err.(interface{ Unwrap() []error }); ok {
		for _, err := range ex.Unwrap() {
			if !apierrors.IsConflict(err) {
				return false
			}
		}
		return true
	}
	return apierrors.IsConflict(err)
}
