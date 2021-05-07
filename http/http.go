package http

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/aikesliu/utils/marshal"
)

// read request body to struct as json
func Req2Json(r *http.Request, req interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return marshal.JsonStr2Struct(body, req)
}

// json resp
func JResp(w http.ResponseWriter, v interface{}, err error) {
	resp := make(map[string]interface{})
	resp["ret"] = 0
	if err != nil {
		resp["ret"] = 1
		resp["msg"] = err.Error()
	}
	resp["data"] = v
	str, _ := marshal.Struct2JsonStr(resp)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(str))
	w.Write([]byte("\n"))
}

// json post
func JPost(url string, req interface{}) (error, map[string]interface{}) {
	client := &http.Client{}
	reqStr, _ := marshal.Struct2JsonStr(req)
	httpReq, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(reqStr)))
	if err != nil {
		return err, nil
	}
	httpReq.Header.Add("Content-Type", "application/json")
	resp, errDo := client.Do(httpReq)
	if errDo != nil {
		return err, nil
	}
	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return errRead, nil
	}
	data := make(map[string]interface{})
	errJson := marshal.JsonStr2Struct(body, &data)
	return errJson, data
}
