import 'package:flutter/material.dart';

import 'gradient_button.dart';

class UseGradientButtonWidget extends StatefulWidget {
  const UseGradientButtonWidget({Key? key}) : super(key: key);

  @override
  _UseGradientButtonWidgetState createState() =>
      _UseGradientButtonWidgetState();
}

class _UseGradientButtonWidgetState extends State<UseGradientButtonWidget> {
  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        GradientButton(
          child: const Text('Submit'),
          colors: const [Colors.orange, Colors.red],
          height: 50.0,
          onPressed: onTap,
        ),
        GradientButton(
          child: const Text('Submit'),
          colors: [Colors.lightGreen, Colors.green[700] as Color],
          height: 50.0,
          onPressed: onTap,
        ),
        GradientButton(
          child: const Text('Submit'),
          height: 50.0,
          colors: [Colors.lightBlue[300] as Color, Colors.blueAccent],
          onPressed: onTap,
        ),
      ],
    );
  }

  void onTap() {
    debugPrint('button click');
  }
}
