import 'package:flutter/material.dart';

class UseGestureWidget extends StatefulWidget {
  const UseGestureWidget({Key? key}) : super(key: key);

  @override
  _UseGestureWidgetState createState() => _UseGestureWidgetState();
}

class _UseGestureWidgetState extends State<UseGestureWidget> {
  String _operation = "No Gesture detected!";

  @override
  Widget build(BuildContext context) {
    return Center(
      child: GestureDetector(
        child: Container(
          alignment: Alignment.center,
          color: Colors.blue,
          width: 200.0,
          height: 100.0,
          child: Text(
            _operation,
            style: const TextStyle(color: Colors.white),
          ),
        ),
        onTap: () => updateText('Tap'),
        onDoubleTap: () => updateText('DoubleTap'),
        onLongPress: () => updateText('LongPress'),
      ),
    );
  }

  void updateText(String text) {
    setState(() {
      _operation = text;
    });
  }
}
