package singleton01

// Singleton hungry man mode
type Singleton struct{}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

// GetInstance get instance
func GetInstance() *Singleton {
	return singleton
}
