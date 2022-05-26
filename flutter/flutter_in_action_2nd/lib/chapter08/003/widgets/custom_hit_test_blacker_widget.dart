import 'package:flutter/material.dart';

import 'hit_test_blocker_widget.dart';

class CustomHitTestBlackerWidget extends StatelessWidget {
  const CustomHitTestBlackerWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        HitTestBlockerWidget(child: wChild(1, 200)),
        HitTestBlockerWidget(child: wChild(2, 200)),
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
