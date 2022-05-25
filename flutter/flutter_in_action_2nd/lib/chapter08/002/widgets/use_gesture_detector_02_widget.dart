import 'package:flutter/material.dart';

class UseGestureDetector02Widget extends StatefulWidget {
  const UseGestureDetector02Widget({Key? key}) : super(key: key);

  @override
  _UseGestureDetector02WidgetState createState() =>
      _UseGestureDetector02WidgetState();
}

class _UseGestureDetector02WidgetState extends State<UseGestureDetector02Widget>
    with SingleTickerProviderStateMixin {
  double _top = 0.0;
  double _left = 0.0;

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        Positioned(
          top: _top,
          left: _left,
          child: GestureDetector(
              child: const CircleAvatar(child: Text('A')),
              onPanDown: (DragDownDetails e) {
                debugPrint('用户手指按下: ${e.globalPosition}');
              },
              onPanUpdate: (DragUpdateDetails e) {
                setState(() {
                  _left += e.delta.dx;
                  _top += e.delta.dy;
                });
              },
              onPanEnd: (DragEndDetails e) {
                debugPrint(e.velocity.toString());
              }),
        ),
      ],
    );
  }
}
