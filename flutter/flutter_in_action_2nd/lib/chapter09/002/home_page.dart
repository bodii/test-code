import 'package:flutter/material.dart';

import 'widgets/custom_route_animation_widget.dart';
import 'widgets/use_custom_fade_route_animation_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('自定义路由切换动画')),
      body: const UseCustomFadeRouteAnimationWidget(),
    );
  }
}
