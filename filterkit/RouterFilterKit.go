package filterkit

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"git.gumpcome.com/go_kit/logkit"
	"time"
	"github.com/astaxie/beego"
)

var apiStatistics = make(map[string]*apiInfo)

type apiInfo struct {
	url       string
	startTime time.Time
	elapsed   time.Duration
	params    string
}

func BeforeRouterFilter(ctx *context.Context) {
	ctx.Request.ParseForm()
	url := ctx.Input.URL()
	params := fmt.Sprint(ctx.Request.Form)
	now := time.Now()

	info := apiStatistics[url]
	if info == nil {
		apiStatistics[url] = &apiInfo{
			elapsed: 0,
		}
	}

	info = apiStatistics[url]
	info.url = url
	info.startTime = now
	info.params = params
}

func FinishRouterFilter(ctx *context.Context) {
	url := ctx.Input.URL()
	info := apiStatistics[url]
	log := logkit.GetLog()
	if info == nil || log == nil {
		return
	}
	info.elapsed = time.Since(info.startTime)

	log.Debug("%s|%s|%s", info.url, info.elapsed.String(), info.params)
}

func InItBeegoFilter(pattern string) {
	log := logkit.GetLog()
	if log == nil {
		return
	}
	beego.InsertFilter(pattern, beego.BeforeRouter, BeforeRouterFilter, false)
	beego.InsertFilter(pattern, beego.FinishRouter, FinishRouterFilter, false)
	log.Info("%s %s", pattern, "filter init success")
}
