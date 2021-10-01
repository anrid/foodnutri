package main

import (
	"os"

	"github.com/anrid/foodnutri/pkg/db"
	"github.com/spf13/pflag"
)

func main() {
	nut := pflag.StringP("nutrient", "n", "", "Select a nutrient, e.g. Cholesterol (required).")

	pflag.Parse()

	if *nut == "" {
		pflag.PrintDefaults()
		os.Exit(-1)
	}

	// Load food DB.
	db := db.NewFoodDB()

	// List top foods by nutrient content.
	db.NutrientTopList(*nut)
}
