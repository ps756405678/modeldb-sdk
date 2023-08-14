package entry

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/ps756405678/modeldb-sdk/domain"
)

const (
	getModelFunUrlApi  = "http://192.168.0.68:8082/api/model/func/url"
	getModelHookUrlApi = "http://192.168.0.68:8082/api/model/func/hook/url"
	emitEventApi       = "http://192.168.0.68:8082/api/model/func/url"
)

func GetModelFuncUrl(modelId string, methodName string) (url string, err error) {
	var callFuncReq = map[string]any{
		"model_db_id": modelId,
		"func_name":   methodName,
	}
	url, err = getUrl(callFuncReq, getModelFunUrlApi)
	return
}

func GetModelHookUrl(modelId string, hookName string) (url string, err error) {
	var hookUrlReq = map[string]any{
		"model_db_id": modelId,
		"hook_name":   hookName,
	}

	url, err = getUrl(hookUrlReq, getModelHookUrlApi)
	return
}

func EmitEvent(httpReq *http.Request, eventName string, params any) (err error) {
	var emitReq = domain.EmitEventReq{
		ApplicationId: httpReq.Header.Get("Application-Id"),
		InstanceId:    httpReq.Header.Get("Instance-Id"),
		ModelDBId:     httpReq.Header.Get("Model-Id"),
		EventName:     eventName,
		Params:        params,
	}
	bData, err := json.Marshal(emitReq)
	if err != nil {
		return
	}

	request, err := http.NewRequest(http.MethodPost, emitEventApi, bytes.NewReader(bData))
	if err != nil {
		return
	}

	var client = http.Client{}

	httpResp, err := client.Do(request)
	if err != nil {
		return
	}

	buff := make([]byte, httpResp.ContentLength)
	httpResp.Body.Read(buff)

	var result map[string]any
	err = json.Unmarshal(buff, &result)
	if err != nil {
		return
	}

	if result["errcode"].(int) != 0 {
		err = errors.New(result["msg"].(string))
		return
	}
	return
}

func getUrl(req any, api string) (url string, err error) {
	bData, err := json.Marshal(req)
	if err != nil {
		return
	}

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(bData))
	if err != nil {
		return
	}

	var client = http.Client{}

	httpResp, err := client.Do(request)
	if err != nil {
		return
	}

	buff := make([]byte, httpResp.ContentLength)
	httpResp.Body.Read(buff)

	var result map[string]any
	err = json.Unmarshal(buff, &result)
	if err != nil {
		return
	}

	if result["errcode"].(int) != 0 {
		err = errors.New(result["msg"].(string))
		return
	}

	url = "http://" + result["data"].(string)
	return
}
