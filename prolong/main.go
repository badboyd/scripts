package main

import (
	"fmt"
	"os"
	"strconv"

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
	if len(args) >= 1 {
		adIDs = args
	}

	spineCli = spine.NewClient(nil, spine.BaseURL(spineURL))

	for _, adIDStr := range adIDs {
		adID, err := strconv.Atoi(adIDStr)
		if err != nil {
			fmt.Println(err)
			continue
		}

		ad, err := loadAd(adID)
		if err != nil {
			fmt.Printf("Cannot load ad %d: %v", adID, err)
			continue
		}

		fmt.Println(prolong(ad))
	}
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
	price := 0
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
		Price:        price,
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
