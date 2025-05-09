package redis

import "encoding/json"

func (c *ObjectCacher[T]) Unmarshal(data []byte, out any) error {
	return json.Unmarshal(data, out)
}

func (c *ObjectCacher[T]) Marshal(in any) ([]byte, error) {
	return json.Marshal(in)
}
