package cluster

type Components interface {
	Create()
	Delete()
	Setup()
}
