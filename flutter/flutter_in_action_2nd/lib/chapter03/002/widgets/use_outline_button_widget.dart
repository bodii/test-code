import 'package:flutter/material.dart';

class UseOutlineButtonWidget extends StatelessWidget {
  const UseOutlineButtonWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      padding: const EdgeInsets.only(top: 10),
      child: Center(
        child: OutlinedButton(
          child: const Text('outlined'),
          onPressed: () {},
        ),
      ),
    );
  }
}
