package treer

// Treer exposes the Trie structure capabilities.
type Treer interface {
	Get(key string) interface{}
	Put(key string, value interface{}) bool
	Delete(key string) bool
	Walk(walker WalkFunc) error
}
