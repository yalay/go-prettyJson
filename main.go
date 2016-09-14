package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	oriJson := []byte(`{"id":1000,"name":"牛尔sohu白玉兰系列8.8","description":"","campaignId":28,"defaultPrice":3,"dayBudget":10000000,"impressionSetting":{},"maxAdImpression":6,"online":false,"comment":"","audienceType":2,"bidType":"kCpc","imPeriodStrategy":{"allowBits":[16777215,16777215,16777215,16777215,16777215,16777215,16777215],"periodDiscount":null},"categories":null,"accountId":28,"startTime":1375891200,"endTime":1423065540,"adTags":null,"deepRetargetingId":0,"retargetingId":0,"tagId":null,"sids":null,"bidFlag":"kNoNormalBid","adGender":"kAll","netType":null,"deviceType":0,"clickUrl":"","destUrl":"","monitorUrl":"","appid":"","mediaPackage":null,"tkCampId":"","gaMultiFlag":false,"trafficSetting":0,"promoteType":"kWeb","maxAdClick":1,"monitorUrl2":"","intervalType":"day","bidPrice":3,"isAgent":false,"dailyClickLimit":0,"dailyImpressionLimit":0,"adgFlag":4,"contractCode":"","clickMonitorUrls":null,"zyzTransUrls":null,"platform":"performance","downloadTrackList":null,"pkgname":"","dimCarriers":null,"appname":"","maxWholeAdImpression":0,"maxWholeAdClick":0,"status":410,"filterAppOldUser":false}`)
	var buf bytes.Buffer
	err := json.Indent(&buf, oriJson, "", "    ")
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Println(buf.String())
	return
}
