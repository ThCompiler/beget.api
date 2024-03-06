package main

import (
	"fmt"
)

func main() {
	var a int = 2
	fmt.Printf("%d\n", a)
	//client := core.Client{
	//	Login:    "vetan2o5",
	//	Password: "KRoAmaD8vNG",
	//}
	var tmp []int
	fmt.Println(cap(tmp))

	//req, err := core.PrepareRequest[result.FileList](
	//	client,
	//	backup.CallGetFileBackupList(),
	//)
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//resp, err := req.Do()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Printf("%v", *resp)
}
