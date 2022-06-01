import 'package:flutter/material.dart';

import 'animated_decorated_box.dart';
import 'animated_decorated_box_02.dart';

class UseAnimatedDecoratedBoxWidget extends StatefulWidget {
  const UseAnimatedDecoratedBoxWidget({Key? key}) : super(key: key);

  @override
  _UseAnimatedDecoratedBoxWidgetState createState() =>
      _UseAnimatedDecoratedBoxWidgetState();
}

class _UseAnimatedDecoratedBoxWidgetState
    extends State<UseAnimatedDecoratedBoxWidget> {
  Color _decorationColor = Colors.blue;
  Duration duration = const Duration(seconds: 1);

  @override
  Widget build(BuildContext context) {
    return Center(
      child: AnimatedDecoratedBox02(
        duration: duration,
        decoration: BoxDecoration(color: _decorationColor),
        child: TextButton(
          onPressed: () {
            setState(() {
              _decorationColor = Colors.red;
              debugPrint('click');
            });
          },
          child: const Text(
            'AnimatedDecoratedBox',
            style: TextStyle(color: Colors.white),
          ),
        ),
      ),
    );
  }
}
