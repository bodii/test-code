import 'package:flutter/material.dart';

class UseAlign01Widget extends StatelessWidget {
  const UseAlign01Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      height: 120.0,
      width: 120.0,
      color: Colors.blue.shade50,
      child: const Align(
        alignment: Alignment.topRight,
        child: FlutterLogo(
          size: 60,
        ),
      ),
    );
  }
}
