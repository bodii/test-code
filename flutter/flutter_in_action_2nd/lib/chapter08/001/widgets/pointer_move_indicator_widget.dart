import 'package:flutter/material.dart';

class PointerMoveIndicatorWidget extends StatefulWidget {
  const PointerMoveIndicatorWidget({Key? key}) : super(key: key);

  @override
  _PointerMoveIndicatorWidgetState createState() =>
      _PointerMoveIndicatorWidgetState();
}

class _PointerMoveIndicatorWidgetState
    extends State<PointerMoveIndicatorWidget> {
  PointerEvent? _event;
  @override
  Widget build(BuildContext context) {
    return Listener(
      child: Container(
        alignment: Alignment.center,
        color: Colors.blue,
        width: 300.0,
        height: 150.0,
        child: Text(
          '${_event?.localPosition ?? ''}',
          style: const TextStyle(color: Colors.white),
        ),
      ),
      onPointerDown: (PointerDownEvent event) => setState(() => _event = event),
      onPointerMove: (PointerMoveEvent event) => setState(() => _event = event),
      onPointerUp: (PointerUpEvent event) => setState(() => _event = event),
    );
  }
}
