import 'package:flutter/material.dart';

class HitTestBeHaviorWidget extends StatelessWidget {
  const HitTestBeHaviorWidget({Key? key}) : super(key: key);

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
      // behavior: HitTestBehavior.opaque, // 放开此行，点击只会输出2
      behavior: HitTestBehavior.translucent, // 放开此行，点击会同时输出2和1
      onPointerDown: (e) => debugPrint(index.toString()),
      // SizedBox 没有子元素，当它被点击时，它的 hitTest 就会返回 false，
      // 此时 Listener  的 behavior 设置为 opaque 和translucent 就会有区别
      child: const SizedBox.expand(),
    );
  }
}
