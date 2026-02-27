//go:build amd64 && (linux || windows || darwin)

package fast

import (
	"github.com/zulfirman/socket.io-go/parser/json/serializer"
	"github.com/zulfirman/socket.io-go/parser/json/serializer/sonic"
)

func New() serializer.JSONSerializer {
	defaultConfig := DefaultConfig()
	return sonic.New(defaultConfig.SonicConfig)
}

func NewWithConfig(config Config) serializer.JSONSerializer {
	return sonic.New(config.SonicConfig)
}

func Type() SerializerType {
	return SerializerTypeSonic
}
