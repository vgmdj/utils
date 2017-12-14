package navigation

import (
	"fmt"
	"log"
	"strconv"
	"util/dispatch"
	"util/navigation/area"
)

const (
	/*
		txSearchReq 腾讯地点搜索服务
		必填
		KeyWord  string //POI搜索关键字，用于全文检索字段
		Boundary string //搜索地理范围
		Key      string // key 值

		选填
		Filter    string  //筛选条件： filter=category=公交站
		OrderBy   string  //排序方式 例1：orderby=_distance desc
		PageSize  float32 //每页条目数，最大限制为20条	page_size=10
		PageIndex float32 //第x页，默认第1页	page_index=2
		Output    string  //返回格式：支持JSON/JSONP，默认JSON	output=json
		CallBack  string  //JSONP方式回调函数
	*/

	invalidParameter = 310
	invalidKey       = 311
	invalidReq       = 306
	unauthorizedSrc  = 110
	OK               = 0

	POITypeNormal               = 0
	POITypeBusStation           = 1
	POITypeMetro                = 2
	POITypeBusRoute             = 3
	POITypeAdministrativeRegion = 4

	txSearchURL    = "http://apis.map.qq.com/ws/place/v1/search"
	txGeocoderURL  = "http://apis.map.qq.com/ws/geocoder/v1/"
	txDistanceURL  = "http://apis.map.qq.com/ws/distance/v1/"
	txDirectionURL = "http://apis.map.qq.com/ws/direction/v1/driving/"

	radius  = 30000
	keyword = "加油站"
	key     = "MJPBZ-GTLKO-26SW6-STDXP-IQH4H-UVBDB"
	filter  = "category=加油站,中石化"
	//pageSize = "10"
)

var (
	ErrMsg map[int]string
)

type (
	//txSearchResp 腾讯地点搜索返回信息
	txSearchResp struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
		Count   int    `json:"count"`
		Data    []poi  `json:"data"`
	}

	poi struct {
		ID      string `json:"id"`
		Title   string `json:"title"`
		Address string `json:"address"`
		//Tel      string   `json:"tel"`
		//Category string   `json:"category"`
		//Type     int      `json:"type"`
		Location location `json:"location"`
		Distance float64  `json:"_distance"`
		AdInfo   adInfo   `json:"ad_info"`
	}

	location struct {
		Lat float64 `json:"lat"`
		Lng float64 `json"lng"`
	}

	adInfo struct {
		AdCode   int    `json:"adcode"`
		Province string `json:"province"`
	}

	txGeocoderResp struct {
		Status  int            `json:"status"`
		Message string         `json:"message"`
		Result  geoCoderResult `json:"result"`
	}

	geoCoderResult struct {
		AddrComponent addrComponent `json:"address_component"`
	}

	addrComponent struct {
		Nation   string `json:"nation"`
		Province string `json:"province"`
		City     string `json:"city"`
	}
)

func init() {
	ErrMsg = make(map[int]string)
	ErrMsg[invalidParameter] = "请求参数信息有误"
	ErrMsg[invalidKey] = "Key格式错误"
	ErrMsg[invalidReq] = "请求有护持信息请检查字符串"
	ErrMsg[unauthorizedSrc] = "请求来源未被授权"

}

//TxSearch POI搜索服务
func TxSearch(lat float64, lng float64, pageIndex string, pageSize string) (searchInfo txSearchResp, err error) {
	qmap := make(map[string]string)
	qmap["boundary"] = fmt.Sprintf("nearby(%f,%f,%d)", lat, lng, radius)
	qmap["keyword"] = keyword
	qmap["key"] = key
	qmap["filter"] = filter
	qmap["orderby"] = "_distance"
	qmap["output"] = "json"
	qmap["page_size"] = pageSize
	qmap["page_index"] = pageIndex

	query := dispatch.Query{
		QMap:      qmap,
		IsOrdered: false,
	}

	if err = dispatch.Get(true, txSearchURL, &searchInfo, query); err != nil {
		log.Println(txSearchURL, query)
		return
	}

	for k, v := range searchInfo.Data {
		adcode := v.AdInfo.AdCode / 10000
		adcode = adcode * 10000
		searchInfo.Data[k].AdInfo.Province = area.GetAreaNameByCode(strconv.Itoa(adcode))
	}

	return
}

func TxGeocoder(lat float64, lng float64) (geocoder txGeocoderResp, err error) {
	query := make(map[string]string)
	query["location"] = fmt.Sprintf("%f,%f", lat, lng)
	query["key"] = key
	query["output"] = "json"

	disQuery := dispatch.Query{
		QMap:      query,
		IsOrdered: false,
	}

	if err = dispatch.Get(true, txGeocoderURL, &geocoder, disQuery); err != nil {
		log.Println(txGeocoderURL, query)
		return
	}

	return
}
