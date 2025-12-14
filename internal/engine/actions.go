package engine

import (
	"fmt"

	"workflow_engine/internal/actions/httpcall"
	"workflow_engine/internal/actions/transform"
)

func ExecuteAction(
	actionType string,
	config map[string]interface{},
	input map[string]interface{},
) (map[string]interface{}, error) {

	switch actionType {
	case "transform":
		return transform.Run(config, input)
	case "http_call":
		return httpcall.Run(config, input)
	default:
		return nil, fmt.Errorf("unknown action type: %s", actionType)
	}
}
