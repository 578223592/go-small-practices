package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"log"
	"math/rand"
	"time"
)

func getData(id int64) string {
	log.Printf("query...%d\n", id)
	time.Sleep(10 * time.Millisecond) // 模拟一个比较耗时的操作,10ms
	return "liwenzhou.com" + fmt.Sprintf("%d", id)
}

func main() {
	log.SetFlags(log.Lmicroseconds)
	g := new(singleflight.Group)
	var i int64 = 0
	for {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		// 调用
		go func() {
			icopy := i
			i++
			v1, _, shared := g.Do("getData", func() (interface{}, error) {
				ret := getData(icopy)
				return ret, nil
			})
			log.Printf("call: v1:%v, shared:%v , i:%d\n", v1, shared, i)
		}()

	}

}

/**
query...1
1st call: v1:liwenzhou.com1, shared:true
2nd call: v2:liwenzhou.com1, shared:true
*/

