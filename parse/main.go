package main

import (
	"fmt"
	"regexp"
	"strings"
)

var lines = `
	e.GET("/health", healthCheck)
	// admin route
	e.POST("/v1/admin/login", wrapHandler(authenticate))
	e.POST("/v1/admin/validate_token", wrapHandler(validateToken))
	e.PUT("/v1/admin/:id", wrapHandler(saveAdmin))
	e.GET("/v1/admin/:id", wrapHandler(GetAdmin))
	e.DELETE("/v1/admin/:id", wrapHandler(deleteAdmin))

	//filter route
	e.GET("/v1/filter/block_list", wrapHandler(listBlockList))
	e.POST("/v1/filter/block_list", wrapHandler(addBlockList))
	e.PUT("/v1/filter/block_list", wrapHandler(editBlockList))
	e.GET("/v1/filter/block_list/:list_id", wrapHandler(listBlockList))
	e.DELETE("/v1/filter/block_list/:list_id", wrapHandler(deleteBlockList))
	e.GET("/v1/filter/blocked_item", wrapHandler(listBlockedItem))
	e.GET("/v1/filter/blocked_item/:list_id", wrapHandler(listBlockedItem))
	e.POST("/v1/filter/blocked_item", wrapHandler(addBlockedItem))
	e.PUT("/v1/filter/blocked_item", wrapHandler(editBlockedItem))
	e.DELETE("/v1/filter/blocked_item/:list_id", wrapHandler(deleteBlockedItem))
	e.GET("/v1/filter/block_rule", wrapHandler(listBlockRules))
	e.POST("/v1/filter/block_rule", wrapHandler(addBlockRule))
	e.PUT("/v1/filter/block_rule", wrapHandler(editBlockRule))
	e.GET("/v1/filter/block_rule/:rule_id", wrapHandler(listBlockRules))
	e.DELETE("/v1/filter/block_rule/:rule_id", wrapHandler(deleteBlockRule))

	//ads route
	e.GET("/v1/ads/ad_id", wrapHandler(getAdID))
	e.GET("/v1/ads/history", wrapHandler(getAdHistory))
	e.GET("/v1/ads/load_action", wrapHandler(loadAdAction))
	e.GET("/v1/ads/load_params", wrapHandler(loadAdParams))
	e.GET("/v1/ads/get_info", wrapHandler(getAdInfo))
	e.GET("/v1/ads/get_multiple_info", wrapHandler(getMultipleAdInfo))
	e.POST("/v1/ads/change_status", wrapHandler(AdStatusChange))
	e.POST("/v1/ads/unlock", wrapHandler(Unlock))
	e.POST("/v1/ads/refuse", wrapHandler(Refuse))
	e.POST("/v1/ads/accept", wrapHandler(Accept))
	e.POST("/v1/ads/change_multiple_status", wrapHandler(changeStatusMultipleAds))
	e.POST("/v1/ads/move_between_shop_and_chotot", wrapHandler(MoveBetweenShopAndChotot))
	e.POST("/v1/ads/clear", wrapHandler(clearAd))
	e.POST("/v1/ads/link_ads_shop", wrapHandler(linkAdsShop))
	e.POST("/v1/ads/bump", wrapHandler(Bump))
	e.POST("/v1/ads/bump_ad", wrapHandler(Bump)) // backward commpatible
	e.POST("/v1/ads/bump_schedule", wrapHandler(bumpSchedule))
	e.GET("/v1/ads/get_multiple_changed_values", wrapHandler(getMultipleChangedValues))
	e.POST("/v1/ads/new", wrapHandler(NewAd))
	e.GET("/v1/ads/get_unverified_ad_ids_by_phone/:phone", wrapHandler(getUnverifiedAdIdsByPhone))
	e.GET("/v1/ads/get_last_accepted_action_id/:ad_id", wrapHandler(getLastAcceptedActionID))
	e.POST("/v1/ads/search", wrapHandler(searchAds))
	e.POST("/v1/ads/move_unpaid_ad", wrapHandler(MoveUnpaidAd))
	e.POST("/v1/ads/move_unverified_ad", wrapHandler(MoveUnverifiedAd))
	e.GET("/v1/ads/get_unverified_ad_in_queue", wrapHandler(getUnverifiedAdsInQueue))

	e.POST("/v1/queues/get", wrapHandler(getQueue)) // POST: it locks the retrieved ads
	e.GET("/v1/queues/summary", wrapHandler(getQueueSummary))

	e.POST("/v1/img", wrapHandler(ImgPut))

	e.GET("/v1/scheduler", wrapHandler(getCommand))
	e.POST("/v1/scheduler/register", wrapHandler(registerCommand))
	e.POST("/v1/scheduler/lock", wrapHandler(lockCommand))
	e.POST("/v1/scheduler/set_result", wrapHandler(setCommandResult))

	e.POST("/v1/notice", wrapHandler(setNotice))
	e.DELETE("/v1/notice", wrapHandler(delNotice))

	//misc route
	e.GET("/v1/conf", wrapHandler(getDynamicConf))
	e.GET("/v1/conf/:key", wrapHandler(getDynamicConfByKey))

	e.GET("/v1/usergroup_info", wrapHandler(getUsergroupInfo))
	e.GET("/v1/phone_by_uid", wrapHandler(getPhoneFromUID))
	e.GET("/v1/bump_info", wrapHandler(getBumpByServiceID))

	e.GET("/v1/mama/next_ad", wrapHandler(mamaGetNextAd))
	e.GET("/v1/mama/config", wrapHandler(mamaLoadConfig))
	e.GET("/v1/mama/version", wrapHandler(mamaGetVersion))
	e.POST("/v1/mama/flush_queue", wrapHandler(mamaFlushQueue))

	e.POST("/v1/admail", wrapHandler(adMail))

	e.GET("/stats", getStats)
`

func main() {
	// fmt.Println(lines)
	for _, v := range strings.Split(lines, "\n") {
		// fmt.Println(i)
		// fmt.Println(v)
		toHandlerInfo(v)
	}
}

var rex = regexp.MustCompile(`e\.(.*)\(\"(.*)\", wrapHandler\((.*)\)\)`)
var handlerInfofmt = "handlerInfo{\"%s\", \"%s\", %s},\n"

func toHandlerInfo(s string) {
	if s == "" {
		return
	}
	tmp := rex.FindStringSubmatch(strings.TrimSpace(s))
	if len(tmp) > 1 {
		fmt.Printf(handlerInfofmt, tmp[1], tmp[2], tmp[3])
	}
}
