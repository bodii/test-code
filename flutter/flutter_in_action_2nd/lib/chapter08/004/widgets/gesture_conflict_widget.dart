import 'package:flutter/material.dart';

class GestureConflictWidget extends StatefulWidget {
  const GestureConflictWidget({Key? key}) : super(key: key);

  @override
  _GestureConflictWidgetState createState() => _GestureConflictWidgetState();
}

class _GestureConflictWidgetState extends State<GestureConflictWidget> {
  double _left = 0.0;

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        Positioned(
          left: _left,
          child: GestureDetector(
            child: const CircleAvatar(child: Text('A')),
            onHorizontalDragUpdate: (DragUpdateDetails details) {
              setState(
                () => _left += details.delta.dx,
              );
            },
            onHorizontalDragEnd: (DragEndDetails details) {
              debugPrint('onHorizontalDragEnd');
            },
            onTapDown: (TapDownDetails details) {
              debugPrint('down');
            },
            onTapUp: (TapUpDetails details) {
              debugPrint('up');
            },
          ),
        ),
      ],
    );
  }
}
