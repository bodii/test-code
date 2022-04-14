import 'package:flutter/material.dart';
import 'package:food_delivery_app/configs/colors.dart';
import 'package:food_delivery_app/controllers/food_nums_controller.dart';
import 'package:get/get.dart';

class FloatShoppingBag extends StatefulWidget {
  const FloatShoppingBag({Key? key}) : super(key: key);

  @override
  _FloatShoppingBagState createState() => _FloatShoppingBagState();
}

class _FloatShoppingBagState extends State<FloatShoppingBag> {
  bool visible = true;
  final foodNumsLogic = Get.put(FoodNumsController());

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.end,
      children: [
        Offstage(
          offstage: !visible,
          child: Container(
            height: 56,
            width: 56,
            padding: const EdgeInsets.all(7.5),
            alignment: Alignment.center,
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(30),
              color: kPrimaryColor,
            ),
            child: IconButton(
              padding: const EdgeInsets.all(0),
              icon: const Icon(
                Icons.shopping_bag_outlined,
                size: 30,
              ),
              onPressed: () {
                setState(() {
                  visible = !visible;
                });
              },
              highlightColor: Colors.transparent.withOpacity(0),
              splashColor: Colors.transparent.withOpacity(0),
              focusColor: Colors.transparent.withOpacity(0),
              hoverColor: Colors.transparent.withOpacity(0),
            ),
          ),
        ),
        Offstage(
          offstage: visible,
          child: Container(
            height: 50,
            width: 92,
            padding: const EdgeInsets.symmetric(horizontal: 6),
            decoration: BoxDecoration(
              borderRadius: BorderRadius.circular(30),
              color: kPrimaryColor,
            ),
            child: Row(
              children: [
                IconButton(
                  padding: const EdgeInsets.all(0),
                  alignment: Alignment.centerLeft,
                  icon: const Icon(
                    Icons.shopping_bag_outlined,
                    size: 30,
                  ),
                  onPressed: () {
                    setState(() {
                      visible = !visible;
                    });
                  },
                  highlightColor: Colors.transparent.withOpacity(0),
                  splashColor: Colors.transparent.withOpacity(0),
                  focusColor: Colors.transparent.withOpacity(0),
                  hoverColor: Colors.transparent.withOpacity(0),
                ),
                Container(
                  alignment: Alignment.center,
                  width: 30,
                  height: 30,
                  decoration: BoxDecoration(
                    borderRadius: BorderRadius.circular(30),
                    color: Colors.white,
                  ),
                  child: Obx(() {
                    return Text(
                      '${foodNumsLogic.nums}',
                      textAlign: TextAlign.center,
                      style: const TextStyle(
                        color: Colors.black,
                        fontSize: 18,
                        fontWeight: FontWeight.bold,
                      ),
                    );
                  }),
                ),
              ],
            ),
          ),
        ),
      ],
    );
  }
}
