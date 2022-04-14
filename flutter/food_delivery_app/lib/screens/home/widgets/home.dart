import 'package:flutter/material.dart';
import 'package:food_delivery_app/configs/colors.dart';
import 'package:food_delivery_app/models/restaurant.dart';
import 'package:food_delivery_app/widgets/float_shopping_bag.dart';
import 'package:food_delivery_app/widgets/food_list.dart';
import 'package:food_delivery_app/widgets/food_list_view.dart';
import 'package:food_delivery_app/widgets/home_app_bar.dart';
import 'package:food_delivery_app/widgets/restaurant_info.dart';
import 'package:smooth_page_indicator/smooth_page_indicator.dart';

class HomePage extends StatefulWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  int selected = 0;
  final pageController = PageController();
  final restaurant = Restaurant.generateRecommendRestaurant();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const HomeAppBarWidget(),
          RestaurantInfo(),
          FoodList(selected, (int index) {
            setState(() {
              selected = index;
            });
            pageController.jumpToPage(index);
          }, restaurant),
          Expanded(
            child: FoodListView(
              selected,
              (int index) {
                selected = index;
              },
              pageController,
              restaurant,
            ),
          ),
          Container(
            padding: const EdgeInsets.symmetric(horizontal: 25),
            height: 60,
            child: SmoothPageIndicator(
              controller: pageController,
              count: restaurant.menu.length,
              effect: CustomizableEffect(
                dotDecoration: DotDecoration(
                  width: 8,
                  height: 8,
                  color: Colors.grey.withOpacity(0.5),
                  borderRadius: BorderRadius.circular(8),
                ),
                activeDotDecoration: DotDecoration(
                  width: 10,
                  height: 10,
                  color: kPrimaryColor,
                  borderRadius: BorderRadius.circular(10),
                  dotBorder: const DotBorder(
                    color: kPrimaryColor,
                    padding: 2,
                    width: 2,
                  ),
                ),
              ),
              onDotClicked: (int index) {
                setState(() {
                  selected = index;
                });
                pageController.jumpToPage(index);
              },
            ),
          ),
        ],
      ),
      floatingActionButton: const FloatShoppingBag(),
    );
  }
}
