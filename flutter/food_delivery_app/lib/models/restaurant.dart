import 'package:food_delivery_app/models/food.dart';

class Restaurant {
  String name;
  String waitTime;
  String distance;
  String label;
  String logoUrl;
  String desc;
  num score;
  Map<String, List<Food>> menu;

  Restaurant(
    this.name,
    this.waitTime,
    this.distance,
    this.label,
    this.logoUrl,
    this.desc,
    this.score,
    this.menu,
  );

  static Restaurant generateRecommendRestaurant() {
    return Restaurant(
      'Golden pepper fish restaurant',
      '10 min',
      '1km',
      'Chinese food',
      'assets/images/Golden_pepper_fish_restaurant.jpg',
      'Golden pepper fish restaurant',
      4.0,
      {
        'Specialty': Food.generateRecommendFoods(),
        'Stir Fry': Food.generateStirFry(),
        'Cold Dishes': Food.generateColdDishes(),
        'Staple Food': Food.generateStapleFoods(),
        'Soup': Food.generateSoups(),
      },
    );
  }
}
