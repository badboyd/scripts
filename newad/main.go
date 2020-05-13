package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	// "image"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/mitchellh/mapstructure"

	"git.chotot.org/ad-service/spine/apps/types"
	spine "git.chotot.org/ad-service/spine/client"
	"git.chotot.org/ad-service/spine/pkg/types/ad"
)

type Token struct {
	Token     string `json:"token"`
	AccountID int    `json:"account_id"`
	Email     string `json:"email"`
}

type newAdResponse struct {
	AdID int `json:"ad_id"`
}

var httpClient = &http.Client{
	Timeout: 30 * time.Second,
}

var (
	spineBaseURL     string
	brainstemBaseURL string
	loginBaseURL     string

	spineCli *spine.Client

	shortType = map[string]string{
		"buy":  "k",
		"sell": "s",
		"rent": "h",
		"swap": "b",
		"let":  "u",
	}
)

func newAd(in map[string]interface{}) (*ad.NewAdResponse, error) {
	if in == nil {
		return nil, fmt.Errorf("NIL_INPUT")
	}
	newAdReq := ad.NewAdRequest{}
	decodeConfig := mapstructure.DecoderConfig{
		WeaklyTypedInput: true,
		TagName:          "json",
		Result:           &newAdReq,
	}
	decoder, err := mapstructure.NewDecoder(&decodeConfig)
	if err != nil {
		return nil, err
	}
	if err = decoder.Decode(in); err != nil {
		return nil, err
	}

	decodeConfig.Result = &newAdReq.Params
	decoder, err = mapstructure.NewDecoder(&decodeConfig)
	if err != nil {
		return nil, err
	}
	if err = decoder.Decode(in); err != nil {
		return nil, err
	}
	formatArea(&newAdReq)
	formatRegion(&newAdReq)

	newAdReq.Type = "s"

	newAdRes := ad.NewAdResponse{}
	// fmt.Printf("%+v\n", newAdReq)
	if err := spineCli.NewAd(newAdReq, &newAdRes); err != nil {
		return nil, err
	}
	return &newAdRes, nil
}

var (
	inputLinkPattern    = "https://drive.google.com/open?id=%s"
	downloadLinkPattern = "https://drive.google.com/uc?export=download&id=%s"
)

func getDownloadLink(link string) string {
	id := ""
	if _, err := fmt.Sscanf(link, inputLinkPattern, &id); err != nil {
		fmt.Println(err)
	}
	return fmt.Sprintf(downloadLinkPattern, id)
}

func uploadImage(imageURL string) (url string, err error) {
	var output types.ImgPutResponse
	resp, err := http.Get(getDownloadLink(imageURL))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	err = spineCli.ImgPut(&types.ImgPutRequest{
		Image:     body,
		Thumbnail: body,
	}, &output)
	return output.ImageName, err
}

func main() {
	args := os.Args[1:]
	if len(args) < 5 {
		fmt.Println("not enough params")
		return
	}
	spineBaseURL = args[1]
	spineCli = spine.NewClient(nil, spine.BaseURL(spineBaseURL))
	loginBaseURL = args[2]

	username := args[3]
	passwd := args[4]

	fileName := args[0]
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	r := csv.NewReader(file)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		return
	}

	limit, err := strconv.Atoi(args[5])
	if err != nil {
		fmt.Println(err)
		return
	}

	if limit > len(records) || limit <= 0 {
		limit = len(records)
	}

	token := authenticate(username, passwd)

	params := records[0]

	cnt := 0

	limiter := time.Tick(50 * time.Millisecond)
	for i := 1; i < limit; i++ {
		<-limiter
		// fmt.Println(records[i])
		if err := do(params, records[i], token); err != nil {
			fmt.Println(err)
		} else {
			cnt++
		}
	}

	fmt.Println(cnt, "/", len(records))
}

func formatImgRequest(in map[string]interface{}) {
	imgs := []string{}
	for _, v := range strings.Split(fmt.Sprint(in["image"]), "\n") {
		// fmt.Println(v)
		if v == "" {
			continue
		}

		res, err := uploadImage(v)
		if err != nil {
			fmt.Println(err)
			continue
		}
		imgs = append(imgs, res)
	}
	in["image"] = imgs
}

