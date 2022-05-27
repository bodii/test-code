import 'package:flutter/material.dart';

import 'widgets/both_direction_test_widget.dart';
import 'widgets/gesture_conflict_widget.dart';
import 'widgets/listenter_gesture_conflict_widget.dart';
import 'widgets/use_custom_gesture_detector_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('手势原理与手势冲突')),
      body: const UseCustomGestureDetectorWidget(),
    );
  }
}
