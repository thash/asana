package utils

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"time"
)

const (
	CacheFileName = ".asana.cache"
)

func Home() string {
	current, err := user.Current()
	Check(err)
	return current.HomeDir
}

func Check(err error) {
	if err != nil {
		log.Fatal("fatal: %v\n", err)
	}
}

func EndlessSelect(max int, index int) int {
	print("\nChoose one out of them: ")
	fmt.Scanf("%d", &index)
	if index <= max {
		return index
	}
	return EndlessSelect(max, index)
}

func Older(duration string, cacheFile string) bool {
	st, err := os.Stat(cacheFile)
	if os.IsNotExist(err) {
		return true
	}
	d, err := time.ParseDuration(duration)
	Check(err)
	return time.Now().After(st.ModTime().Add(d))
}

func CacheFile() string {
	return Home() + "/" + CacheFileName
}
