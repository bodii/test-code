import 'package:flutter/material.dart';

class UseTransformTranslate01Widget extends StatelessWidget {
  const UseTransformTranslate01Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return DecoratedBox(
      decoration: const BoxDecoration(
        color: Colors.red,
      ),
      // 默认原点为左上角，左移20像素，向上平移5像素
      child: Transform.translate(
        offset: const Offset(-20.0, -5.0),
        child: const Text('Hello world'),
      ),
    );
  }
}
