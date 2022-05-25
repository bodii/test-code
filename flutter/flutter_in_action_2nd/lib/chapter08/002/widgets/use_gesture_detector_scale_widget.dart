import 'package:flutter/material.dart';

class UseGestureDetectorScaleWidget extends StatefulWidget {
  const UseGestureDetectorScaleWidget({Key? key}) : super(key: key);

  @override
  _UseGestureDetectorScaleWidgetState createState() =>
      _UseGestureDetectorScaleWidgetState();
}

class _UseGestureDetectorScaleWidgetState
    extends State<UseGestureDetectorScaleWidget> {
  double _width = 200.0;

  @override
  Widget build(BuildContext context) {
    return Center(
      child: GestureDetector(
          child: Image.asset(
            'assets/images/sea.png',
            width: _width,
          ),
          onScaleUpdate: (ScaleUpdateDetails details) {
            setState(() {
              _width = 200 * details.scale.clamp(.8, 10.0);
            });
          }),
    );
  }
}
