import 'package:flutter/material.dart';

class UseMatrix401Widget extends StatelessWidget {
  const UseMatrix401Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.black,
      child: Transform(
        alignment: Alignment.topRight, // 相对于坐标系原点的对齐方式
        transform: Matrix4.skewY(0.3), // 沿y轴倾斜0.3弧度
        child: Container(
          padding: const EdgeInsets.all(8.0),
          color: Colors.deepOrange,
          child: const Text('Apartment for rent!'),
        ),
      ),
    );
  }
}
