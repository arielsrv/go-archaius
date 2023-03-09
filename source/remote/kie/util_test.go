package kie_test

import (
	"github.com/arielsrv/go-archaius/source/remote/kie"
	"testing"

	"github.com/arielsrv/go-archaius/source/remote"
	"github.com/stretchr/testify/assert"
)

func TestGenerateLabels(t *testing.T) {
	optionsLabels := map[string]string{
		remote.LabelApp:         "app",
		remote.LabelEnvironment: "env",
		remote.LabelService:     "service",
		"foo":                   "bar",
	}
	dimensionApp, err := kie.GenerateLabels(kie.DimensionApp, optionsLabels)
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{
		remote.LabelApp:         "app",
		remote.LabelEnvironment: "env",
	}, dimensionApp)

	dimensionService, err := kie.GenerateLabels(kie.DimensionService, optionsLabels)
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{
		remote.LabelApp:         "app",
		remote.LabelEnvironment: "env",
		remote.LabelService:     "service",
	}, dimensionService)
}
