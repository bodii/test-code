import 'package:flutter/material.dart';

import 'widgets/use_gesture_widget.dart';
import 'widgets/use_gesture_detector_02_widget.dart';
import 'widgets/use_gesture_detector_drag_widget.dart';
import 'widgets/use_gesture_detector_scale_widget.dart';
import 'widgets/use_gesture_recognizer_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('手势识别')),
      body: const UseGestureRecognizerWidget(),
    );
  }
}
