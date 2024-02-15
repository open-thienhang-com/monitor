package base

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"mono.thienhang.com/pkg/common/logger"
)

func callAPI(token string, aType, url string, bodyData map[string]interface{}) (interface{}, time.Duration, error) {
	start := time.Now()
	var body map[string]interface{}
	body = bodyData

	data := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(data)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.Encode(body)

	req, err := http.NewRequest(aType, url, data)
	logger.Warn(req)
	if err != nil {
		logger.Error(err)
		return nil, 0, errors.New("can not make a request by abnormal reason")
	}
	// Header - API get user information
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", token)
	//
	logger.Error(token)
	//
	client := &http.Client{
		Timeout: 3 * time.Second,
	}
	resp, err := client.Do(req)
	end := time.Since(start)
	if err != nil {
		logger.Error(err)
		return nil, end, errors.New("can not send request to API ")
	}
	if resp == nil || resp.Body == nil {
		return nil, end, errors.New("can not send request to API by null body")
	}
	defer func() {
		//if resp != nil {
		resp.Body.Close()
		//}
	}()
	logger.Error(resp)
	switch resp.StatusCode {
	case http.StatusOK:
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		if result["data"] != nil {
			res := result["data"]
			return res, end, nil
		}
		return nil, end, errors.New("API return null value")
	case http.StatusUnauthorized:
		// callAPI(token, wsToken, aType, url, body)
		return nil, end, errors.New("forbidden API")
	default:
		// Parse body
		var result map[string]interface{}
		json.NewDecoder(resp.Body).Decode(&result)
		return nil, end, errors.New("abnormal error wwith api authen")
	}
}
