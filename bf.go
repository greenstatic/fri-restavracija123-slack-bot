package main

import (
	"fmt"
	"github.com/andybalholm/cascadia"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	bfScrapeUrl    = "https://www.studentska-prehrana.si/sl/restaurant/MenuForDate"
	bfRestaurantid = 2023
)

type Bf struct {
}

func (_ *Bf) Name() string {
	return "BF"
}

func (bf *Bf) DailyMenu(t time.Time) (Menu, error) {
	url := bfScrapeUrl

	logrus.WithFields(logrus.Fields{"url": url, "cafeteria": bf.Name()}).Info("Creating HTTP POST Request for daily menu")

	reqBody := strings.NewReader(fmt.Sprintf("restaurantId=%d&date=%s", bfRestaurantid, t.Format("2.1.2006")))

	r, err := http.Post(url, "application/x-www-form-urlencoded", reqBody)
	if err != nil {
		return Menu{}, stacktrace.Propagate(err, "Failed to perform HTTP GET request")
	}

	return bf.parseDailyMenuHtml(r.Body)
}

func (bf *Bf) parseDailyMenuHtml(htmlMenu io.Reader) (Menu, error) {
	root, err := html.Parse(htmlMenu)
	if err != nil {
		return Menu{}, stacktrace.Propagate(err, "failed to parse html")
	}

	menuItemNodes := cascadia.QueryAll(root, cascadia.MustCompile(".shadow-wrapper"))
	menuItems := make([]MenuItem, 0, len(menuItemNodes))

	for _, menuItemNode := range menuItemNodes {
		nameNode := cascadia.Query(menuItemNode, cascadia.MustCompile("h5 strong"))
		if nameNode != nil && nameNode.FirstChild != nil {
			item := MenuItem{
				Name: deShitifyStudentskaPrehranaMenuName(nameNode.FirstChild.Data),
			}

			mealTypeNode := cascadia.Query(menuItemNode, cascadia.MustCompile("i img"))

			if mealTypeNode != nil {
				switch getHtmlAttribute(mealTypeNode.Attr, "src") {
				case "/Images/icnvegetarian.png":
					item.IsVegetarian = true
				case "/Images/icnfish.png":
					item.IsFish = true
				}
			}

			menuItems = append(menuItems, item)
		}
	}

	return menuItems, nil
}

func deShitifyStudentskaPrehranaMenuName(s string) string {
	if len(s) > 6 {
		if _, err := strconv.Atoi(string(s[0])); err == nil {
			// First character is digit
			s = s[6:]
		}
	}

	s = strings.ToLower(s)

	if len(s) > 0 {
		firstLetter := strings.ToUpper(string(s[0]))
		s = firstLetter + s[1:]
	}

	return s
}

func getHtmlAttribute(attrs []html.Attribute, key string) string {
	for _, attr := range attrs {
		if attr.Key == key {
			return attr.Val
		}
	}

	return ""
}
