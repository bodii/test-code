import 'package:flutter/material.dart';

import 'widgets/animated_switcher_counter_widget.dart';
import 'widgets/custom_slide_transition_widget.dart';
import 'widgets/use_slide_transition_x_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('动画切换')),
      body: const UseSlideTransitionXWidget(),
    );
  }
}
