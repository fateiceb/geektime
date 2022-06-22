package main

import (
	"crypto/md5"
	"fmt"
	// "github.com/sony/sonyflake"
)

// var sf *sonyflake.Sonyflake

// func init() {
// 	var f sonyflake.Settings
// 	f.StartTime = time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC)
// 	sf = sonyflake.NewSonyflake(f)
// 	if sf == nil {
// 		panic("id generator init error.")
// 	}
// }
func main() {
	fmt.Println(Md5String("1d"))
}
func Md5String(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}
