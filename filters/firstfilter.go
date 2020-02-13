package filters

import (
	"fmt"
	"github.com/astaxie/beego/context"
	"time"
)

var LogManager = func(ctx *context.Context) {
	fmt.Println("IP :: " + ctx.Request.RemoteAddr + ", Time:: " + time.Now().Format(time.RFC850))
}
