package singleton02

import "sync"

// Singleton lazy man mode
type Singleton struct{}

var (
	lazySingleton *Singleton
	once          = &sync.Once{}
)

// GetLazyInstance lazy man mode
func GetLazyInstance() *Singleton {
	if lazySingleton == nil {
		once.Do(func() { lazySingleton = &Singleton{} })
	}
	return lazySingleton
}
