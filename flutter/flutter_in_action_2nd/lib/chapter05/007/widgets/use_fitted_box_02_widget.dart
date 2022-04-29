import 'package:flutter/material.dart';

class UseFittedBox02Widget extends StatelessWidget {
  const UseFittedBox02Widget({Key? key}) : super(key: key);

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
          FittedBox(
            child: wRow(' 90000000000000000 '),
          ),
          wRow(' 800 '),
          FittedBox(child: wRow(' 800 '))
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
