package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
)

func main() {
	http.HandleFunc("/", feedHandler)
	http.HandleFunc("/feed", func(w http.ResponseWriter, r *http.Request) {
		feed := fetchFeed()
		// Set the content type to HTML
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// Start the HTML document
		fmt.Fprintf(w, "<html><body>")

		for _, item := range feed.Items {
			// Format the feed item using HTML tags
			fmt.Fprintf(w, "<h2>%s</h2>", item.Title)
			fmt.Fprintf(w, "<p><b>Description:</b> %s</p>", item.Description)
			fmt.Fprintf(w, "<p><b>Link:</b> <a href=\"%s\">%s</a></p>", item.Link, item.Link)
			fmt.Fprintf(w, "<hr>")
		}

		// End the HTML document
		fmt.Fprintf(w, "</body></html>")

	})
	// go fetchFeedPeriodically()
	// server := &http.Server{
	// 	Addr:         ":8081",
	// 	Handler:      http.DefaultServeMux,
	// 	TLSConfig:    nil,
	// 	TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	// }
	fmt.Println("Starting server on http://localhost:3306")
	log.Fatal(http.ListenAndServe(":3306", nil))
	// err := server.ListenAndServe()
	// if err != nil {
	// 	log.Fatal("Server error: ", err)
	// }
}

func fetchFeedPeriodically() {
	ticker := time.NewTicker(1 * time.Minute)

	for range ticker.C {
		go fetchFeed()
	}
}

func fetchFeed() *gofeed.Feed {
	fp := gofeed.NewParser()

	// Enter the valid RSS feed URL. Note that this is just an example feed url
	feedURL := "https://www.upwork.com/ab/feed/jobs/rss?q=JavaScript+OR+TypeScript+OR+jQuery+OR+HTML+OR+CSS+OR+React+OR+Next.js+OR+RESTful+API+OR+API+Integration+OR+Node.js+OR+ExpressJS+OR+Redux+OR+GraphQL+OR+Web+OR+Mobile&sort=recency&paging=0%3B10&api_params=1&securityToken=fe1fefd8c06ee5fd011fb49739610b1783945e3422bb86c573eb8328182475afb15cbc2f88bf8c6012e183a557fdf61be5591f3f8871d357655342299371dd14&userUid=1682285877314486272&orgUid=1682285877314486273"
	// feedURL := "https://www.upwork.com/ab/feed/jobs/rss?q=Go+OR+React+OR+Laravel+OR+VUE+OR+Microservice&sort=recency&job_type=hourly%2Cfixed&client_hires=0%2C1-9%2C10-&proposals=0-4%2C5-9%2C10-14%2C15-19&budget=500-999%2C1000-4999%2C5000-&verified_payment_only=1&location=Argentina%2CArmenia%2CAustralia%2CFinland%2CFrance%2CGermany%2CKuwait%2CLuxembourg%2CMalaysia%2CNetherlands%2CNew+Zealand%2CNorway%2CSingapore%2CUnited+Arab+Emirates%2CUnited+Kingdom%2CUnited+States%2CAustralia+and+New+Zealand%2CNorthern+America%2CNorthern+Europe%2CWestern+Europe%2CAmericas%2CEurope%2COceania&paging=0%3B20&api_params=1&securityToken=7b26c330ff1063ae5e66f3c62881c8840aea9029dfc11abb0ea233bd56f2490f23335b271aaf4b40365f7215d077bd282d0c2e87ed5d1c125102405702520292&userUid=1681235169147748352&orgUid=1681235169147748353"

	feed, err := fp.ParseURL(feedURL)
	if err != nil {
		log.Printf("Error while parsing the feed: %v", err)
		return nil
	}

	for _, item := range feed.Items {
		fmt.Println(item.Title)
	}

	log.Println("Feed fetched successfully")
	return feed
}

func feedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the Feed Handler!")
}
