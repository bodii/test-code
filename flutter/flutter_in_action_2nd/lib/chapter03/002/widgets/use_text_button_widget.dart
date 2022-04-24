import 'package:flutter/material.dart';

class UseTextButtonWidget extends StatelessWidget {
  const UseTextButtonWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Center(
      child: TextButton(
        child: const Text('normal'),
        onPressed: () {},
      ),
    );
  }
}
