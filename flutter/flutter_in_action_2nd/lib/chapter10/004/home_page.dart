import 'package:flutter/material.dart';

import 'widgets/use_gradient_circular_progress_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('自绘:圆形背景渐变进度条')),
      body: const UseGradientCircularProgressWidget(),
    );
  }
}
