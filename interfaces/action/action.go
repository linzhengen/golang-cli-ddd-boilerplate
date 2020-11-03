package action

// Action is s interface
type Action interface {
	HelloWorldAction
}

type action struct {
	HelloWorldAction
}

// NewAction is Action constructor.
func NewAction(
	hw HelloWorldAction,
) Action {
	return &action{
		HelloWorldAction: hw,
	}
}
