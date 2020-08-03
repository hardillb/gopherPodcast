package main

import (
		"os"
    "fmt"
    "strings"
    "io/ioutil"
    "github.com/mmcdole/gofeed"
    //"github.com/eidolon/wordwrap"
    "jaytaylor.com/html2text"
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

	//wrapper := wordwrap.Wrapper(60, false)

	output := feed.Title + "\n\n"

	for _, item := range feed.Items {
		dirname := strings.ReplaceAll(item.Title, " ", "_")
		filename := strings.ReplaceAll(item.Title, " ", "_")
		if _, err := os.Stat(dirname); os.IsNotExist(err) {
			os.Mkdir(dirname, 0755)
			body, _ := html2text.FromString(item.Content, html2text.Options{PrettyTables: true})
			writeFile(body, dirname + "/" + filename)
		}
		output += " -- " + item.Title + " --\n\n"
		output += "0" + item.Title + "\t" + dirname + "/" + filename + "\n\n"
	}

	fmt.Println(output)
}

func writeFile(body string, location string) error {
	d1 := []byte(body)
	err := ioutil.WriteFile(location, d1, 0644)
	return err
}