import 'package:flutter/material.dart';
import 'package:food_delivery_app/configs/colors.dart';
import 'package:food_delivery_app/models/food.dart';
import 'package:food_delivery_app/screens/detail/widgets/detail.dart';

class FoodItem extends StatelessWidget {
  final Food food;

  const FoodItem(this.food, {Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return GestureDetector(
      onTap: () {
        Navigator.push(
          context,
          MaterialPageRoute(builder: (BuildContext context) {
            return DetailPage(food);
          }),
        );
        // Navigator.pushNamed(context, '/detail');
        // Navigator.push(
        //   context,
        //   MaterialPageRoute(builder: (context) => const DetailPage()),
        // );
      },
      child: Container(
        height: 110,
        decoration: BoxDecoration(
          color: Colors.grey[50],
          borderRadius: BorderRadius.circular(20),
        ),
        child: Row(
          children: [
            Container(
              padding: const EdgeInsets.all(5),
              width: 110,
              height: 110,
              child: ClipRRect(
                child: Image.asset(
                  food.imgUrl,
                  fit: BoxFit.cover,
                ),
                borderRadius: BorderRadius.circular(110),
              ),
            ),
            Expanded(
              child: Container(
                padding: const EdgeInsets.only(
                  top: 20,
                  left: 10,
                  right: 10,
                ),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Row(
                      mainAxisAlignment: MainAxisAlignment.spaceBetween,
                      children: [
                        SizedBox(
                          width: MediaQuery.of(context).size.width * 0.45,
                          child: Text(
                            food.name,
                            style: const TextStyle(
                              fontSize: 16,
                              fontWeight: FontWeight.bold,
                              height: 1.5,
                            ),
                            overflow: TextOverflow.ellipsis,
                          ),
                        ),
                        const Icon(Icons.arrow_forward_ios_outlined, size: 15),
                      ],
                    ),
                    Text(
                      food.desc,
                      style: TextStyle(
                        color: food.hightLight
                            ? kPrimaryColor
                            : Colors.grey.withOpacity(0.8),
                        height: 1.5,
                      ),
                    ),
                    const SizedBox(
                      height: 5,
                    ),
                    Row(
                      children: [
                        const Text(
                          '\$',
                          style: TextStyle(
                            fontSize: 10,
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                        Text(
                          '${food.price}',
                          style: const TextStyle(
                            fontSize: 18,
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                      ],
                    ),
                  ],
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
