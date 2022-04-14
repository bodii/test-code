import 'package:flutter/material.dart';
import 'package:food_delivery_app/models/food.dart';
import 'package:food_delivery_app/widgets/float_shopping_bag.dart';
import 'package:food_delivery_app/widgets/food_detail_app_bar.dart';
import 'package:food_delivery_app/widgets/food_nums.dart';

class DetailPage extends StatefulWidget {
  Food food;
  DetailPage(this.food, {Key? key}) : super(key: key);

  @override
  State<DetailPage> createState() => _DetailPageState();
}

class _DetailPageState extends State<DetailPage> {
  @override
  Widget build(BuildContext context) {
    final _food = widget.food;

    return Scaffold(
      // backgroundColor: Colors.grey.withOpacity(0.1),
      body: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Container(
            width: double.infinity,
            height: 265,
            decoration: const BoxDecoration(
              image: DecorationImage(
                image: AssetImage('assets/images/header01.jpg'),
                fit: BoxFit.cover,
              ),
            ),
            child: Column(
              children: [
                const FoodDetailAppBarWidget(),
                const SizedBox(height: 25),
                Stack(
                  alignment: Alignment.bottomCenter,
                  children: [
                    Column(children: [
                      const SizedBox(
                        height: 80,
                      ),
                      Container(
                        width: double.infinity,
                        height: 80,
                        decoration: const BoxDecoration(
                          color: Color.fromARGB(250, 250, 250, 250),
                          borderRadius:
                              BorderRadius.vertical(top: Radius.circular(40)),
                        ),
                      )
                    ]),
                    Container(
                      width: 160,
                      height: 160,
                      alignment: Alignment.center,
                      decoration: BoxDecoration(
                        borderRadius: BorderRadius.circular(160),
                        boxShadow: [
                          BoxShadow(
                            color: Colors.grey.withOpacity(0.1),
                            offset: const Offset(0.0, 0.5),
                            blurRadius: 2,
                            spreadRadius: 0.2,
                          ),
                        ],
                      ),
                      child: ClipPath.shape(
                        shape: const CircleBorder(),
                        child: Align(
                          alignment: Alignment.topCenter,
                          heightFactor: 1.0,
                          child: Image.asset(
                            _food.imgUrl,
                            fit: BoxFit.cover,
                          ),
                        ),
                      ),
                    ),
                  ],
                ),
              ],
            ),
          ),
          Expanded(
            child: ListView(
              children: [
                Container(
                  alignment: Alignment.center,
                  height: 80,
                  child: Text(
                    _food.name,
                    style: const TextStyle(
                      fontWeight: FontWeight.bold,
                      fontSize: 23.0,
                    ),
                    overflow: TextOverflow.fade,
                    textAlign: TextAlign.center,
                  ),
                ),
                Row(
                  // 等待时间、评分、热量
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Container(
                      padding: const EdgeInsets.all(1),
                      alignment: Alignment.center,
                      child: Row(
                        children: [
                          Icon(
                            Icons.query_builder_outlined,
                            color: Colors.blue[400],
                          ),
                          Text(
                            _food.waitTime,
                            style: const TextStyle(
                              fontSize: 15,
                            ),
                          ),
                        ],
                      ),
                    ),
                    Container(
                      padding: const EdgeInsets.symmetric(horizontal: 25),
                      child: Row(
                        children: [
                          Icon(
                            Icons.star_outline_outlined,
                            color: Colors.orangeAccent[100],
                          ),
                          Text(
                            _food.score.toString(),
                            style: const TextStyle(
                              fontSize: 15,
                            ),
                          ),
                        ],
                      ),
                    ),
                    Container(
                      padding: const EdgeInsets.symmetric(horizontal: 25),
                      child: Row(
                        children: [
                          Icon(
                            Icons.whatshot_outlined,
                            color: Colors.red[400],
                          ),
                          Text(
                            _food.cal,
                            style: const TextStyle(
                              fontSize: 15,
                            ),
                          ),
                        ],
                      ),
                    ),
                  ],
                ),
                Container(
                  // 价格和数量
                  width: 200,
                  height: 50,
                  alignment: Alignment.center,
                  margin: const EdgeInsets.symmetric(vertical: 20),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      Stack(
                        alignment: Alignment.centerLeft,
                        children: [
                          Container(
                            width: 100,
                            height: 36,
                            decoration: const BoxDecoration(
                              borderRadius: BorderRadius.horizontal(
                                left: Radius.circular(30),
                              ),
                              color: Color.fromARGB(255, 244, 244, 245),
                            ),
                          ),
                          Container(
                            margin: const EdgeInsets.only(left: 20),
                            child: Row(
                              crossAxisAlignment: CrossAxisAlignment.start,
                              children: [
                                const Text(
                                  '\$',
                                  style: TextStyle(
                                    fontSize: 13,
                                    fontWeight: FontWeight.bold,
                                  ),
                                ),
                                Text(
                                  _food.price.toString(),
                                  style: const TextStyle(
                                    fontSize: 26,
                                    fontWeight: FontWeight.w900,
                                  ),
                                ),
                              ],
                            ),
                          ),
                          const FoodNumsWidget(),
                        ],
                      ),
                    ],
                  ),
                ),
                Container(
                  // 成分
                  width: double.infinity,
                  height: 165,
                  alignment: Alignment.topLeft,
                  child: Column(
                    children: [
                      Container(
                        alignment: Alignment.topLeft,
                        padding: const EdgeInsets.only(left: 30, bottom: 12),
                        child: const Text(
                          'Ingredients',
                          style: TextStyle(
                            fontSize: 18.0,
                            fontWeight: FontWeight.w900,
                          ),
                        ),
                      ),
                      Row(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: _buildIngredients(_food.ingredients),
                      ),
                    ],
                  ),
                ),
                // 关于
                Container(
                  margin: const EdgeInsets.fromLTRB(30, 30, 30, 100),
                  child: Column(
                    mainAxisSize: MainAxisSize.min,
                    crossAxisAlignment: CrossAxisAlignment.start,
                    children: [
                      const Text(
                        'About',
                        style: TextStyle(
                          fontWeight: FontWeight.bold,
                          fontSize: 18,
                        ),
                      ),
                      const SizedBox(height: 10),
                      Text(
                        _food.about,
                        softWrap: true,
                        style: TextStyle(
                          fontWeight: FontWeight.w300,
                          fontSize: 15,
                          color: Colors.grey[500],
                        ),
                      ),
                    ],
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
      floatingActionButton: const FloatShoppingBag(),
    );
  }
}

List<Widget> _buildIngredients(List<Map<String, String>> ingredients) {
  var list = ingredients.asMap().map(
      (i, v) => MapEntry(i, _getIngredientsInfo(v.values.first, v.keys.first)));

  return list.values.toList();
}

Widget _getIngredientsInfo(String imagePath, String name) {
  return Container(
    alignment: Alignment.center,
    width: 68,
    height: 105,
    margin: const EdgeInsets.symmetric(
      horizontal: 5,
    ),
    decoration: BoxDecoration(
      borderRadius: BorderRadius.circular(60),
      color: Colors.white70,
    ),
    child: Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        SizedBox(
          width: 40,
          height: 40,
          child: ClipRRect(
            child: Image.asset(
              imagePath,
              fit: BoxFit.cover,
            ),
            borderRadius: BorderRadius.circular(120),
          ),
        ),
        const SizedBox(
          height: 10,
        ),
        Text(
          name,
          textAlign: TextAlign.center,
          style: const TextStyle(
            fontWeight: FontWeight.bold,
            fontSize: 13,
          ),
        ),
      ],
    ),
  );
}
