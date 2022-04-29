import 'package:flutter/material.dart';

class UseContainer01Widget extends StatelessWidget {
  const UseContainer01Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.only(top: 50.0, left: 120.0),
      constraints: const BoxConstraints.tightFor(width: 200.0, height: 150),
      decoration: const BoxDecoration(
        gradient: RadialGradient(
          colors: [Colors.red, Colors.orange],
          center: Alignment.topLeft,
          radius: .98,
        ),
        boxShadow: [
          BoxShadow(
            color: Colors.black54,
            offset: Offset(2.0, 2.0),
            blurRadius: 4.0,
          ),
        ],
      ),
      transform: Matrix4.rotationZ(.2), // 倾斜
      alignment: Alignment.center,
      child: const Text(
        '5.20',
        style: TextStyle(color: Colors.white, fontSize: 40.0),
      ),
    );
  }
}
