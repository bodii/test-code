import 'package:flutter/material.dart';

class UseGestureDetectorDragWidget extends StatefulWidget {
  const UseGestureDetectorDragWidget({Key? key}) : super(key: key);

  @override
  _UseGestureDetectorDragWidgetState createState() =>
      _UseGestureDetectorDragWidgetState();
}

class _UseGestureDetectorDragWidgetState
    extends State<UseGestureDetectorDragWidget> {
  double _top = 0.0;

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        Positioned(
          top: _top,
          child: GestureDetector(
            child: const CircleAvatar(child: Text('A')),
            onVerticalDragUpdate: (DragUpdateDetails details) {
              setState(() {
                _top += details.delta.dy;
              });
            },
          ),
        ),
      ],
    );
  }
}
