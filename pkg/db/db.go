package db

import (
	"archive/zip"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"sort"
	"strconv"
	"strings"
)

type FoodDB struct {
	Rows       int
	Foods      map[string]*Food
	Nutrients  map[string]*Nutrient // All nutrients in the database.
	Categories map[string]*Category // All food categories in the database.
}

// Any substance consumed by humans for nutrition, taste and/or aroma.
type Food struct {
	FDCID          string          // Unique permanent identifier of the food.
	Description    string          // Description of the food.
	CategoryID     string          // ID of the food category the food belongs to.
	ScientificName string          // The scientific name for the food.
	FoodKey        string          // A string of characters used to identify both the current and all historical records for a specific food.
	CategoryCode   string          // See `Category.Code`.
	CategoryDesc   string          // See `Category.Description`.
	Nutrients      []*FoodNutrient // Nutrient values this food.
}

// A nutrient value for a food.
type FoodNutrient struct {
	ID         string
	FDCID      string  // ID of the food this food nutrient pertains to.
	NutrientID string  // ID of the nutrient to which the food nutrient pertains.
	Amount     float64 // Amount of the nutrient per 100g of food. Specified in unit defined in the nutrient table.
	Name       string  // Name of the nutrient.
	UnitName   string  // The standard unit of measure for the nutrient (per 100g of food).
}

// The chemical constituent of a food (e.g. calcium, vitamin E)
// officially recognized as essential to human health.
type Nutrient struct {
	ID          string
	Name        string // Name of the nutrient.
	UnitName    string // The standard unit of measure for the nutrient (per 100g of food).
	NutrientNbr int64  // A unique code identifying a nutrient or food constituent.
}

// Foods of defined similarity.
type Category struct {
	ID          string
	Code        string // Food group code.
	Description string // Description of the food group.
}

const (
	FoundationFoodZippedCSV = "./FoodData_Central_Foundation_Food_csv_2021-04-28.zip"
	SupportingDataZippedCSV = "./FoodData_Central_Supporting_Data_csv_2021-04-28.zip"
)

func NewFoodDB() *FoodDB {
	db := &FoodDB{
		Foods:      make(map[string]*Food),
		Nutrients:  make(map[string]*Nutrient),
		Categories: make(map[string]*Category),
	}

	ReadZippedCSV(SupportingDataZippedCSV, db.ReadCSVRecord)
	ReadZippedCSV(FoundationFoodZippedCSV, db.ReadCSVRecord)

	// Dump(db.Foods["1750347"])

	newDB := db.NewDBWithUniqueFoodNamesAndAverageNutritionalValues()
	newDB.NutrientTopList("Cholesterol")

	return db
}

func (db *FoodDB) NewDBWithUniqueFoodNamesAndAverageNutritionalValues() *FoodDB {
	ndb := &FoodDB{
		Foods: make(map[string]*Food),
	}

	foods := make(map[string]*Food)
	nuts := make(map[string][]float64)
	// nutsAvgs := make(map[string]float64)

	for _, f := range db.Foods {
		for _, n := range f.Nutrients {
			key := f.Description + "<>" + n.Name + "<>" + n.UnitName
			nuts[key] = append(nuts[key], n.Amount)

			if _, found := foods[key]; found {
				continue
			}

			foods[key] = f
			ndb.Foods[f.FDCID] = f
			f.Nutrients = nil
		}
	}

	// Calculate averages.
	for k, vs := range nuts {
		if len(vs) > 0 {
			// Average nutrient value.
			var avg float64
			for _, v := range vs {
				avg += v
			}
			avg /= float64(len(vs))

			// fmt.Printf("k=%s v=%d avg=%f\n", k, len(vs), avg)

			f, found := foods[k]
			if !found {
				log.Fatalf("could not look up food with key %s", k)
			}

			parts := strings.SplitN(k, "<>", 3)

			f.Nutrients = append(f.Nutrients, &FoodNutrient{
				Name:     parts[1],
				UnitName: parts[2],
				Amount:   avg,
			})
		}
	}

	return ndb
}

