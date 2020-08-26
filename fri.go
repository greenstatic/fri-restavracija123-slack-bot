package main

import (
	"encoding/json"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	friBaseUrl      = "https://www.restavracija123.si/api/"
	friId           = "3798" // acquired by monitoring traffic from: https://www.restavracija123.si
	friVeganId      = "13186"
	friVegetarianId = "3725"
)

type Fri struct {
}

type friMenuItem struct {
	FoodId   string `json:"foodId"`
	Title    string `json:"title"`
	MenuType []struct {
		TypeId string
		Title  string
	} `json:"menu_type"`
}

type friDailyMenuJsonResponse struct {
	Dnevna      []friMenuItem `json:"dnevna"`
	Priporocamo []friMenuItem `json:"priporocamo"`
}

func (_ *Fri) Name() string {
	return "FRI"
}

func (fri *Fri) DailyMenu(t time.Time) (Menu, error) {
	url := friBaseUrl + "getDailyMenu/" + friId + "/" + t.Format("2006-01-02")

	logrus.WithFields(logrus.Fields{"url": url, "cafeteria": fri.Name()}).Info("Creating HTTP GET Request for daily menu")

	r, err := http.Get(url)
	if err != nil {
		return Menu{}, stacktrace.Propagate(err, "Failed to perform HTTP GET request")
	}

	raw, err := ioutil.ReadAll(r.Body)

	menuResp := friDailyMenuJsonResponse{}
	if err := json.Unmarshal(raw, &menuResp); err != nil {

		// Check if empty list
		if string(raw) == "[]" {
			return Menu{}, nil
		}

		return Menu{}, stacktrace.Propagate(err, "Failed to parse JSON")
	}

	return friMenuItemsFromIntermediateStruct(friMenuItemsFromApiResponse(menuResp)), nil
}

func friMenuItemsFromApiResponse(d friDailyMenuJsonResponse) []friMenuItem {
	menuItems := make([]friMenuItem, 0)

	for _, entry := range d.Dnevna {
		menuItems = append(menuItems, entry)
	}

	for _, entry := range d.Priporocamo {
		menuItems = append(menuItems, entry)
	}

	return menuItems
}

func friMenuItemsFromIntermediateStruct(items []friMenuItem) []MenuItem {
	menuItems := make([]MenuItem, 0, len(items))

	for _, item := range items {
		mItem := MenuItem{
			Name: item.Title,
		}

		if StringContains(item.MenuTypeIds(), friVeganId) {
			mItem.IsVegan = true
		}

		if StringContains(item.MenuTypeIds(), friVegetarianId) {
			mItem.IsVegetarian = true
		}

		menuItems = append(menuItems, mItem)
	}

	return menuItems
}

func (m friMenuItem) MenuTypeIds() []string {
	xs := make([]string, 0, len(m.MenuType))
	for _, e := range m.MenuType {
		xs = append(xs, e.TypeId)
	}

	return xs
}
