import 'package:flutter/material.dart';
import 'package:food_delivery_app/controllers/food_nums_controller.dart';
import 'package:get/get.dart';

class FoodNumsWidget extends StatefulWidget {
  const FoodNumsWidget({
    Key? key,
  }) : super(key: key);

  @override
  State<FoodNumsWidget> createState() => _FoodNumsWidgetState();
}

class _FoodNumsWidgetState extends State<FoodNumsWidget> {
  final foodNumsLogic = Get.put(FoodNumsController());

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.only(left: 70),
      width: 122,
      height: 36,
      decoration: BoxDecoration(
        borderRadius: BorderRadius.circular(30),
        color: Colors.orange[300],
      ),
      child: Row(
        mainAxisAlignment: MainAxisAlignment.spaceBetween,
        children: [
          ChangeFoodNumsWidget(
            icon: Icons.remove,
            event: () {
              setState(() {
                foodNumsLogic.decrease();
              });
            },
          ),
          Container(
            constraints: const BoxConstraints(minWidth: 24),
            padding: const EdgeInsets.all(3),
            child: Obx(() {
              return Text(
                '${foodNumsLogic.nums}',
                textAlign: TextAlign.center,
                maxLines: 1,
                style: const TextStyle(
                  fontSize: 15,
                  fontWeight: FontWeight.w900,
                  color: Colors.black,
                ),
              );
            }),
            decoration: const BoxDecoration(
              borderRadius: BorderRadius.all(Radius.circular(13)),
              color: Colors.white,
            ),
          ),
          ChangeFoodNumsWidget(
            icon: Icons.add,
            event: () {
              setState(() {
                foodNumsLogic.increase();
              });
            },
          ),
        ],
      ),
    );
  }
}

class ChangeFoodNumsWidget extends StatelessWidget {
  IconData icon;
  void Function()? event;

  ChangeFoodNumsWidget({Key? key, required this.icon, required this.event})
      : super(key: key);

  @override
  build(BuildContext context) {
    return IconButton(
      icon: Icon(icon),
      color: Colors.black,
      iconSize: 16,
      padding: EdgeInsets.zero,
      splashColor: Colors.transparent,
      highlightColor: Colors.transparent,
      hoverColor: Colors.transparent,
      onPressed: event,
    );
  }
}
