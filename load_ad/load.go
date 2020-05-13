package main

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/structs"

	"git.chotot.org/ad-service/spine/apps/types"
	"git.chotot.org/ad-service/spine/client"
)

type adMap map[string]interface{}

func (am adMap) adParam(name string) string {
	if am == nil {
		return ""
	}

	params, ok := am["Req"].(map[string]interface{})["Params"]
	if !ok || params == nil {
		return ""
	}

	for _, v := range params.([]interface{}) {
		param, ok := v.(map[string]interface{})
		if !ok {
			continue
		}
		if param["Name"].(string) == name {
			return fmt.Sprint(param["Value"])
		}
	}
	return ""
}

func formatAdParams(i interface{}) {
	tmp, ok := i.(map[string]interface{})
	if !ok {
		return
	}

	paramsInf, ok := tmp["Params"]
	if !ok {
		return
	}

	paramsArr, ok := paramsInf.([]interface{})
	if !ok {
		return
	}

	paramsMap := make(map[string]interface{})
	for _, paramInf := range paramsArr {
		paramMap, ok := paramInf.(map[string]interface{})
		if !ok {
			continue
		}
		paramsMap[paramMap["Name"].(string)] = paramMap["Value"]
	}
	delete(tmp, "Params")
	tmp["Params"] = paramsMap
}

func formatActionParams(i interface{}) {
	tmp, ok := i.(map[string]interface{})
	if !ok {
		return
	}

	paramsInf, ok := tmp["ActionParams"]
	if !ok {
		return
	}

	paramsArr, ok := paramsInf.([]interface{})
	if !ok {
		return
	}

	paramsMap := make(map[string]interface{})
	for _, paramInf := range paramsArr {
		paramMap, ok := paramInf.(map[string]interface{})
		if !ok {
			continue
		}
		paramsMap[paramMap["Name"].(string)] = paramMap["Value"]
	}
	delete(tmp, "ActionParams")
	tmp["ActionParams"] = paramsMap
}

func formatUserParams(i interface{}) {
	tmp, ok := i.(map[string]interface{})
	if !ok {
		return
	}

	paramsInf, ok := tmp["UserParams"]
	if !ok {
		return
	}

	paramsArr, ok := paramsInf.([]interface{})
	if !ok {
		return
	}

	paramsMap := make(map[string]interface{})
	for _, paramInf := range paramsArr {
		paramMap, ok := paramInf.(map[string]interface{})
		if !ok {
			continue
		}
		paramsMap[paramMap["Name"].(string)] = paramMap["Value"]
	}
	delete(tmp, "UserParams")
	tmp["UserParams"] = paramsMap
}

func main() {
	spineClient := client.NewClient(nil,
		client.BaseURL("http://10.60.3.47:5656"),
		client.UserAgent("badboyd"),
	)

	loadAdReq := types.LoadActionRequest{
		AdID:             34888626,
		ActionID:         1,
		GetAssociatedAds: true,
	}
	loadAdRes := types.LoadActionResponse{}
	if err := spineClient.LoadAction(loadAdReq, &loadAdRes); err != nil {
		panic(err)
	}

	loadAdMap := structs.Map(loadAdRes)
	formatAdParams(loadAdMap)
	formatUserParams(loadAdMap)
	formatActionParams(loadAdMap)
	loadAdMap = map[string]interface{}{
		"Req": loadAdMap,
	}
	tmp, err := json.MarshalIndent(loadAdMap, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(tmp))
}
