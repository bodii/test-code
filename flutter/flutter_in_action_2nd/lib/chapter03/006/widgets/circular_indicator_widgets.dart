import 'package:flutter/material.dart';

class CircularIndicatorWidgets extends StatelessWidget {
  const CircularIndicatorWidgets({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        CircularProgressIndicator(
          backgroundColor: Colors.grey[200],
          valueColor: const AlwaysStoppedAnimation(Colors.blue),
        ),
        CircularProgressIndicator(
          backgroundColor: Colors.grey[200],
          valueColor: const AlwaysStoppedAnimation(Colors.blue),
          value: .5,
        ),
      ].map((e) {
        return Container(
          child: e,
          width: 50,
          height: 50,
          margin: const EdgeInsets.symmetric(vertical: 10),
        );
      }).toList(),
    );
  }
}
