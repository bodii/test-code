import 'package:flutter/material.dart';
import 'package:food_delivery_app/models/restaurant.dart';
import 'package:food_delivery_app/widgets/food_item.dart';

class FoodListView extends StatelessWidget {
  final int selected;
  final Function callback;
  final PageController pageController;
  final Restaurant restaurant;
  const FoodListView(
      this.selected, this.callback, this.pageController, this.restaurant,
      {Key? key})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    final categroy = restaurant.menu.keys.toList();
    return Container(
      padding: const EdgeInsets.symmetric(horizontal: 25),
      child: PageView(
        controller: pageController,
        onPageChanged: (index) => callback(index),
        children: categroy
            .map((e) => ListView.separated(
                  itemBuilder: (context, index) =>
                      FoodItem(restaurant.menu[categroy[selected]]![index]),
                  separatorBuilder: (_, index) => const SizedBox(height: 15),
                  itemCount: restaurant.menu[categroy[selected]]!.length,
                ))
            .toList(),
      ),
    );
  }
}
