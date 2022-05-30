import 'package:flutter/material.dart';

class ScaleAnimation02Widget extends StatefulWidget {
  const ScaleAnimation02Widget({Key? key}) : super(key: key);

  @override
  _ScaleAnimationWidgetState createState() => _ScaleAnimationWidgetState();
}

class _ScaleAnimationWidgetState extends State<ScaleAnimation02Widget>
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

    // 使用弹性曲线
    animation = CurvedAnimation(
      parent: controller,
      curve: Curves.bounceIn,
    );

    // 图片宽高从0到300
    animation = Tween(begin: 0.0, end: 300.0).animate(animation)
      ..addListener(() {
        setState(() => {});
      });

    // 启动动画
    controller.forward();
  }

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Image.asset(
        'assets/images/avatar.png',
        width: animation.value,
        height: animation.value,
      ),
    );
  }

  @override
  void dispose() {
    controller.dispose();
    super.dispose();
  }
}
