import 'package:flutter/material.dart';
import 'package:food_delivery_app/controllers/favorite_controller.dart';
import 'package:food_delivery_app/routers/routers.dart';
import 'package:get/get.dart';

class FoodDetailAppBarWidget extends StatefulWidget {
  const FoodDetailAppBarWidget({
    Key? key,
  }) : super(key: key);

  @override
  State<FoodDetailAppBarWidget> createState() => _FoodDetailAppBarWidgetState();
}

class _FoodDetailAppBarWidgetState extends State<FoodDetailAppBarWidget> {
  final favorite = Get.put(FavoriteController());

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.only(top: 25),
      padding: const EdgeInsets.only(
        top: 15,
        left: 25,
        right: 25,
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          GestureDetector(
            child: Container(
              padding: const EdgeInsets.all(8),
              decoration: BoxDecoration(
                shape: BoxShape.circle,
                color: Colors.grey[100],
              ),
              child: const Icon(Icons.arrow_back),
            ),
            onTap: () {
              Navigator.pushNamed(context, homePage);
            },
          ),
          GestureDetector(
            child: Container(
              padding: const EdgeInsets.all(8),
              decoration: BoxDecoration(
                shape: BoxShape.circle,
                color: Colors.grey[100],
              ),
              child: Column(
                children: [
                  Offstage(
                    child: const Icon(Icons.favorite_border_outlined),
                    offstage: favorite.getVal(),
                  ),
                  Offstage(
                    child: const Icon(
                      Icons.favorite,
                      color: Colors.pink,
                    ),
                    offstage: !favorite.getVal(),
                  ),
                ],
              ),
            ),
            onTap: () {
              setState(() {
                favorite.change();
              });
            },
          ),
        ],
      ),
    );
  }
}
