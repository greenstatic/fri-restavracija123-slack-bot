package bot

import (
	"github.com/greenstatic/fri-restavracija123-slack-bot/restavracija123"
	"strings"
)

func MenuMarkdownContent(m []restavracija123.MenuItem) string {
	s := strings.Builder{}
	s.WriteString(randomEmoji() + " Dnevni meni\n")

	menuItems := make([]string, 0, len(m))

	for _, e := range m {

		t := RemoveAsterisks(e.Title)
		emoji := FoodEmoji(e)
		if emoji != "" {
			t += " " + emoji
		}

		menuItems = append(menuItems, t)
	}

	s.WriteString(MarkdownList(menuItems))

	return s.String()
}

func RemoveAsterisks(str string) string {
	return strings.ReplaceAll(str, "*", "")
}

func MarkdownList(xs []string) string {
	s := strings.Builder{}

	for i, e := range xs {
		s.WriteString("* ")
		s.WriteString(e)

		// if not last element
		if i != len(xs)-1 {
			s.WriteString("\n")
		}
	}

	return s.String()
}

const (
	veganId      = "13186"
	vegetarianId = "3725"
)

func FoodEmoji(food restavracija123.MenuItem) string {
	typeIds := food.MenuTypeIds()
	if StringContains(typeIds, veganId) {
		return veganEmoji
	}

	if StringContains(typeIds, vegetarianId) {
		return vegetarianEmoji
	}

	return ""
}

func StringContains(xs []string, s string) bool {
	for _, e := range xs {
		if e == s {
			return true
		}
	}

	return false
}
