import 'package:flutter/material.dart';

import 'red_box.dart';

class UseConstrainedBoxWidget extends StatelessWidget {
  const UseConstrainedBoxWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        ConstrainedBox(
          constraints: const BoxConstraints(
            minWidth: double.infinity,
            minHeight: 50.0,
          ),
          child: const SizedBox(
            height: 5.0,
            child: redBox,
          ),
        ),
        const SizedBox(
          width: 80.0,
          height: 80.0,
          child: redBox,
        ),
      ].map((e) {
        return Padding(
          child: e,
          padding: const EdgeInsets.all(5),
        );
      }).toList(),
    );
  }
}
