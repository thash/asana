package utils

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"runtime"
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
		log.Fatalf("fatal: %v\n", err)
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

// base: github.com/github/hub
func BrowserLauncher() (string, error) {
	browser := os.Getenv("BROWSER")
	if browser == "" {
		browser = searchBrowserLauncher(runtime.GOOS)
	}

	if browser == "" {
		return "", errors.New("Please set $BROWSER to a web launcher")
	}

	return browser, nil
}

func searchBrowserLauncher(goos string) (browser string) {
	switch goos {
	case "darwin":
		browser = "open"
	case "windows":
		browser = "cmd /c start"
	default:
		candidates := []string{"xdg-open", "cygstart", "x-www-browser", "firefox",
			"opera", "mozilla", "netscape"}
		for _, b := range candidates {
			path, err := exec.LookPath(b)
			if err == nil {
				browser = path
				break
			}
		}
	}

	return browser
}
