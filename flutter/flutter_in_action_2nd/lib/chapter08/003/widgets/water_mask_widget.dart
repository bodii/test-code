import 'package:flutter/material.dart';

import '../water_mark/water_mark.dart';

class WaterMaskWidget extends StatelessWidget {
  const WaterMaskWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        wChild(1, Colors.white, 200),
        IgnorePointer(
          child: WaterMark(
            painter: TextWaterMarkPainter(text: 'wendux', rotate: -20),
          ),
        ),
      ],
    );
  }

  Widget wChild(int index, color, double size) {
    return Listener(
      onPointerDown: (e) => debugPrint(index.toString()),
      child: Container(
        width: size,
        height: size,
        color: Colors.grey,
      ),
    );
  }
}
