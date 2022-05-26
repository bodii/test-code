import 'package:flutter/material.dart';

import 'widgets/use_pointer_down_listener_widget.dart';
import 'widgets/water_mask_widget.dart';
import 'widgets/stack_event_widget.dart';
import 'widgets/hit_test_behavior_widget.dart';
import 'widgets/all_children_can_response_event_widget.dart';
import 'widgets/custom_hit_test_blacker_widget.dart';
import 'widgets/use_gesture_hit_test_block_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('事件处理流程demo')),
      body: const UseGestureHitTestBlockWidget(),
    );
  }
}
