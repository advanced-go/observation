package timeseries1

import (
	"fmt"
	"github.com/advanced-go/stdlib/json"
)

func ExampleEntry() {
	buf, status := json.Marshal(entryData)
	if !status.OK() {
		fmt.Printf("test: Entry{} -> [status:%v]\n", status)
	} else {
		fmt.Printf("test: Entry{} -> %v\n", string(buf))
	}

	//Output:
	//fail

}

func _ExampleScanColumnsTemplate() {
	//log := scanColumnsTemplate[AccessLog](nil)

	//fmt.Printf("test: scanColumnsTemplate[AccessLog](nil) -> %v\n", log)

	//Output:
	//fail
}

func _ExampleScannerInterface_V1() {

	//log, status := scanRowsTemplateV1[AccessLog, AccessLog](nil)
	//fmt.Printf("test: scanRowsTemplateV1() -> [status:%v] [elem:%v] [log:%v] \n", status, reflect.TypeOf(log).Elem(), log[0].DurationString)

	//Output:
	//test: scanRowsTemplateV1() -> [status:OK] [elem:timeseries.AccessLog] [log:SCAN() TEST DURATION STRING]

}

func _ExampleScannerInterface() {
	//log, status := scanRowsTemplate[accessLogV2](nil)

	//log, status := scanRowsTemplate[AccessLog](nil)
	//fmt.Printf("test: scanRowsTemplate() -> [status:%v] [elem:%v] [log:%v] \n", status, reflect.TypeOf(log).Elem(), log[0].DurationString)

	//Output:
	//test: scanRowsTemplateV1() -> [status:OK] [elem:timeseries.AccessLog] [log:SCAN() TEST DURATION STRING]

}
