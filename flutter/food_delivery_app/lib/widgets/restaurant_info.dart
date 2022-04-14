import 'package:flutter/material.dart';
import 'package:food_delivery_app/configs/colors.dart';
import 'package:food_delivery_app/models/restaurant.dart';

class RestaurantInfo extends StatelessWidget {
  final restaurant = Restaurant.generateRecommendRestaurant();

  RestaurantInfo({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.only(top: 30),
      padding: const EdgeInsets.symmetric(horizontal: 25),
      child: Column(
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  SizedBox(
                    width: MediaQuery.of(context).size.width * 0.6,
                    child: Text(
                      restaurant.name,
                      style: const TextStyle(
                        fontSize: 25,
                        fontWeight: FontWeight.bold,
                      ),
                      overflow: TextOverflow.fade,
                      maxLines: 2,
                      softWrap: true,
                    ),
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  Row(
                    children: [
                      Container(
                        padding: const EdgeInsets.all(5),
                        decoration: BoxDecoration(
                            color: Colors.grey.withOpacity(0.4),
                            borderRadius: BorderRadius.circular(5)),
                        child: Text(
                          restaurant.waitTime,
                          style: const TextStyle(
                            color: Colors.white,
                          ),
                        ),
                      ),
                      const SizedBox(width: 10),
                      Text(
                        restaurant.distance,
                        style: TextStyle(
                          fontSize: 16,
                          fontWeight: FontWeight.bold,
                          color: Colors.grey.withOpacity(0.4),
                        ),
                      ),
                      const SizedBox(
                        width: 10,
                      ),
                      Text(
                        restaurant.label,
                        style: TextStyle(
                          fontSize: 16,
                          fontWeight: FontWeight.bold,
                          color: Colors.grey.withOpacity(0.4),
                        ),
                      ),
                    ],
                  ),
                ],
              ),
              ClipRRect(
                borderRadius: BorderRadius.circular(50),
                child: Image.asset(restaurant.logoUrl, width: 80),
              ),
            ],
          ),
          const SizedBox(
            height: 5,
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceBetween,
            children: [
              SizedBox(
                width: MediaQuery.of(context).size.width * 0.71,
                child: Text(
                  '"${restaurant.desc}"',
                  style: const TextStyle(
                    fontSize: 16,
                  ),
                  overflow: TextOverflow.fade,
                ),
              ),
              Row(
                children: [
                  const Icon(
                    Icons.star_outline,
                    color: kPrimaryColor,
                  ),
                  Text(
                    '${restaurant.score}',
                    style: const TextStyle(
                      fontSize: 18,
                      fontWeight: FontWeight.bold,
                    ),
                  ),
                ],
              ),
            ],
          ),
        ],
      ),
    );
  }
}
