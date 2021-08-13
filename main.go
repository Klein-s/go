package main

import (
	"github.com/klein/go-mvc/bootstrap"
	"github.com/klein/go-mvc/config"
	c "github.com/klein/go-mvc/pkg/config"
	"github.com/klein/go-mvc/routes"
)


func init()  {
	config.Initialize()
}

func main()  {

	bootstrap.SetupDB()
	r := routes.RegisterWebRoutes()

	r.Run(":"+ c.GetString("app.port"))
	//ch := make(chan int)
	//var y int
	//for i := 0; i<=1000000; i++ {
	//	go run(ch, i)
	//	x := <-ch
	//	y = y + x
	//}
	//
	//fmt.Println(y)
}


//func run(ch chan int, num int) {
//	ch <- num + 1
//	time.Sleep(1 * time.Second)
//}