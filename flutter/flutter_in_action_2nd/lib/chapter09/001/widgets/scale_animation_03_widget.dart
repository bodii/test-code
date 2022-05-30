import 'package:flutter/material.dart';

class AnimatedImage extends AnimatedWidget {
  const AnimatedImage({Key? key, required Animation<double> animation})
      : super(
          key: key,
          listenable: animation,
        );

  @override
  Widget build(BuildContext context) {
    final animation = listenable as Animation<double>;
    return Center(
      child: Image.asset(
        'assets/images/avatar.png',
        width: animation.value,
        height: animation.value,
      ),
    );
  }
}

class ScaleAnimation03Widget extends StatefulWidget {
  const ScaleAnimation03Widget({Key? key}) : super(key: key);

  @override
  _ScaleAnimation03WidgetState createState() => _ScaleAnimation03WidgetState();
}

class _ScaleAnimation03WidgetState extends State<ScaleAnimation03Widget>
    with SingleTickerProviderStateMixin {
  late Animation<double> animation;
  late AnimationController controller;

  @override
  void initState() {
    super.initState();
    controller = AnimationController(
      duration: const Duration(seconds: 2),
      vsync: this,
    );
    // 图片宽高从0到300
    animation = Tween(begin: 0.0, end: 300.0).animate(controller);

    // 启动动画
    controller.forward();
  }

  @override
  Widget build(BuildContext context) {
    return AnimatedImage(
      animation: animation,
    );
  }

  @override
  void dispose() {
    controller.dispose();
    super.dispose();
  }
}
