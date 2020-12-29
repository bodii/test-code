package learningcache

// Getter loads data for a key.
type Getter interface {
	Get(key string) ([]byte, error)
}

// GetterFunc implements Getter with a function.
type GetterFunc func(key string) ([]byte, error)

// Get implements Getter interface function
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}
