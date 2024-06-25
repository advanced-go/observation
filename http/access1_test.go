package http

import (
	"github.com/advanced-go/observation/access1"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/httpx"
	"net/http"
)

const (
	entriesJson = "file:///c:/Users/markb/GitHub/observation/access1/access1test/access-v1.json"
)

var (
	content            = httpx.NewListContent[access1.Entry, struct{}, struct{}](false, matchEntry, nil, nil)
	resource           = httpx.NewResource[access1.Entry, struct{}, struct{}](module.TimeseriesAccessResourceV1, content, nil)
	authority, hostErr = httpx.NewHost(module.TimeseriesAuthority, mapResource, resource.Do)
)

/*
func initializeDocuments() {
	defer controller.DisableLogging(true)()
	if hostErr != nil {
		fmt.Printf("error: new resource %v", hostErr)
	}
	//entries, status := json.New[[]access1.Entry](entriesJson, nil)
	//if !status.OK() {
	//	fmt.Printf("initializeDocuments.New() -> [status:%v]\n", status)
	//	return
	//}
	cfg, ok := module.GetRoute(module.TimeseriesRouteName)
	if !ok {
		fmt.Printf("initializeDocuments.GetRoute() [ok:%v]\n", ok)
	}
	ctrl := controller.New(cfg, authority.Do)
	controller.RegisterController(ctrl)
	//_, status = put[core.Output](context.Background(), nil, entries)
	//if !status.OK() {
	//	fmt.Printf("initializeDocuments.put() [status:%v]\n", status)
	//}

}


*/

func matchEntry(req *http.Request, item *access1.Entry) bool {
	filter := core.NewOrigin(req.URL.Query())
	target := core.Origin{Region: item.Region, Zone: item.Zone, SubZone: item.SubZone, Host: item.Host}
	if core.OriginMatch(target, filter) {
		return true
	}
	return false
}

func mapResource(r *http.Request) string {
	return module.TimeseriesAccessResourceV1

}
