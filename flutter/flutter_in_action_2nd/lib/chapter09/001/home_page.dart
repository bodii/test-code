import 'package:flutter/material.dart';

import 'widgets/scale_animation_widget.dart';
import 'widgets/scale_animation_02_widget.dart';
import 'widgets/scale_animation_03_widget.dart';
import 'widgets/scale_animation_04_widget.dart';
import 'widgets/grow_transition_widget.dart';
import 'widgets/animation_state_listener_demo_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Animation 动画基本结构及状态监听')),
      body: const AnimationStateListenerDemowidget(),
    );
  }
}
