import 'package:flutter/material.dart';
import 'dart:math';

class UseTransformRotate01Widget extends StatelessWidget {
  const UseTransformRotate01Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return DecoratedBox(
      decoration: const BoxDecoration(color: Colors.red),
      child: Transform.rotate(
        // 旋转90度
        angle: pi / 2,
        child: const Text('Hello world'),
      ),
    );
  }
}
