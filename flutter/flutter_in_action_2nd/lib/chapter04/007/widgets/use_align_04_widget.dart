import 'package:flutter/material.dart';

class UseAlign04Widget extends StatelessWidget {
  const UseAlign04Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      width: 120,
      height: 120,
      color: Colors.blue[50],
      child: const Align(
        alignment: FractionalOffset(0.2, 0.6),
        child: FlutterLogo(size: 60),
      ),
    );
  }
}
