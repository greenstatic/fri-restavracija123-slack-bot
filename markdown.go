package main

import (
	"fmt"
	"strings"
)

func MenuMarkdownContent(c Cafeteria, m Menu) string {
	s := strings.Builder{}
	s.WriteString(MarkdownBold(fmt.Sprintf("%s dnevni meni:", c.Name())))
	s.WriteString("\n")

	menuItems := make([]string, 0, len(m))

	for _, e := range m {
		t := RemoveAsterisks(e.Name)
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

func MarkdownBold(s string) string {
	return fmt.Sprintf("*%s*", s)
}

func FoodEmoji(item MenuItem) string {
	if item.IsVegan {
		return veganEmoji
	}

	if item.IsVegetarian {
		return vegetarianEmoji
	}

	if item.IsFish {
		return fishEmoji
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
