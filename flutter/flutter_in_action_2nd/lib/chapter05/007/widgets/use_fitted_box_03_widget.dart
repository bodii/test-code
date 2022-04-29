import 'package:flutter/material.dart';

import 'single_line_fitted_box_widget.dart';

class UseFittedBox03Widget extends StatelessWidget {
  const UseFittedBox03Widget({Key? key}) : super(key: key);

  Widget wRow(String text) {
    Widget child = Text(text);
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceEvenly,
      children: [child, child, child],
    );
  }

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        children: [
          SingleLineFittedBoxWidget(
            child: wRow(' 90000000000000000 '),
          ),
          wRow(' 800 '),
          SingleLineFittedBoxWidget(child: wRow(' 800 '))
        ]
            .map((e) => Padding(
                  padding: const EdgeInsets.symmetric(vertical: 20),
                  child: e,
                ))
            .toList(),
      ),
    );
  }
}
