package twins

import (
	"encoding/json"
	"fmt"
	"github.com/Leo-Guo/twins/service/compare"
	"testing"
)

type testReq struct {
	StoreId     int   `json:"store_id"`
	Class1Id    int   `json:"class1_id"`
	SortType    int   `json:"sort_type"`
	FilterTypes []int `json:"filter_types"`
	Page        int   `json:"page"`
	PageSize    int   `json:"page_size"`
}

type testRet struct {
	Ret     int         `json:"ret"`
	Date    interface{} `json:"date"`
	NowTime int64       `json:"now_time"`
}


func TestRequestDiff(t *testing.T){
	reqA := &testReq{StoreId: 63503, Class1Id: 2011, SortType: 1, FilterTypes: []int{1, 2}, Page: 1, PageSize: 50}
	res := &testRet{}
	_, responseA := SendRequest("http://localhost:8234", "/app/list/get-product-list-by-class-and-filter", "POST", reqA, res)
	reqB := &testReq{StoreId: 63503, Class1Id: 2011, SortType: 1, FilterTypes: []int{1, 2}, Page: 1, PageSize: 50}
	_, responseB := SendRequest("http://127.0.0.1:9010", "/app/list/get-product-list-by-class-and-filter", "POST", reqB, res)

	var (
		json1 map[string]interface{}
		json2 map[string]interface{}
	)

	errA := json.Unmarshal(responseA, &json1)
	errB := json.Unmarshal(responseB, &json2)
	if errA != nil || errB != nil{
		fmt.Println(errA)
		fmt.Println(errB)
	}

	retString,isDiff := compare.JsonCompareDiff(json1,json2,1)

	println(retString,isDiff)
}

