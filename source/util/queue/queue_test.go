package queue_test

import (
	"errors"
	"github.com/arielsrv/go-archaius/source/util/queue"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParallelize(t *testing.T) {
	errs := []error{nil}
	f := func(i int, errCh chan error) {
		if errs[i] != nil {
			errCh <- errs[i]
		}
	}
	err := queue.Concurrent(len(errs), len(errs), f)
	assert.NoError(t, err)

	errs = append(errs, errors.New("error string"))
	err = queue.Concurrent(len(errs), len(errs), f)
	assert.Error(t, err)
}
