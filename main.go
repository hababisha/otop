package main


import "github.com/hababisha/otop/router"


func main(){
	r:= router.Router()
	r.Run(":8080")
}