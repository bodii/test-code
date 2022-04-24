import 'package:flutter/material.dart';

class UsePackageFontWidget extends StatelessWidget {
  const UsePackageFontWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const Text(
      'Tangerine font demo',
      style: TextStyle(
        fontFamily: 'Tangerine',
      ),
    );
  }
}
