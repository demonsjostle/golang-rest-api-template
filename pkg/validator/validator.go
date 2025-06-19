package validator

type Validator interface {
	// Add your validator methods here
}

func New() Validator {
	return &validator{}
}

type validator struct {
	// Add your validator fields here
}
