import 'package:flutter/material.dart';

class UseRotatedBox01Widget extends StatelessWidget {
  const UseRotatedBox01Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.center,
      children: const [
        DecoratedBox(
          decoration: BoxDecoration(color: Colors.red),
          child: RotatedBox(
            quarterTurns: 1, // 旋转90度（1/4圈)
            child: Text('Hello world'),
          ),
        ),
        Text(
          '你好',
          style: TextStyle(color: Colors.green, fontSize: 18.0),
        ),
      ],
    );
  }
}
