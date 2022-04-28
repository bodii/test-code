import 'package:flutter/material.dart';

class UseTransformScale02Widget extends StatelessWidget {
  const UseTransformScale02Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        DecoratedBox(
          decoration: const BoxDecoration(color: Colors.red),
          child: Transform.scale(
            scale: 1.5,
            child: const Text('Hello world'),
          ),
        ),
        const Text(
          '你好',
          style: TextStyle(color: Colors.green, fontSize: 18.0),
        ),
      ],
    );
  }
}
