package corews

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func JSON(v any) ([]byte, error) {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(v)
	if err != nil {
		return nil, fmt.Errorf("encode JSON: %w", err)
	}

	return buf.Bytes(), nil
}
