package main

import (
	"fmt"
	"io/ioutil"
	"mapreduce/mr"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var count int

func Map(filename string, contents string) []mr.KeyValue {
	me := os.Getpid()
	f := fmt.Sprintf("mr-worker-jobcount-%d-%d", me, count)
	count++
	err := ioutil.WriteFile(f, []byte("x"), 0666)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Duration(2000+rand.Intn(3000)) * time.Millisecond)
	return []mr.KeyValue{mr.KeyValue{"a", "x"}}
}

func Reduce(key string, values []string) string {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		panic(err)
	}
	invocations := 0
	for _, f := range files {
		if strings.HasPrefix(f.Name(), "mr-worker-jobcount") {
			invocations++
		}
	}
	return strconv.Itoa(invocations)
}
