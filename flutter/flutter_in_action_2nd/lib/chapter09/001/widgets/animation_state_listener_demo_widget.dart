import 'package:flutter/material.dart';

class GrowTransition extends StatelessWidget {
  const GrowTransition({
    Key? key,
    required this.animation,
    this.child,
  }) : super(key: key);

  final Widget? child;
  final Animation<double> animation;

  @override
  Widget build(BuildContext context) {
    return Center(
      child: AnimatedBuilder(
          animation: animation,
          child: child,
          builder: (BuildContext context, child) {
            return SizedBox(
              child: child,
              width: animation.value,
              height: animation.value,
            );
          }),
    );
  }
}

class AnimationStateListenerDemowidget extends StatefulWidget {
  const AnimationStateListenerDemowidget({Key? key}) : super(key: key);

  @override
  _AnimationStateListenerDemowidgetState createState() =>
      _AnimationStateListenerDemowidgetState();
}

class _AnimationStateListenerDemowidgetState
    extends State<AnimationStateListenerDemowidget>
    with SingleTickerProviderStateMixin {
  late Animation<double> animation;
  late AnimationController controller;

  @override
  void initState() {
    super.initState();
    controller = AnimationController(
      duration: const Duration(seconds: 3),
      vsync: this,
    );

    animation = Tween(begin: 0.0, end: 300.0).animate(controller);
    animation.addStatusListener((status) {
      if (status == AnimationStatus.completed) {
        // 如果执行完时反向执行动画
        controller.reverse();
      } else if (status == AnimationStatus.dismissed) {
        {
          // 如果动画状态回到初始时，开始正向执行
          controller.forward();
        }
      }
    });

    controller.forward();
  }

  @override
  Widget build(BuildContext context) {
    return GrowTransition(
      child: Image.asset('assets/images/avatar.png'),
      animation: animation,
    );
  }
}
