import 'package:flutter/material.dart';

import 'red_box.dart';

class UseUnconstrainedBoxWidget extends StatelessWidget {
  const UseUnconstrainedBoxWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ConstrainedBox(
      constraints: const BoxConstraints(
        minWidth: 60.0,
        minHeight: 100.0,
      ),
      child: UnconstrainedBox(
        child: ConstrainedBox(
          constraints: const BoxConstraints(
            minWidth: 90.0,
            minHeight: 20.0,
          ),
          child: redBox,
        ),
      ),
    );
  }
}