func login(username string, pw string, loginUrl string) (*Token, error) {
	values := map[string]string{"phone": username, "password": pw}
	jsonValue, _ := json.Marshal(values)
	resp, err := http.Post(loginUrl, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var s Token
	json.Unmarshal(body, &s)
	return &s, err
}

func addAccountID(in map[string]interface{}) {
	phone := fmt.Sprint(in["phone"])
	passwd := fmt.Sprint(in["passwd"])
	tkn, err := login(phone, passwd, loginBaseURL)
	if err != nil {
		return
	}
	in["account_id"] = tkn.AccountID
	in["email"] = "trucnguyen.117@gmail.com"
	in["lang"] = "vi"
}

func do(names, values []string, token string) error {
	in := make(map[string]interface{})
	for i, v := range names {
		in[v] = values[i]
	}

	addAccountID(in)
	// newad
	formatImgRequest(in)
	newAdRes, err := newAd(in)
	if err != nil {
		return err
	}
	// fmt.Printf("newAdRes: %+v\n", newAdRes)
	// clear ad
	err = clear(newAdRes)
	if err != nil {
		return err
	}
	// acceptad
	loadAdRes, err := loadAd(newAdRes.AdID)
	if err != nil {
		return err
	}

	if err := accept(token, loadAdRes); err != nil {
		return err
	}

	fmt.Println("ad_id: ", newAdRes.AdID)
	return nil
}

func authenticate(username, passwd string) string {
	req := spine.AuthenticateRequest{
		Username: username,
		Password: passwd,
	}

	res := spine.AuthenticateResponse{}
	if err := spineCli.Authenticate(req, &res); err != nil {
		panic(err)
	}

	return res.Token
}

func loadAd(adID int) (*types.LoadActionResponse, error) {
	req := types.LoadActionRequest{AdID: adID}
	res := types.LoadActionResponse{}
	err := spineCli.LoadAction(req, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func clear(in *ad.NewAdResponse) error {
	req := types.ClearAdRequest{
		AdID:       in.AdID,
		ActionID:   in.ActionID,
		RemoteAddr: "10.50.10.129",
	}
	res := types.ClearAdResponse{}
	// defer func() {
	// fmt.Println(res)
	// }()
	return spineCli.ClearAd(req, &res)
}

func accept(token string, ad *types.LoadActionResponse) error {
	actionID := 1
	doNotSendEmail := true

	category, err := strconv.Atoi(ad.Ad.Category)
	if err != nil {
		return err
	}

	hasImages := true

	adType := shortType[ad.Ad.Type]

	req := types.AcceptRequest{
		Token:         token,
		AdID:          ad.Ad.AdID,
		ActionID:      actionID,
		Subject:       &ad.Ad.Subject,
		Body:          ad.Ad.Body,
		RemoteAddr:    "10.50.10.129",
		CompanyAd:     &ad.Ad.CompanyAd,
		AdType:        &adType,
		DoNotSendMail: &doNotSendEmail,
		RemoteBrowser: "spine",
		Category:      category,
		Region:        &ad.Ad.Region,
		AdHasImages:   &hasImages,
		Lang:          ad.Ad.Lang,
	}

	var tmp int64
	// if ad.Ad.Price != nil {
	// 	tmp = int64(*ad.Ad.Price)
	if ad.Ad.OrigPrice != nil {
		tmp, err = strconv.ParseInt(*ad.Ad.OrigPrice, 10, 64)
		if err != nil {
			return err
		}
	}
	req.Price = &tmp

	params, err := getAdParams(ad)
	if err != nil {
		return err
	}
	convertBrandModedForVehicle(params)
	req.AdParams = *params

	res := types.AcceptResponse{}
	return spineCli.ReviewAccept(req, &res)
}

func convertBrandModedForVehicle(params *ad.Params) {
	if params.MotorbikeBrand != nil && params.MotorbikeModel != nil {
		switch {
		case *params.MotorbikeBrand == 4 && *params.MotorbikeModel == 69:
			newBrand := 5
			params.MotorbikeBrand = &newBrand
		case *params.MotorbikeBrand == 30 && *params.MotorbikeModel == 243:
			newModel := 242
			params.MotorbikeModel = &newModel
		case *params.MotorbikeBrand == 2 && *params.MotorbikeModel == 352:
			newModel := 281
			params.MotorbikeModel = &newModel
		}
	}

}

func getAdParams(in *types.LoadActionResponse) (*ad.Params, error) {
	if in.Params == nil {
		return nil, nil
	}

	tmp := make(map[string]interface{})
	for _, param := range in.Params {
		if param.Value != nil {
			tmp[param.Name] = *param.Value
		} else {
			tmp[param.Name] = nil
		}
	}

	params := ad.Params{}
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "json",
		WeaklyTypedInput: true,
		Result:           &params,
	})
	if err != nil {
		return nil, err
	}

	err = dec.Decode(tmp)
	return &params, err
}

func formatArea(ad *ad.NewAdRequest) {
	area := 0
	areaV2Ptr := ad.Params.AreaV2
	if areaV2Ptr != nil && *areaV2Ptr > 0 {
		area = *areaV2Ptr % 1000
	}

	if area > 0 {
		ad.Params.Area = area
	}
}

func formatRegion(ad *ad.NewAdRequest) {
	regionV2 := ad.Params.RegionV2
	region := regionV2 / 1000
	if region > 0 {
		ad.Region = region
	}
}
