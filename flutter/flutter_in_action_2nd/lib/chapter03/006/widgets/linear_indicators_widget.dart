import 'package:flutter/material.dart';

class LinearIndicatorswidget extends StatelessWidget {
  const LinearIndicatorswidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        LinearProgressIndicator(
          backgroundColor: Colors.grey[200],
          valueColor: const AlwaysStoppedAnimation(Colors.blue),
        ),
        LinearProgressIndicator(
          backgroundColor: Colors.grey[200],
          valueColor: const AlwaysStoppedAnimation(Colors.blue),
          value: .5,
        ),
      ].map((e) {
        return Container(
          width: 200,
          height: 12,
          child: e,
          margin: const EdgeInsets.symmetric(vertical: 20),
        );
      }).toList(),
    );
  }
}
