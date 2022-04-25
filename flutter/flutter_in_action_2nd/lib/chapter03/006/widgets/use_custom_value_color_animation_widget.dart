import 'package:flutter/material.dart';

class UseCustomValueColorAnimationWidget extends StatefulWidget {
  const UseCustomValueColorAnimationWidget({Key? key}) : super(key: key);

  @override
  State<UseCustomValueColorAnimationWidget> createState() =>
      _UseCustomValueColorAnimationWidgetState();
}

class _UseCustomValueColorAnimationWidgetState
    extends State<UseCustomValueColorAnimationWidget>
    with SingleTickerProviderStateMixin {
  late AnimationController _animationController;

  @override
  void initState() {
    // 动画执行进间3秒
    _animationController =
        AnimationController(vsync: this, duration: const Duration(seconds: 3));
    _animationController.forward();
    _animationController.addListener(() => setState(() {}));
    super.initState();
  }

  @override
  void dispose() {
    _animationController.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return SingleChildScrollView(
      child: Column(
        children: [
          Padding(
            padding: const EdgeInsets.all(10),
            child: LinearProgressIndicator(
              backgroundColor: Colors.grey[200],
              valueColor: ColorTween(begin: Colors.grey, end: Colors.blue)
                  .animate(_animationController),
              value: _animationController.value,
            ),
          ),
        ],
      ),
    );
  }
}
