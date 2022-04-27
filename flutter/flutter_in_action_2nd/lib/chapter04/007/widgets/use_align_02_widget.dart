import 'package:flutter/material.dart';

class UseAlign02Widget extends StatelessWidget {
  const UseAlign02Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.blue.shade50,
      child: const Align(
        widthFactor: 2,
        heightFactor: 2,
        alignment: Alignment.topRight,
        child: FlutterLogo(
          size: 60,
        ),
      ),
    );
  }
}
