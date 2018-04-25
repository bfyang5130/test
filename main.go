// Command keys is a chromedp example demonstrating how to send key events to
// an element.
package main

import (
	"context"
	"log"
	"time"
    "strings"
	"github.com/chromedp/chromedp"
	"fmt"
	//"github.com/chromedp/cdproto/cdp"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}
     //定义一个字符串，用来装页面返回来的内容
     var res string
	// run task list
	err = c.Run(ctxt, sendkeys(&res))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", res)
	/**
	countRes:=len(res)

	for i := 0; i < countRes; i++ {
		fmt.Printf("%#v\n", res[i])
	}
	*/
	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}
	s := strings.NewReader(res)
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(s)
	if err != nil {
		log.Fatal(err)
	}
	// Find the review items
	doc.Find(".test_1").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		band := s.Text()
		fmt.Printf("Review %d: %s\n", i, band,)
	})
}

func sendkeys(res *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate("F:/html/index.html"),
		chromedp.WaitVisible(`div.test_1`, chromedp.ByQuery),
		chromedp.OuterHTML(`body`,res,chromedp.ByQuery),
		//chromedp.WaitVisible(`#su`, chromedp.ByID),
		//chromedp.SetValue(`#kw`, "权力的游戏百度云", chromedp.ByID),
		//chromedp.Click(`//input[@id="su"]`, chromedp.NodeVisible),
		//chromedp.Text(`div.test_1`, res, chromedp.NodeVisible, chromedp.ByQueryAll),
		//chromedp.
		chromedp.Sleep(3 * time.Second),
	}
}

func printText(res *string){
	fmt.Printf("GET THE CODE:%s",res)
}