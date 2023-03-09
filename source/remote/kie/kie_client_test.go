package kie_test

import (
	"github.com/arielsrv/go-archaius/source/remote/kie"
	"strconv"
	"testing"

	"github.com/arielsrv/go-archaius/source/remote"
	client "github.com/go-chassis/kie-client"
	"github.com/stretchr/testify/assert"
)

func TestNewKie(t *testing.T) {
	k, err := kie.NewKie(remote.Options{
		ServerURI: "http://",
		Labels:    map[string]string{remote.LabelApp: "default"}})
	assert.NoError(t, err)
	assert.Equal(t, "default", k.Options().Labels[remote.LabelApp])
}

func TestMergeConfig(t *testing.T) {
	k, err := kie.NewKie(remote.Options{
		ServerURI: "http://",
		Labels: map[string]string{
			remote.LabelApp:         "app",
			remote.LabelEnvironment: "env",
			remote.LabelService:     "service",
			remote.LabelVersion:     "1.0.0",
		}})
	assert.NoError(t, err)
	for i, dimension := range kie.DimensionPrecedence {
		k.SetDimensionConfigs(&client.KVResponse{
			Data: []*client.KVDoc{
				{
					Key:    "foo",
					Status: "enabled",
					Value:  strconv.Itoa(i + 1),
				},
			},
		}, dimension)
	}
	assert.Equal(t, strconv.Itoa(len(kie.DimensionPrecedence)), k.MergeConfig()["foo"].(string))
}
