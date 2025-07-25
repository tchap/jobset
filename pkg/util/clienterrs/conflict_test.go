package clienterrs

import (
	"errors"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"testing"
)

func TestOnlyConflictErrors(t *testing.T) {
	otherErr := errors.New("nuked")
	conflictErr := apierrors.NewConflict(schema.GroupResource{}, "conflict", errors.New("resource updated"))

	testCases := []struct {
		name         string
		err          error
		onlyConflict bool
	}{
		{
			name:         "nil error",
			err:          nil,
			onlyConflict: false,
		},
		{
			name:         "simple other error",
			err:          otherErr,
			onlyConflict: false,
		},
		{
			name:         "simple conflict error",
			err:          conflictErr,
			onlyConflict: true,
		},
		{
			name:         "mixed error slice",
			err:          errors.Join(otherErr, conflictErr),
			onlyConflict: false,
		},
		{
			name:         "conflict error slice",
			err:          errors.Join(conflictErr, conflictErr),
			onlyConflict: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := OnlyConflictErrors(tc.err); got != tc.onlyConflict {
				t.Errorf("Expected OnlyConflictErrors to return %v, got %v", tc.onlyConflict, got)
			}
		})
	}
}
