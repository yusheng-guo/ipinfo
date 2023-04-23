package main

import "github.com/yushengguo557/ipinfo/cmd"

func main() {
	cmd.Execute()
}

// func main() {
// 	c := colly.NewCollector()

// 	// Find and visit all links
// 	c.OnHTML("a", func(e *colly.HTMLElement) {
// 		e.Request.Visit(e.Attr("href"))
// 	})

// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL)
// 	})

// 	c.Visit("https://www.bilibili.com/")
// }
