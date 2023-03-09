package kie_test

import (
	"github.com/arielsrv/go-archaius/source/remote/kie"
	"testing"

	"github.com/arielsrv/go-archaius"
	"github.com/arielsrv/go-archaius/source/remote"
	"github.com/stretchr/testify/assert"
)

func TestNewKieSource(t *testing.T) {
	opts := &archaius.RemoteInfo{
		DefaultDimension: map[string]string{
			remote.LabelApp:     "default",
			remote.LabelService: "cart",
		},
		TenantName: "default",
		URL:        "http://",
	}
	_, err := kie.NewKieSource(opts)
	assert.NoError(t, err)
}
