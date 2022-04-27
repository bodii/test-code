import 'package:flutter/material.dart';

class UseAlign03Widget extends StatelessWidget {
  const UseAlign03Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.blue.shade50,
      child: const Align(
        widthFactor: 2.0,
        heightFactor: 2.0,
        alignment: Alignment(2.0, 0.0),
        child: FlutterLogo(
          size: 60,
        ),
      ),
    );
  }
}
