import 'package:flutter/material.dart';

class StackEventWidget extends StatelessWidget {
  const StackEventWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        wChild(1),
        wChild(2),
      ],
    );
  }

  Widget wChild(int index) {
    return Listener(
      behavior: HitTestBehavior.opaque, // 放开此行，点击只会输出2
      onPointerDown: (e) => debugPrint(index.toString()),
      child: Container(
        width: 100.0,
        height: 100.0,
        color: Colors.grey,
      ),
    );
  }
}
