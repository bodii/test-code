import 'package:flutter/material.dart';

import 'custom_tap_gesture_recognizer.dart';

class UseCustomGestureDetectorWidget extends StatefulWidget {
  const UseCustomGestureDetectorWidget({Key? key}) : super(key: key);

  @override
  _UseCustomGestureDetectorWidgetState createState() =>
      _UseCustomGestureDetectorWidgetState();
}

class _UseCustomGestureDetectorWidgetState
    extends State<UseCustomGestureDetectorWidget> {
  @override
  Widget build(BuildContext context) {
    return customGestureDetector(
      onTap: () => debugPrint('2'),
      child: Container(
        width: 200,
        height: 200,
        color: Colors.red,
        alignment: Alignment.center,
        child: GestureDetector(
          onTap: () => debugPrint('1'),
          child: Container(
            width: 50,
            height: 50,
            color: Colors.grey,
          ),
        ),
      ),
    );
  }
}
