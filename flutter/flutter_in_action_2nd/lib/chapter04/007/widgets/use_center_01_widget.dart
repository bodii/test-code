import 'package:flutter/material.dart';

class UseCenter01Widget extends StatelessWidget {
  const UseCenter01Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: const [
        DecoratedBox(
          decoration: BoxDecoration(color: Colors.red),
          child: Center(
            child: Text('xxxx'),
          ),
        ),
        SizedBox(height: 10),
        DecoratedBox(
          decoration: BoxDecoration(color: Colors.red),
          child: Center(
            widthFactor: 1,
            heightFactor: 1,
            child: Text('xxxx'),
          ),
        ),
        SizedBox(height: 10),
        DecoratedBox(
          decoration: BoxDecoration(color: Colors.red),
          child: Center(
            widthFactor: 2.0,
            heightFactor: 2.0,
            child: Text('xxxx'),
          ),
        ),
      ],
    );
  }
}
