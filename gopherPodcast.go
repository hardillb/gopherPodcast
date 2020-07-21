package main

import (
		"os"
		"io"
    "fmt"
    "strings"
    "net/http"

    "github.com/mmcdole/gofeed"
)

func main() {

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) == 0 {
		fmt.Println("You need to pass a feed URL")
		os.Exit(0)
	}

	fp := gofeed.NewParser()
	feed, err := fp.ParseURL(argsWithoutProg[0])

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	output := feed.Title + "\n\n"

	for _, item := range feed.Items {
		//fmt.Println(i,item.Title)
		dirname := strings.ReplaceAll(item.Title, " ", "_")
		if _, err := os.Stat(dirname); os.IsNotExist(err) {
			os.Mkdir(dirname, 0755)
			filename := strings.ReplaceAll(item.Title + ".mp3", " ", "_")
			downloadFile(item.Enclosures[0].URL, dirname + "/" + filename)
			output += "9" + item.Title + ".mp3\t" + dirname + "/" + filename + "\n"
		}
	}

	fmt.Println(output)
}

func downloadFile(url string, location string) error {
	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(location)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}