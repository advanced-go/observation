package timeseries1

import (
	"context"
	"fmt"
	"github.com/advanced-go/observation/module"
	"github.com/advanced-go/stdlib/controller"
	"github.com/advanced-go/stdlib/core"
	"github.com/advanced-go/stdlib/host"
	"github.com/advanced-go/stdlib/httpx"
	"github.com/advanced-go/stdlib/json"
	"github.com/advanced-go/stdlib/messaging"
	"net/http"
	"time"
)

const (
	entriesJson = "file:///c:/Users/markb/GitHub/observation/timeseries1/timeseries1test/access-v1.json"
)

func init() {
	a, err1 := host.RegisterControlAgent(PkgPath, messageHandler)
	if err1 != nil {
		fmt.Printf("init(\"%v\") failure: [%v]\n", PkgPath, err1)
	}
	a.Run()
	initializeDocuments()
}

func messageHandler(msg *messaging.Message) {
	start := time.Now()
	switch msg.Event() {
	case messaging.StartupEvent:
		// Any processing for a Startup event would be here
		messaging.SendReply(msg, core.NewStatusDuration(http.StatusOK, time.Since(start)))
	case messaging.ShutdownEvent:
	case messaging.PingEvent:
		// Any processing for a Shutdown/Ping event would be here
		messaging.SendReply(msg, core.NewStatusDuration(http.StatusOK, time.Since(start)))
	}
}

var (
	content            = httpx.NewListContent[Entry, struct{}, struct{}](false, matchEntry, nil, nil)
	resource           = httpx.NewResource[Entry, struct{}, struct{}](module.TimeseriesAccessResource, content, nil)
	authority, hostErr = httpx.NewHost(module.TimeseriesAuthority, mapResource, resource.Do)
)

func initializeDocuments() {
	defer controller.DisableLogging(true)()
	if hostErr != nil {
		fmt.Printf("error: new resource %v", hostErr)
	}
	entries, status := json.New[[]Entry](entriesJson, nil)
	if !status.OK() {
		fmt.Printf("initializeDocuments.New() -> [status:%v]\n", status)
		return
	}
	cfg, ok := module.GetRoute(module.TimeseriesRouteName)
	if !ok {
		fmt.Printf("initializeDocuments.GetRoute() [ok:%v]\n", ok)
	}
	ctrl := controller.New(cfg, authority.Do)
	controller.RegisterController(ctrl)
	_, status = put[core.Output](context.Background(), nil, entries)
	if !status.OK() {
		fmt.Printf("initializeDocuments.put() [status:%v]\n", status)
	}

}

func matchEntry(req *http.Request, item *Entry) bool {
	filter := core.NewOrigin(req.URL.Query())
	target := core.Origin{Region: item.Region, Zone: item.Zone, SubZone: item.SubZone, Host: item.Host}
	if core.OriginMatch(target, filter) {
		return true
	}
	return false
}

func mapResource(r *http.Request) string {
	return module.TimeseriesAccessResource

}
