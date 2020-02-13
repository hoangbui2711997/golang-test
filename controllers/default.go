package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"math/big"
	"sync"
	"time"
)

type Student struct {
	Name string
	Age int
}

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *UserController) Get() {
	var wg sync.WaitGroup

	total := big.NewInt(1)
	c.Data["Website"] = "My website"
	c.Data["Email"] = "hoang2711997@gmail.com"
	chans := make(chan string, 10)
	numbers := []int64{100, 200, 300, 400, 500, 600, 700, 800, 900, 1000, 1100}
	c.Data["numbers"] = numbers
	c.Data["object"] = Student{ Name: "Test", Age: 10 }
	if beegoCache.IsExist("total") {
		c.Data["total"] = beegoCache.Get("total")
		c.TplName = "user.tpl"
		return
	}
	for num := range numbers {
		wg.Add(1)
		log.Println("num:", num)
		//go worker(numbers[num] * 10, chans, &wg)
		go worker(numbers[num] / 10, chans, &wg)
		workResult, _ := new(big.Int).SetString(<-chans, 10)
		total.Mul(total, workResult)
		log.Println("total:", total.Text(10))
	}
	c.Data["total"] = total.Text(10)
	beegoCache.Put("total", total.Text(10), 1000000 * time.Second)
	c.TplName = "user.tpl"
}

func worker(count int64, c chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	var result int64 = 1
	var i int64 = 1
	product := big.NewInt(1)
	log.Println("__result:", result)
	for ; i <= count; i++ {
		product.Mul(product, new(big.Int).SetInt64(i))
	}
	//newBigInt := big.NewInt(int64(100))
	//bigInt1, _ := new(big.Int).SetString("1723681723698127649817263981726172368172369812764981726398172617236817236981276498172639817261723681723698127649817263981726", 10)
	//bigInt2, _ := new(big.Int).SetString("1723681723698127649817263981726172368172369812764981726398172617236817236981276498172639817261723681723698127649817263981726", 10)
	//log.Println(bigInt1.Text(10))
	//product := new(big.Int).Mul(bigInt1, bigInt2)
	//log.Println(product.Text(10))

	log.Println("result:", result)
	c <- product.Text(10)
	fmt.Printf("Worker count: %d done\n", count)
}

type FirstController struct {
	beego.Controller
}

type Employee struct {
	Id int `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

type Employees []Employee
var employees Employees

func init() {
	employees = Employees{
		Employee{Id: 1, FirstName: "Foo", LastName: "Bar"},
		Employee{Id: 2, FirstName: "Baz", LastName: "Qux"},
		Employee{Id: 3, FirstName: "After", LastName: "Deploy"},
	}
}

func (this *FirstController) GetEmployees() {
	// 200 mean success
	this.Ctx.ResponseWriter.WriteHeader(200)
	this.Data["json"] = employees
	this.ServeJSON()
}
