import 'package:flutter/material.dart';

class BothDirectionTestWidget extends StatefulWidget {
  const BothDirectionTestWidget({Key? key}) : super(key: key);

  @override
  _BothDirectionTestWidgetState createState() =>
      _BothDirectionTestWidgetState();
}

class _BothDirectionTestWidgetState extends State<BothDirectionTestWidget> {
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
              // 垂直方向拖动事件
              onVerticalDragUpdate: (DragUpdateDetails details) {
                setState(() {
                  _top += details.delta.dy;
                });
              },
              onHorizontalDragUpdate: (DragUpdateDetails details) {
                setState(() {
                  _left += details.delta.dx;
                });
              }),
        ),
      ],
    );
  }
}
