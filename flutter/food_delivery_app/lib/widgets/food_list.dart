import 'package:flutter/material.dart';
import 'package:food_delivery_app/configs/colors.dart';
import 'package:food_delivery_app/models/restaurant.dart';

class FoodList extends StatelessWidget {
  final int selected;
  final Function callback;
  final Restaurant restaurant;

  const FoodList(this.selected, this.callback, this.restaurant, {Key? key})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    final category = restaurant.menu.keys.toList();
    return Container(
      height: 70,
      padding: const EdgeInsets.symmetric(vertical: 15),
      child: ListView.separated(
        padding: const EdgeInsets.symmetric(horizontal: 25),
        scrollDirection: Axis.horizontal,
        itemBuilder: (context, index) => GestureDetector(
          onTap: () => callback(index),
          child: Container(
            padding: const EdgeInsets.symmetric(vertical: 10, horizontal: 14),
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(20),
              color: selected == index ? kPrimaryColor : Colors.grey[100],
            ),
            child: Text(
              category[index],
              style: const TextStyle(
                fontWeight: FontWeight.bold,
              ),
            ),
          ),
        ),
        separatorBuilder: (_, index) => const SizedBox(
          width: 20,
        ),
        itemCount: category.length,
      ),
    );
  }
}
