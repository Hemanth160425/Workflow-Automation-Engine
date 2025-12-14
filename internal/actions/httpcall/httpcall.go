package httpcall

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func Run(
	config map[string]interface{},
	input map[string]interface{},
) (map[string]interface{}, error) {

	url := config["url"].(string)
	method := config["method"].(string)

	bodyTemplate, _ := config["body_template"].(map[string]interface{})
	body := make(map[string]interface{})

	for k, v := range bodyTemplate {
		s := v.(string)
		if strings.HasPrefix(s, "{{") && strings.HasSuffix(s, "}}") {
			key := strings.Trim(s, "{}")
			body[k] = input[key]
		} else {
			body[k] = v
		}
	}

	bodyBytes, _ := json.Marshal(body)

	req, err := http.NewRequest(method, url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var output map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&output)

	return output, nil
}
