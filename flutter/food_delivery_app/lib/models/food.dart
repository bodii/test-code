class Food {
  String imgUrl;
  String desc;
  String name;
  String waitTime;
  num score;
  String cal;
  num price;
  num guntity;
  List<Map<String, String>> ingredients;
  String about;
  bool hightLight;

  Food(this.imgUrl, this.desc, this.name, this.waitTime, this.score, this.cal,
      this.price, this.guntity, this.ingredients, this.about,
      {this.hightLight = false});

  static List<Food> generateRecommendFoods() {
    List<Food> stirFrys = generateStirFry();
    List<Food> coldDishes = generateColdDishes();

    return [
      stirFrys[0],
      stirFrys[1],
      coldDishes[0],
    ];
  }

  /// Stir Fry
  static List<Food> generateStirFry() {
    return [
      Food(
        'assets/images/Yu-Shiang_Shredded_Pork.jpg',
        'No1. in Sales',
        'Yu-Shiang Shredded Pork',
        '15 min',
        4.8,
        '571 kcal',
        22,
        1,
        [
          {'Pork': 'assets/images/pork.jpg'},
          {'Agaric': 'assets/images/agaric.jpg'},
          {'Carrot': 'assets/images/carrot.jpg'},
          {'Green Pepper': 'assets/images/green_pepper.jpg'},
        ],
        "Shredded pork with fish flavor is a common Sichuan dish. "
            "Fish flavor is one of the main traditional flavors of "
            "Sichuan cuisine. The finished dish has the smell of fish, "
            "and its flavor is modulated by condiments. This method "
            "originated from the unique fish cooking and seasoning "
            "method of Sichuan folk, and now it has become a taste "
            "type loved by the public.",
      ),
      Food(
        'assets/images/scrambled_egg_with_tomato.jpg',
        'No2. in Sales',
        'Scrambled Egg With Tomato',
        '10 min',
        4.1,
        '94 kcal',
        18,
        1,
        [
          {'Egg': 'assets/images/egg.jpg'},
          {'Tomato': 'assets/images/tomato.jpg'},
          {'Scallion': 'assets/images/scallion.jpg'},
        ],
        "Scrambled egg with tomato is a home cooked Shandong dish with tomato "
            "and egg as the main material. It tastes sour, sweet and delicious, "
            "especially after dinner.",
      ),
    ];
  }

  /// Cold Dishes
  static List<Food> generateColdDishes() {
    return [
      Food(
        'assets/images/Cold_Lotus_Root_Slices.jpg',
        'No3. in Sales',
        'Cold Lotus Root Slices ',
        '5 min',
        4.3,
        '70 kcal',
        16,
        1,
        [
          {'Lotus Root': 'assets/images/Lotus_Root.jpg'},
        ],
        "Lotus root has high iron content and contains a lot of vitamin C "
            "and dietary fiber. Raw food can produce saliva and quench thirst, "
            "clear away heat and annoyance, cooked food can nourish the stomach "
            "and eliminate food, tonify the heart and stop diarrhea. Eating in "
            "summer can supplement nutrition and cure fever and thirst",
      ),
    ];
  }

  /// Staple Food
  static List<Food> generateStapleFoods() {
    return [
      Food(
        'assets/images/Chinese_Hamburger.jpg',
        'No4. in Sales',
        'Chinese Hamburger',
        '12 min',
        4.5,
        '659 kcal',
        8,
        1,
        [
          {'Streaky Pork': 'assets/images/Streaky_Pork.jpg'},
          {'Clay Oven Rolls': 'assets/images/Clay_Oven_Rolls.jpg'},
        ],
        "Roujiamo is one of the traditional specialty foods in "
            "Shaanxi Province, China. The name means "
            "steamed bun "
            "with meat filling",
      ),
    ];
  }

  /// Soup
  static List<Food> generateSoups() {
    return [
      Food(
        'assets/images/White_Gourd_Meatball_Soup.jpg',
        'No5. in Sales',
        'White Gourd Meatball Soup',
        '18 min',
        4.3,
        '571 kcal',
        19,
        1,
        [
          {'Meatballs': 'assets/images/Meatballs.jpg'},
          {'White Gourd': 'assets/images/White_Gourd.jpg'},
        ],
        "White gourd meatball soup is a dish. The main raw "
            "materials are white gourd, meatballs and so on. "
            "The soup tastes delicious, the meatballs are "
            "smooth and nutritious. It is a soup suitable for all ages.",
      ),
    ];
  }
}
