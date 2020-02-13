package main

import (
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/session/redis"
	_ "my-first-beego-project/routers"
)

func main() {
	// for set static path, by default beego set url "/static" to "static" path
	// remove default config path
	beego.DelStaticPath(`static`)
	beego.SetStaticPath("/images", "static/img")

	// The file format of LoadAppConfig. By default this is ini. Other valid formats include xml, yaml, and json.
	// The application configuration file path. By default this is conf/app.conf.
	beego.LoadAppConfig("yaml", "conf/app.conf")
	//// App config
	// Enable auto render. By default this is True. This value should be set to false for API applications, as there is no need to render templates.
	beego.BConfig.WebConfig.AutoRender = true
	// Enable Docs. By default this is False.
	beego.BConfig.WebConfig.EnableDocs = false
	// Sets the Flash Cookie name. By default this is BEEGO_FLASH.
	beego.BConfig.WebConfig.FlashName = "BEEGO_FLASH"
	// Enable listing of the static directory. By default this is False and will return a 403 error.
	beego.BConfig.WebConfig.DirectoryIndex = false
	// Sets the static file dir(s). By default this is static.
	// 1. Single dir, StaticDir = download. Same as beego.SetStaticPath("/download","download")
	// 2. Multiple dirs, StaticDir = download:down download2:down2. Same as beego.SetStaticPath("/download","down") and beego.SetStaticPath("/download2","down2")
	beego.BConfig.WebConfig.StaticDir = map[string]string{"download":"download"}
	// StaticExtensionsToGzip
	// Sets a list of file extensions that will support compression by Gzip. The formats .css and .js are supported by default.
	beego.BConfig.WebConfig.StaticExtensionsToGzip = []string{".css", ".js"}
	//Same as in config file StaticExtensionsToGzip = .css, .js
	// Left mark of the template, {{ by default.
	beego.BConfig.WebConfig.TemplateLeft = "{{"
	// Right mark of the template, }} by default.
	beego.BConfig.WebConfig.TemplateRight = "}}"
	// Set the location of template files. This is set to views by default.
	beego.BConfig.WebConfig.ViewsPath = "views"
	// Enable XSRF
	beego.BConfig.WebConfig.EnableXSRF = false
	// Set the XSRF expire time. By default this is set to 0.
	beego.BConfig.WebConfig.XSRFExpire = 0

	//// HTTP Server config
	// Enable graceful shutdown. By default this is False.
	beego.BConfig.Listen.Graceful = false
	// Set the http timeout. By default thi is ‘0’, no timeout.
	beego.BConfig.Listen.ServerTimeOut = 0
	// Set the address type. default is tcp6 but we can set it to true to force use TCP4.
	beego.BConfig.Listen.ListenTCP4 = true
	// Enable HTTP listen. By default this is set to True.
	beego.BConfig.Listen.EnableHTTP = true
	// Set the address the app listens to. By default this value is empty and the app will listen to all IPs.
	beego.BConfig.Listen.HTTPAddr = ""
	// Set the port the app listens on. By default this is 8080
	beego.BConfig.Listen.HTTPPort = 8080
	// Enable HTTPS. By default this is False. When enabled HTTPSCertFile and HTTPSKeyFile must also be set.
	beego.BConfig.Listen.EnableHTTPS = false
	// beego.BConfig.Listen.HTTPSAddr = ""
	beego.BConfig.Listen.HTTPSAddr = ""
	// Set the port the app listens on. By default this is 10443
	beego.BConfig.Listen.HTTPSPort = 10443
	// Set the SSL cert path. By default this value is empty.
	beego.BConfig.Listen.HTTPSCertFile = ""
	// beego.BConfig.Listen.HTTPSCertFile = "conf/ssl.crt"
	// Set the SSL key path. By default this value is empty.
	beego.BConfig.Listen.HTTPSKeyFile = ""
	// beego.BConfig.Listen.HTTPSKeyFile = "conf/ssl.key"
	// Enable supervisor module. By default this is False.
	beego.BConfig.Listen.EnableAdmin = false
	// Set the address the admin app listens to. By default this is blank and the app will listen to any IP.
	beego.BConfig.Listen.AdminAddr = ""
	// Set the port the admin app listens on. By default this is 8088.
	beego.BConfig.Listen.AdminPort = 8088
	// Enable fastcgi. By default this is False.
	beego.BConfig.Listen.EnableFcgi = false
	// Enable fastcgi standard I/O or not. By default this is False.
	beego.BConfig.Listen.EnableStdIo = false

	//// Session config
	// Enable session. By default this is False.
	beego.BConfig.WebConfig.Session.SessionOn = false
	// Set the session provider. By default this is memory.
	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	// Set the session cookie name stored in the browser. By default this is beegosessionID.
	beego.BConfig.WebConfig.Session.SessionName = "beegosessionID"
	// Set the session expire time. By default this is 3600s.
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600
	// Set the valid expiry time of the cookie in browser for session. By default this is 3600s.
	beego.BConfig.WebConfig.Session.SessionCookieLifeTime = 3600
	// Enable SetCookie. By default this is True.
	beego.BConfig.WebConfig.Session.SessionAutoSetCookie = true
	// Set the session cookie domain. By default this is empty.
	beego.BConfig.WebConfig.Session.SessionDomain = ""

	//// Log config
	// Enable output access logs. By default these logs will not be output under ‘prod’ mode.
	beego.BConfig.Log.AccessLogs = false
	// Toggle printing line numbers. By default this is True. This config is not supported in config file.
	beego.BConfig.Log.FileLineNum = true
	// Log outputs config. This config is not supported in config file.
	beego.BConfig.Log.Outputs = map[string]string{"console": ""}
	// or
	beego.BConfig.Log.Outputs["console"] = ""

	// another config
	//beego.AppConfig.String("mysqluser")
	//beego.AppConfig.String("mysqlpass")
	//beego.AppConfig.String("mysqlurls")
	//beego.AppConfig.String("mysqldb")

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Run()
}
