package controllers

import (
  "fmt"
  "github.com/astaxie/beego"
  "github.com/astaxie/beego/cache"
  "log"
  "reflect"
  "time"
)

type CacheController struct
{
  beego.Controller
}

var beegoCache cache.Cache
var err error

func init() {
  log.Println("@@Cache")
  beegoCache, _ = cache.NewCache("memory", `{"interval": 60}`)
  beegoCache.Put("foo", "bar", 1000000 * time.Second)
  log.Println("Foo:", beegoCache.Get("foo"))
  log.Printf("%+v", reflect.TypeOf(beegoCache))
}


func (this *CacheController) GetFromCache() {
  fmt.Printf("%+v", reflect.TypeOf(beegoCache))
  log.Println("@@GetFromCache")
  foo := beegoCache.Get("foo")

  this.Ctx.WriteString("Hello " + fmt.Sprintf("%v", foo))
}

func (this *CacheController) Finish() {
  log.Println("@Finish")
}

func (this *CacheController) Prepare() {
  log.Println("@Prepare")
}

func (this *CacheController) GetResult() {
  keyCache := this.Ctx.Input.Param(":key")
  log.Println("keyCache:", keyCache)
  this.Ctx.WriteString(fmt.Sprintf("%v", beegoCache.Get(keyCache)))
}
