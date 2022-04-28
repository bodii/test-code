import 'package:flutter/material.dart';

class UseTransformScale01Widget extends StatelessWidget {
  const UseTransformScale01Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return DecoratedBox(
      decoration: const BoxDecoration(color: Colors.red),
      child: Transform.scale(
        scale: 1.5,
        child: const Text('Hello world'),
      ),
    );
  }
}
