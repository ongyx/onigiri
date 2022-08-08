package onigiri

// Event is a notification channel of T.
type Event[T any] struct {
	c chan T
}

// NewEvent creates a new event with a buffer of (size) values.
func NewEvent[T any](size int) *Event[T] {
	return &Event[T]{c: make(chan T, size)}
}

// Notify runs the callback function when a value is emitted.
func (e *Event[T]) Notify(f func(T)) {
	go func() {
		f(<-e.c)
	}()
}

// Emit sends the value to the event channel.
func (e *Event[T]) Emit(value T) {
	e.c <- value
}

// Poll checks if a value was emitted without blocking the current goroutine.
// If there is none, ok will be false.
func (e *Event[T]) Poll() (value T, ok bool) {
	select {
	case value = <-e.c:
		ok = true
	default:
	}

	return
}

// Close closes the event channel.
// No event methods should be called after this.
func (e *Event[T]) Close() {
	close(e.c)
}
