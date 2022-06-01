import 'package:flutter/material.dart';

import 'widgets/use_animated_decorated_box_widget.dart';
import 'widgets/use_animateds_widgets.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('自定义动画过渡组件')),
      body: const UseAnimatedDecoratedBoxWidget(),
      // body: const UseAnimatedsWidgets(),
    );
  }
}
