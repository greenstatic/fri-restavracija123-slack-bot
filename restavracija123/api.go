package restavracija123

import (
	"encoding/json"
	"github.com/palantir/stacktrace"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	baseUrl = "https://www.restavracija123.si/api/"
	friId   = "3798" // acquired by monitoring traffic from: https://www.restavracija123.si/
)

type MenuItem struct {
	FoodId   string `json:"foodId"`
	Title    string `json:"title"`
	MenuType []struct {
		TypeId string
		Title  string
	} `json:"menu_type"`
}

type DailyMenuJsonResponse struct {
	Dnevna      []MenuItem `json:"dnevna"`
	Priporocamo []MenuItem `json:"priporocamo"`
}

func DailyMenu(t time.Time) ([]MenuItem, error) {
	url := baseUrl + "getDailyMenu/" + friId + "/" + t.Format("2006-01-02")

	logrus.WithField("url", url).Info("Creating HTTP GET Request for daily menu")

	r, err := http.Get(url)
	if err != nil {
		return []MenuItem{}, stacktrace.Propagate(err, "Failed to perform HTTP GET request")
	}

	raw, err := ioutil.ReadAll(r.Body)

	menuResp := DailyMenuJsonResponse{}
	if err := json.Unmarshal(raw, &menuResp); err != nil {

		// Check if empty list
		if string(raw) == "[]" {
			return []MenuItem{}, nil
		}

		return []MenuItem{}, stacktrace.Propagate(err, "Failed to parse JSON")
	}

	return MenuItemsFromApiResponse(menuResp), nil
}

func MenuItemsFromApiResponse(d DailyMenuJsonResponse) []MenuItem {
	menuItems := make([]MenuItem, 0)

	for _, entry := range d.Dnevna {
		menuItems = append(menuItems, entry)
	}

	for _, entry := range d.Priporocamo {
		menuItems = append(menuItems, entry)
	}

	return menuItems
}

func (m MenuItem) MenuTypeIds() []string {
	xs := make([]string, 0, len(m.MenuType))
	for _, e := range m.MenuType {
		xs = append(xs, e.TypeId)
	}

	return xs
}
