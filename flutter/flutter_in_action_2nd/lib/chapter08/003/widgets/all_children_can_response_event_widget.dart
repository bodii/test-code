import 'package:flutter/material.dart';

class AllChildrenCanResponseEventWidget extends StatelessWidget {
  const AllChildrenCanResponseEventWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        IgnorePointer(child: wChild(1, 200)),
        IgnorePointer(child: wChild(2, 200)),
      ],
    );
  }

  Widget wChild(int index, double size) {
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
