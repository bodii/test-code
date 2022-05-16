import 'package:flutter/material.dart';

class UseGridView01Widgets extends StatelessWidget {
  const UseGridView01Widgets({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return GridView(
      gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
          crossAxisCount: 3, // 横轴三个子widget
          childAspectRatio: 1.0 // 宽高比为1时，子widget
          ),
      children: const [
        Icon(Icons.ac_unit),
        Icon(Icons.airport_shuttle),
        Icon(Icons.all_inclusive),
        Icon(Icons.beach_access),
        Icon(Icons.cake),
        Icon(Icons.free_breakfast),
      ],
    );
  }
}
