package singleton

type Singleton interface {
	AddOne() int
}

type singleton struct {
}

var instance *singleton

func GetInstance() Singleton {}
