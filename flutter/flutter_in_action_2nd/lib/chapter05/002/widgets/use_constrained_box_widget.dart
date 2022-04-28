import 'package:flutter/material.dart';

import 'red_box.dart';

class UseConstrainedBoxWidget extends StatelessWidget {
  const UseConstrainedBoxWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ConstrainedBox(
      constraints: const BoxConstraints(
        minWidth: double.infinity,
        minHeight: 50.0,
      ),
      child: SizedBox(
        height: 5.0,
        child: redBox,
      ),
    );
  }
}
