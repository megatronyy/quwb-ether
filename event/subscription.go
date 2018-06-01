package event

type Subscription interface {
	Err() <-chan error
	Unsubscribe()
}