import 'package:flutter/material.dart';

class ListenerGestureConflictWidget extends StatefulWidget {
  const ListenerGestureConflictWidget({Key? key}) : super(key: key);

  @override
  _ListenerGestureConflictWidgetState createState() =>
      _ListenerGestureConflictWidgetState();
}

class _ListenerGestureConflictWidgetState
    extends State<ListenerGestureConflictWidget> {
  double _leftB = 0.0;
  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        Positioned(
          top: 80.0,
          left: _leftB,
          child: Listener(
            onPointerDown: (PointerDownEvent event) {
              debugPrint('down');
            },
            onPointerUp: (PointerUpEvent event) {
              debugPrint('up');
            },
            child: GestureDetector(
              child: const CircleAvatar(child: Text('A')),
              onHorizontalDragUpdate: (DragUpdateDetails details) {
                setState(
                  () => _leftB += details.delta.dx,
                );
              },
              onHorizontalDragEnd: (DragEndDetails details) {
                debugPrint('onHorizontalDragEnd');
              },
            ),
          ),
        ),
      ],
    );
  }
}
