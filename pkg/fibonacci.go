package pkg

import (
    "context"
    "github.com/go-redis/redis/v8"
	"strconv"	
)

var ctx = context.Background()

func Fib(n int) (int) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	
    res, err := rdb.Get(ctx, strconv.Itoa(n)).Result()	
    if err == redis.Nil {
    if n == 0 {
        return 0
    } else if n == 1 {
		return 1
	} else {
		res := Fib(n-1) + Fib(n-2)
		err = rdb.Set(ctx, strconv.Itoa(n), strconv.Itoa(res), 0).Err()
    if err != nil {
        panic(err)
    }
		return res
	} 
} else {
	res2, _ := strconv.Atoi(res)
	return res2
}
}

func FibSlice(n1 int, n2 int) ([]int, []int) {
	var ret1 []int
	var ret2 []int
	for i:=n1; i<n2+1; i++ {
		ret1= append(ret1, i)
		ret2 = append(ret2,Fib(i))
	//fmt.Println(i, Fib(i))
	}
	return ret1, ret2

}

func FibSliceRPC(n1 int, n2 int) ([]int32) {	
	var ret2 []int32
	for i:=n1; i<n2+1; i++ {		
		ret2 = append(ret2,int32(Fib(i)))	
	}
	return ret2
}