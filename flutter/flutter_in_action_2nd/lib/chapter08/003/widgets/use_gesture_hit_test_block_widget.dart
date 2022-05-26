import 'package:flutter/material.dart';

import 'hit_test_blocker_widget.dart';

class UseGestureHitTestBlockWidget extends StatelessWidget {
  const UseGestureHitTestBlockWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Stack(children: [
      HitTestBlockerWidget(child: wChild(1, 200)),
      HitTestBlockerWidget(child: wChild(2, 200)),
    ]);
  }

  Widget wChild(int index, double size) {
    return GestureDetector(
      onTap: () => debugPrint('$index'),
      child: Container(
        width: size,
        height: size,
        color: Colors.grey,
      ),
    );
  }
}
