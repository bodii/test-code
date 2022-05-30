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

class GrowTransitionWidget extends StatefulWidget {
  const GrowTransitionWidget({Key? key}) : super(key: key);

  @override
  _GrowTransitionWidgetState createState() => _GrowTransitionWidgetState();
}

class _GrowTransitionWidgetState extends State<GrowTransitionWidget>
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