func (db *FoodDB) NutrientTopList(nut string) {
	type top struct {
		F *Food
		N *FoodNutrient
	}
	var unsorted []top

	for _, f := range db.Foods {
		for _, fn := range f.Nutrients {
			if strings.Contains(fn.Name, nut) {
				unsorted = append(unsorted, top{f, fn})
			}
		}
	}

	fmt.Printf("\nTop foods by nutrient '%s'\n\n", nut)

	sort.SliceStable(unsorted, func(i, j int) bool {
		return unsorted[i].N.Amount > unsorted[j].N.Amount
	})

	for i, t := range unsorted {
		shortDesc := t.F.Description
		if len(shortDesc) > 49 {
			shortDesc = shortDesc[0:49]
		}

		shortCat := t.F.CategoryDesc
		if len(shortCat) > 29 {
			shortCat = shortCat[0:29]
		}

		fmt.Printf(
			"%04d. %-50s (%-30s) --  %-20s  (%.02f %s)\n",
			i+1,
			shortDesc,
			shortCat,
			t.N.Name,
			t.N.Amount,
			t.N.UnitName,
		)
	}
}

func (db *FoodDB) ReadCSVRecord(csvFile string, cols []string) error {
	db.Rows++

	if csvFile == "food.csv" {
		// Handle food.csv
		// fmt.Printf("handling %s row #%d ..\n", csvFile, db.Rows)

		// Dump(cols)

		f := new(Food)
		f.FDCID = cols[0]
		f.Description = strings.Trim(cols[2], "\t, \r\n")
		f.CategoryID = cols[3]
		// f.ScientificName = cols[6]
		// f.FoodKey = cols[7]

		// Dump(f)

		if f.CategoryID != "" {
			c, found := db.Categories[f.CategoryID]
			if !found {
				log.Fatalf("could not find food category id %s for food %s", f.CategoryID, f.FoodKey)
			}
			f.CategoryCode = c.Code
			f.CategoryDesc = c.Description
		}

		db.Foods[f.FDCID] = f
	}

	if csvFile == "food_nutrient.csv" {
		// Handle food_nutrient.csv
		// fmt.Printf("handling %s row #%d ..\n", csvFile, db.Rows)

		fn := new(FoodNutrient)
		fn.ID = cols[0]
		fn.FDCID = cols[1]
		fn.NutrientID = cols[2]
		fn.Amount, _ = strconv.ParseFloat(cols[3], 64)

		n, found := db.Nutrients[fn.NutrientID]
		if !found {
			log.Fatalf("could not find nutrient id %s", fn.NutrientID)
		}
		fn.Name = n.Name
		fn.UnitName = n.UnitName

		// Dump(fn)

		f, found := db.Foods[fn.FDCID]
		if !found {
			log.Fatalf("could not find food id %s", fn.FDCID)
		}
		f.Nutrients = append(f.Nutrients, fn)
	}

	if csvFile == "nutrient.csv" {
		// Handle nutrient.csv
		// fmt.Printf("handling %s row #%d ..\n", csvFile, db.Rows)

		n := new(Nutrient)
		n.ID = cols[0]
		n.Name = cols[1]
		n.UnitName = cols[2]
		n.NutrientNbr, _ = strconv.ParseInt(cols[3], 10, 64)

		// Dump(n)

		db.Nutrients[n.ID] = n
	}

	if csvFile == "food_category.csv" {
		// Handle food_category.csv
		// fmt.Printf("handling %s row #%d ..\n", csvFile, db.Rows)

		c := new(Category)
		c.ID = cols[0]
		c.Code = cols[1]
		c.Description = cols[2]

		// Dump(n)

		db.Categories[c.ID] = c
	}

	return nil
}

func ReadZippedCSV(file string, onRecord func(filename string, cols []string) error) {
	fmt.Printf("reading ZIP file: %s\n", file)
	zr, err := zip.OpenReader(file)
	if err != nil {
		log.Fatalf("could not open file %s: %s", file, err.Error())
	}
	defer zr.Close()

	for _, f := range zr.File {
		if !strings.HasSuffix(f.Name, ".csv") {
			fmt.Printf("skipping non-CSV file: %s\n", f.Name)
			continue
		}

		fmt.Printf("reading CSV file: %s\n", f.Name)
		r, err := f.Open()
		if err != nil {
			log.Fatalf("could not open file %s: %s", file, err.Error())
		}

		cr := csv.NewReader(r)

		var count int
		for {
			rec, err := cr.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error while reading file %s: %s", file, err.Error())
			}

			count++

			if count == 1 {
				// spew.Dump(rec)
			} else {
				onRecord(f.Name, rec)
			}
		}

		r.Close()
	}
}

func Dump(o interface{}) {
	j, _ := json.MarshalIndent(o, "", "  ")
	fmt.Println(string(j))
}
