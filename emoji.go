package main

import (
	"math/rand"
	"time"
)

const (
	veganEmoji      = ":seedling:"
	vegetarianEmoji = ":green_salad:"
	fishEmoji       = ":fish:"
)

var foodEmojis = []string{":baby_bottle:",
	":beer:",
	":beers:",
	":cocktail:",
	":tropical_drink:",
	":wine_glass:",
	":fork_and_knife:",
	":pizza:",
	":hamburger:",
	":fries:",
	":poultry_leg:",
	":meat_on_bone:",
	":spaghetti:",
	":curry:",
	":fried_shrimp:",
	":bento:",
	":sushi:",
	":fish_cake:",
	":rice_ball:",
	":rice_cracker:",
	":rice:",
	":ramen:",
	":stew:",
	":oden:",
	":dango:",
	":egg:",
	":bread:",
	":doughnut:",
	":custard:",
	":icecream:",
	":ice_cream:",
	":shaved_ice:",
	":birthday:",
	":cake:",
	":cookie:",
	":chocolate_bar:",
	":candy:",
	":lollipop:",
	":honey_pot:",
	":apple:",
	":green_apple:",
	":tangerine:",
	":lemon:",
	":cherries:",
	":grapes:",
	":watermelon:",
	":strawberry:",
	":peach:",
	":melon:",
	":banana:",
	":pear:",
	":pineapple:",
	":sweet_potato:",
	":eggplant:",
	":tomato:",
	":corn:",
}

func randomEmoji() string {
	rand.Seed(time.Now().UnixNano()) // Not best pseudorandom practice, but for this it suffices
	return foodEmojis[rand.Int()%len(foodEmojis)]
}
