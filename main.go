package main

import (
	"github.com/hababisha/otop/router"
	"github.com/hababisha/otop/utils"
)


func main(){
	utils.ConnectDB()
	r:= router.Router()
	r.Run(":8080")
}