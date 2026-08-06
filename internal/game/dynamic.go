package game

import "fmt"

type dynamic_state uint8

const (
	unresolved dynamic_state = 0
	resolving  dynamic_state = 1
	resolved   dynamic_state = 2
)

type Dynamic[T any] struct {
	value    T
	resolved T
	state    dynamic_state
}

func NewDynamic[T any](value T) Dynamic[T] {
	return Dynamic[T]{
		value:    value,
		resolved: value,
		state:    unresolved,
	}
}

func (d *Dynamic[T]) Mutate(updater func(*T)) error {
	if d.state == resolving {
		return fmt.Errorf("!!! Tried to mutate state inside of resolve()")
	}

	updater(&d.value)
	d.state = unresolved
	return nil
}
func (d *Dynamic[T]) Modify(updater func(*T)) error {
	if d.state != resolving {
		return fmt.Errorf("!!! Tried to modify state outside of resolve()")
	}

	updater(&d.resolved)
	return nil
}
func (d *Dynamic[T]) StartResolve() {
	d.state = resolving
}
func (d *Dynamic[T]) EndResolve() {
	d.state = resolved
}

func (d *Dynamic[T]) State(resolve func()) T {
	if d.state == unresolved {
		d.StartResolve()
		resolve()
		d.EndResolve()
	}

	return d.resolved
}
