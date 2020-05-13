package main

import (
	"encoding/json"
	"fmt"
	"time"

	// "github.com/fatih/structs"

	"git.chotot.org/data/kstream"

	"git.chotot.org/go-kafka-consumer/logger"

	// "git.chotot.org/ad-service/spine/apps/types"
	// spine "git.chotot.org/ad-service/spine/client"

	"github.com/antonmedv/expr"

	"git.chotot.org/gandalf/wand"
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
	var tmp wand.VerdictValue
	fmt.Printf("%+v", tmp)

	op := kstream.SimpleOp{
		Name: "expr",
		Op: func(i *kstream.Record) ([]*kstream.Record, error) {
			if i.Val == nil {
				return nil, nil
			}
			normalizeMap := i.Val.(map[string]interface{})
			formatAdParams(normalizeMap)
			formatActionParams(normalizeMap)
			// fmt.Println("normalizedMapParams: ", normalizeMap["Params"])
			// fmt.Println("normalizedMapAd: ", normalizeMap["Ad"])
			// fmt.Println("normalizedMapNads: ", normalizeMap["Nads"])
			// fmt.Println("normalizedMapNadsv2: ", normalizeMap["NadsV2"])
			// fmt.Println("normalizedMapActionParams: ", normalizeMap["ActionParams"])

			data := adMap{
				"Req": normalizeMap,
			}
			dataString, err := json.MarshalIndent(data, "", "   ")
			if err != nil {
				panic(err)
			}
			fmt.Println(string(dataString))

			ruleExpr, err := expr.Parse(`Req.Ad.Category not in [5070,5060,5030,5010]`)
			if err != nil {
				panic(err)
			}

			fmt.Println(ruleExpr.Eval(data))
			timer := time.NewTimer(1 * time.Second)
			<-timer.C
			return nil, nil
		},
	}

	kconf := kstream.ProcessConfig{
		Brokers:         []string{"10.60.3.49:9092", "10.60.3.50:9092", "10.60.3.51:9092"},
		Registry:        "http://10.60.3.50:8081",
		Subscribe:       []string{"robocop.papa_reviewer.ad"},
		AppID:           "dat18",
		ProcessName:     "test",
		ISerder:         kstream.AVRO,
		Log:             logger.GetLogger("dat-test"),
		RewindEvent:     true,
		NormalizeRecord: true,
	}

	kp, err := kstream.NewProcessor(kconf, &op)
	if err != nil {
		panic(err)
	}

	err = kp.Init(nil)
	if err != nil {
		panic(err)
	}

	go func() {
		timer := time.NewTimer(1 * time.Minute)
		<-timer.C
		kp.Stop()
	}()

	kp.Run()

	// data["AdParam"] = data.adParam

	// ruleExpr, err := expr.Parse(`Req.Ad.Phonee == "01655528881" and Ad.Category == "12020"`)
	// ruleExpr, err := expr.Parse(`Ad.Category in [5030,5070] and NadsV2.UnreviewedPhone >= 10`)

}
