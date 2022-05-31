import 'package:flutter/material.dart';

import 'stagger_animation.dart';

class StaggerRoute extends StatefulWidget {
  const StaggerRoute({Key? key}) : super(key: key);

  @override
  _StaggerRouteState createState() => _StaggerRouteState();
}

class _StaggerRouteState extends State<StaggerRoute>
    with TickerProviderStateMixin {
  late AnimationController _controller;

  @override
  void initState() {
    super.initState();

    _controller = AnimationController(
      duration: const Duration(milliseconds: 2000),
      vsync: this,
    );
  }

  _playAnimation() async {
    try {
      // 先正向执行动画
      await _controller.forward().orCancel;
      // 再反向执行动画
      await _controller.reverse().orCancel;
      // ignore: empty_catches
    } on TickerCanceled {}
  }

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          ElevatedButton(
            child: const Text('start animation'),
            onPressed: () => _playAnimation(),
          ),
          Container(
            width: 300.0,
            height: 300.0,
            decoration: BoxDecoration(
              color: Colors.black.withOpacity(0.1),
              border: Border.all(
                color: Colors.black.withOpacity(.5),
              ),
            ),
            child: StaggerAnimation(controller: _controller),
          ),
        ],
      ),
    );
  }
}
