package coverage

import (
	"encoding/json"
	"fmt"
)

type SetupError struct {
	Err string
}

func (e SetupError) Error() string {
	return fmt.Sprintf("error on coverage-proxy setup state: %s", e.Err)
}

func (e SetupError) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Message string `json:"message,omitempty"`
	}{e.Error()})
}
