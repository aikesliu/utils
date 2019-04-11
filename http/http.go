package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"utils/logger"
)

func Req2Json(r *http.Request, req interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("read body failed, err: %v", err)
	}

	jsonStr := string(body)
	if err := json.Unmarshal([]byte(jsonStr), req); err != nil {
		return fmt.Errorf("body: %v is not json, err: %v", string(body), err)
	}
	return nil
}

func ResponseJson(w http.ResponseWriter, data interface{}, err error) {
	resp := make(map[string]interface{})
	resp["ret"] = 0
	if err != nil {
		resp["ret"] = 1
		resp["msg"] = err.Error()
	}
	resp["data"] = data
	bytesData, err := json.Marshal(resp)
	if err != nil {
		logger.E("ResponseJson json.Marshal error: %v", err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write(bytesData)
	}
}

func JsonPost(url string, jsonReq interface{}) (int, map[string]interface{}) {
	var dat map[string]interface{}
	client := &http.Client{}
	data, err := json.Marshal(jsonReq)
	if err != nil {
		logger.E("JsonPost, json req marshal failed, err: %v", err)
		return http.StatusExpectationFailed, dat
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		logger.E("JsonPost, new request failed, err: %v", err)
		return http.StatusExpectationFailed, dat
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		logger.E("JsonPost, new request do failed, err: %v", err)
		return http.StatusExpectationFailed, dat
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.E("JsonPost, new request failed, err: %v", err)
		return http.StatusExpectationFailed, dat
	}
	if err := json.Unmarshal(body, &dat); err != nil {
		logger.E("JsonPost, body: %v to json failed, err: %v", string(body), err)
	}
	return resp.StatusCode, dat
}
