# Food Nutri

List top foods by nutrient content, e.g. Cholesterol.

### Produces output like this:

```bash
% go run cmd/cli/cli.go -n Cholesterol

...

reading ZIP file: ./FoodData_Central_Supporting_Data_csv_2021-04-28.zip
skipping non-CSV file: Download & API Field Descriptions April 2021.pdf
reading CSV file: food_category.csv
reading CSV file: nutrient.csv
reading ZIP file: ./FoodData_Central_Foundation_Food_csv_2021-04-28.zip
reading CSV file: food.csv
reading CSV file: food_nutrient.csv

Top foods by nutrient 'Cholesterol' (amount of the nutrient per 100g of food):

0001. Egg yolks, dried                                   (Dairy and Egg Products        ) --  Cholesterol           (2345.29 MG)
0002. Egg, yolk, dried                                   (Dairy and Egg Products        ) --  Cholesterol           (2340.00 MG)
0003. Egg, whole, dried                                  (Dairy and Egg Products        ) --  Cholesterol           (1700.00 MG)
0004. Whole eggs, dried                                  (Dairy and Egg Products        ) --  Cholesterol           (1696.47 MG)
0005. Egg yolk                                           (Dairy and Egg Products        ) --  Cholesterol           (1000.13 MG)
0006. Egg, yolk, raw, frozen, pasteurized                (Dairy and Egg Products        ) --  Cholesterol           (1000.00 MG)
0007. Whole eggs                                         (Dairy and Egg Products        ) --  Cholesterol           (420.21 MG)
0008. Egg, whole, raw, frozen, pasteurized               (Dairy and Egg Products        ) --  Cholesterol           (420.00 MG)
0009. Eggs, whole                                        (Dairy and Egg Products        ) --  Cholesterol           (411.33 MG)
0010. Eggs, Grade A, Large, egg whole                    (Dairy and Egg Products        ) --  Cholesterol           (411.00 MG)
0011. Butter, stick, salted                              (Dairy and Egg Products        ) --  Cholesterol           (235.00 MG)
0012. BUTTER, STICK, SALTED                              (Dairy and Egg Products        ) --  Cholesterol           (234.62 MG)
0013. Butter, stick, unsalted                            (Dairy and Egg Products        ) --  Cholesterol           (234.00 MG)
0014. BUTTER, STICK, UNSALTED                            (Dairy and Egg Products        ) --  Cholesterol           (233.57 MG)

... [lots of rows omitted] ...

0288. Yogurt, Greek, plain, nonfat                       (Dairy and Egg Products        ) --  Cholesterol           (5.00 MG)
0289. Cholesterol, Greek yogurt, CHOBANI PLAIN NON-FAT   (Dairy and Egg Products        ) --  Cholesterol           (5.00 MG)
0290. Cholesterol, Greek yogurt, FAGE PLAIN NONFAT (CA2  (Dairy and Egg Products        ) --  Cholesterol           (5.00 MG)
0291. Cholesterol, Greek yogurt, CHOBANI PLAIN NON-FAT   (Dairy and Egg Products        ) --  Cholesterol           (5.00 MG)
0292. Cholesterol, Greek yogurt, CHOBANI PLAIN NON-FAT   (Dairy and Egg Products        ) --  Cholesterol           (5.00 MG)
0293. Milk, lowfat, fluid, 1% milkfat, with added vitam  (Dairy and Egg Products        ) --  Cholesterol           (5.00 MG)
0294. Cholesterol, Greek yogurt, FAGE PLAIN NONFAT (AL1  (Dairy and Egg Products        ) --  Cholesterol           (5.00 MG)
0295. Cholesterol, Greek yogurt, FAGE PLAIN NONFAT (CO1  (Dairy and Egg Products        ) --  Cholesterol           (5.00 MG)
0296. Cholesterol, Yogurt, Greek, strawberry, non-fat,   (Dairy and Egg Products        ) --  Cholesterol           (4.00 MG)
0297. Yogurt, Greek, strawberry, nonfat                  (Dairy and Egg Products        ) --  Cholesterol           (4.00 MG)
0298. Milk, nonfat, fluid, with added vitamin A and vit  (Dairy and Egg Products        ) --  Cholesterol           (3.00 MG)
0299. Egg, white, raw, frozen, pasteurized               (Dairy and Egg Products        ) --  Cholesterol           (3.00 MG)
0300. Egg whites                                         (Dairy and Egg Products        ) --  Cholesterol           (2.75 MG)
0301. MILK, SKIM                                         (Dairy and Egg Products        ) --  Cholesterol           (2.48 MG)
```

Another example:

```bash
% go run cmd/cli/cli.go -n Protein

...

Top foods by nutrient 'Protein' (amount of the nutrient per 100g of food):

0001. Egg, white, dried                                  (Dairy and Egg Products        ) --  Protein               (79.90 G)
0002. Flour, soy, defatted                               (Legumes and Legume Products   ) --  Protein               (51.10 G)
0003. Egg, whole, dried                                  (Dairy and Egg Products        ) --  Protein               (48.10 G)
0004. Pork, cured, bacon, cooked, restaurant             (Pork Products                 ) --  Protein               (40.90 G)
0005. Flour, soy, full-fat                               (Legumes and Legume Products   ) --  Protein               (38.60 G)
0006. Egg, yolk, dried                                   (Dairy and Egg Products        ) --  Protein               (34.20 G)
0007. Chicken, broiler or fryers, breast, skinless, bon  (Poultry Products              ) --  Protein               (32.10 G)
0008. Beans, Dry, Great Northern, 446 (0% moisture)      (Legumes and Legume Products   ) --  Protein               (31.60 G)
0009. Beans, Dry, Pinto, 468 (0% moisture)               (Legumes and Legume Products   ) --  Protein               (30.80 G)
0010. Beans, Dry, Great Northern, 579 (0% moisture)      (Legumes and Legume Products   ) --  Protein               (30.50 G)
0011. Proximates, Beef, T-Bone Steak, lean only, cooked  (Beef Products                 ) --  Protein               (30.40 G)
0012. Proximates, Beef, T-Bone Steak, lean only, cooked  (Beef Products                 ) --  Protein               (29.90 G)

... [lots of rows omitted] ...

1088. Pears, raw, bartlett                               (Fruits and Fruit Juices       ) --  Protein               (0.38 G)
1089. Apples, granny smith, with skin, raw               (Fruits and Fruit Juices       ) --  Protein               (0.27 G)
1090. Apples, red delicious, with skin, raw              (Fruits and Fruit Juices       ) --  Protein               (0.19 G)
1091. Apples, fuji, with skin, raw                       (Fruits and Fruit Juices       ) --  Protein               (0.15 G)
1092. Apples, gala, with skin, raw                       (Fruits and Fruit Juices       ) --  Protein               (0.13 G)
1093. Apples, honeycrisp, with skin, raw                 (Fruits and Fruit Juices       ) --  Protein               (0.10 G)
1094. Oil, coconut                                       (Fats and Oils                 ) --  Protein               (0.00 G)
1095. Sugars, granulated                                 (Sweets                        ) --  Protein               (0.00 G)
```
