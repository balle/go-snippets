package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func die(msg string) {
	fmt.Printf(msg)
	os.Exit(0)

}

func main() {
	type Item struct {
		XMLName     xml.Name `xml:"item"`
		Game        string   `xml:"game,attr"`
		NextJp      string   `xml:"nextjp,attr"`
		NextDd      string   `xml:"nextdd,attr"`
		WinNums     string   `xml:"winnum,attr"`
		WinDd       string   `xml:"windd,attr"`
		Myflv       string   `xml:"myflv,attr"`
		WinnumNM    string   `xml:"winnumNM,attr"`
		Name        string   `xml:"name,attr"`
		Title       string   `xml:"title"`
		Link        string   `xml:"link"`
		Description string   `xml:"description"`
		PubDate     string   `xml:"pubdate"`
		Guid        string   `xml:"guid"`
	}

	type Channel struct {
		XMLName     xml.Name `xml:"channel"`
		Title       string   `xml:"title"`
		Link        string   `xml:"link"`
		Description string   `xml:"description"`
		Items       []Item   `xml:"item"`
	}

	type Rss struct {
		XMLName     xml.Name `xml:"rss"`
		Version     string   `xml:"version,attr"`
		Channel     Channel  `xml:"channel"`
		Description string   `xml:"description"`
		Title       string   `xml:"title"`
		Link        string   `xml:"link"`
	}

	res, err := http.Get("https://www.heise.de/rss/heise.rdf")

	if err != nil {
		die(fmt.Sprintf("Cannot get heise rss feed: %v", err))

	}

	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)

	if err != nil {
		die(fmt.Sprintf("Cannot read content: %v", err))
	}

	rss := Rss{}
	err = xml.Unmarshal(content, &rss)

	if err != nil {
		die(fmt.Sprintf("Cannot parse xml: %v", err))
	}

	for _, item := range rss.Channel.Items {
		fmt.Println(item.Title)
		fmt.Println(item.Link)
		fmt.Println("")
	}
}
