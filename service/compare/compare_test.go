package compare

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestCompare(t *testing.T)  {
	strA :=  []byte("{\"store_id\":63503,\"class_1_id\":2011,\"sort_type\":1,\"filter_types\":[1,2],\"page\":1,\"page_size\":50}")
	strB := []byte("{\n    \"store_id\":63503,\n    \"class1_id\":2011,\n    \"sort_type\":1, \n    \"filter_types\":[1,2], \n    \"page\":1,\n    \"page_size\":50\n}")

	var (
		json1 map[string]interface{}
		json2 map[string]interface{}
	)

	errA := json.Unmarshal(strA, &json1)
	errB := json.Unmarshal(strB, &json2)
	if errA != nil || errB != nil{
		fmt.Println(errA)
		fmt.Println(errB)
	}

	retString,isDiff := JsonCompareDiff(json1,json2,1)
	println(retString,isDiff)
}