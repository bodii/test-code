import 'package:flutter/material.dart';

class UseContainer02Widget extends StatelessWidget {
  const UseContainer02Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.only(top: 70),
      child: Column(
        children: [
          Container(
            margin: const EdgeInsets.all(20.0),
            color: Colors.orange,
            child: const Text('Hello world!'),
          ),
          const Padding(
            padding: EdgeInsets.all(20.0),
            child: DecoratedBox(
              decoration: BoxDecoration(color: Colors.orange),
              child: Text('Hello world!'),
            ),
          ),
          Container(
            padding: const EdgeInsets.all(20.0),
            color: Colors.orange,
            child: const Text('Hello wolrd!'),
          ),
          const DecoratedBox(
            decoration: BoxDecoration(
              color: Colors.orange,
            ),
            child: Padding(
              padding: EdgeInsets.all(20.0),
              child: Text('Hello world!'),
            ),
          ),
        ],
      ),
    );
  }
}
