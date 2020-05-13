package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"

	"git.chotot.org/ad-service/spine/apps/types"
	spine "git.chotot.org/ad-service/spine/client"
	"git.chotot.org/ad-service/spine/pkg/types/ad"
)

var (
	spineURL  = "http://127.0.0.1:5656"
	adIDs     = []string{"39541151"}
	shortType = map[string]string{
		"buy":  "k",
		"sell": "s",
		"rent": "h",
		"swap": "b",
		"let":  "u",
	}

	spineCli *spine.Client
)

func main() {
	args := os.Args[1:]
	if len(args) < 5 {
		fmt.Println("not enough params")
		return
	}
	spineURL := args[1]
	spineCli = spine.NewClient(nil, spine.BaseURL(spineURL))

	username := args[2]
	passwd := args[3]

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

	limit, err := strconv.Atoi(args[4])
	if err != nil {
		log.Println(err)
		return
	}

	if limit > len(records) || limit <= 0 {
		limit = len(records)
	}

	token := authenticate(username, passwd)

	limiter := time.Tick(50 * time.Millisecond)
	for i := 1; i < limit; i++ {
		<-limiter
		adID, err := strconv.Atoi(records[i][0])
		if err != nil {
			fmt.Println(err)
			continue
		}

		ad, err := loadAd(adID)
		if err != nil {
			fmt.Printf("Cannot load ad %d: %v", adID, err)
			continue
		}

		log.Println(accept(token, ad))
		log.Println(hideAd(ad))
	}
}

func authenticate(username, passwd string) string {
	req := spine.AuthenticateRequest{
		Username: username,
		Password: passwd,
	}

	res := spine.AuthenticateResponse{}
	if err := spineCli.Authenticate(req, &res); err != nil {
		log.Println("authenticate error:", err)
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

func getImages(in *types.LoadActionResponse) ([]string, error) {
	imgs := make([]string, 0, len(in.Images))
	for _, img := range in.Images {
		imgs = append(imgs, img.Name)
	}

	return imgs, nil
}

func clear(in *ad.NewAdResponse) error {
	req := types.ClearAdRequest{
		AdID:       in.AdID,
		ActionID:   in.ActionID,
		RemoteAddr: "10.50.10.129",
	}
	res := types.ClearAdResponse{}
	defer func() {
		fmt.Println(res)
	}()
	return spineCli.ClearAd(req, &res)
}

func prolong(in *types.LoadActionResponse) error {
	var price int64
	if in.Ad.Price != nil {
		price = *in.Ad.Price
	}

	category, err := strconv.Atoi(in.Ad.Category)
	if err != nil {
		return err
	}

	req := ad.NewAdRequest{
		AdID:         in.Ad.AdID,
		RemoteAddr:   "10.50.10.129",
		Price:        int(price),
		Subject:      in.Ad.Subject,
		Body:         *in.Ad.Body,
		Lang:         *in.Ad.Lang,
		AdType:       "deleted",
		Region:       in.Ad.Region,
		Category:     category,
		Type:         shortType[in.Ad.Type],
		Phone:        in.Ad.Phone,
		AccountID:    *in.Ad.AccountID,
		CompanyAd:    in.Ad.CompanyAd,
		Name:         in.Ad.Name,
		NeedsPayment: false,
	}

	params, err := getAdParams(in)
	if err != nil {
		return err
	}
	req.Params = *params

	if in.Ad.ListID != nil {
		req.ListID = *in.Ad.ListID
	}

	if in.User.Email != nil {
		req.Email = *in.User.Email
	}

	if in.Ad.ShopAlias != nil {
		req.ShopAlias = *in.Ad.ShopAlias
	}

	imgs, err := getImages(in)
	if err != nil {
		return err
	}
	req.Image = imgs

	// fmt.Println(*req.Params.SellerAddr)

	fmt.Println(req)

	res := ad.NewAdResponse{}
	err = spineCli.NewAd(req, &res)
	fmt.Println(res, err)

	err = clear(&res)
	return err
}

func getLastAcceptedID(adID int) int {
	req := types.GetLastAcceptedActionIDRequest{
		AdID: adID,
	}
	res := types.GetLastAcceptedActionIDResponse{}
	if err := spineCli.GetLastAcceptedActionID(req, &res); err != nil {
		log.Println("Cannot get last actionid: ", err)
		return 0
	}
	return *res.ActionID
}

func accept(token string, ad *types.LoadActionResponse) error {
	actionID := getLastAcceptedID(ad.Ad.AdID)
	if actionID == 0 {
		return fmt.Errorf("%d has no action_id", ad.Ad.AdID)
	}

	doNotSendEmail := true

	category := 0
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
		Category:      &category,
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

func hideAd(ad *types.LoadActionResponse) error {
	if ad.Ad.Status != "hidden" {
		return nil
	}

	if ad.Ad.ListID == nil {
		return nil
	}

	req := types.AdStatusChangeRequest{
		ListID:       uint32(*ad.Ad.ListID),
		Status:       "hidden",
		Source:       "desktop_site_flashad",
		HiddenReason: "2",
	}
	res := types.AdStatusChangeResponse{}
	return spineCli.ChangeAdStatus(req, &res)
}

func formatArea(params *ad.Params) {
	area := 0
	areaV2Ptr := params.AreaV2
	if areaV2Ptr != nil && *areaV2Ptr > 0 {
		area = *areaV2Ptr % 1000
	}

	if area > 0 {
		params.Area = &area
	}
}
