package paginator

import (
	jsoniter "github.com/json-iterator/go"
	"sync"
)

var j jsoniter.API
var once sync.Once

func get() jsoniter.API {
	once.Do(func() {
		j = jsoniter.ConfigCompatibleWithStandardLibrary
	})
	return j
}

func MarshalToString(v interface{}) (string, error) {
	return get().MarshalToString(v)
}
