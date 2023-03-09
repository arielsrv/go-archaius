package serializers_test

import (
	"github.com/arielsrv/go-archaius/pkg/serializers"
	"testing"

	"github.com/arielsrv/go-archaius/pkg/serializers/json"
)

type Test struct {
	Team string `json:"team"`
}

func Test_Encode1(t *testing.T) {
	t.Log("Testing serializer encoding function for valid serializer")
	serializers.AvailableSerializers = make(map[string]serializers.Serializer)
	serializers.AvailableSerializers[serializers.JSONEncoder] = json.Serializer{}

	test := &Test{Team: "data"}
	data, _ := serializers.Encode(serializers.JSONEncoder, test)

	stringData := `{"team":"data"}`
	if string(data) != stringData {
		t.Error("error is encoding the data")
	}
}

func Test_Encode2(t *testing.T) {
	t.Log("Testing serializer encoding function for invalid serializer")
	serializers.AvailableSerializers = make(map[string]serializers.Serializer)
	serializers.AvailableSerializers[serializers.JSONEncoder] = json.Serializer{}

	test := &Test{Team: "data"}
	_, err := serializers.Encode("Invalidserializer", test)
	if err == nil {
		t.Error("Encoder is encoding invalid type of serilizer format")
	}
}

func Test_Decode(t *testing.T) {
	t.Log("Testing serializer decode function")
	serializers.AvailableSerializers = make(map[string]serializers.Serializer)
	serializers.AvailableSerializers[serializers.JSONEncoder] = json.Serializer{}
	test := &Test{Team: "data"}

	data, _ := serializers.Encode(serializers.JSONEncoder, test)
	err := serializers.Decode(serializers.JSONEncoder, data, test)

	if err != nil {
		t.Error("error in decoding data")
	}
}
