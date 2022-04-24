import 'package:flutter/material.dart';

class UseImageWidget extends StatelessWidget {
  const UseImageWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Row(
      mainAxisAlignment: MainAxisAlignment.spaceAround,
      children: [
        const Center(
          child: Image(
            image: AssetImage('assets/images/avatar.png'),
            width: 100.0,
          ),
        ),
        Image.asset(
          'assets/images/avatar.png',
          width: 100.0,
        )
      ],
    );
  }
}
