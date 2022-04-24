import 'package:flutter/material.dart';

class UseElevatedButtonWidget extends StatelessWidget {
  const UseElevatedButtonWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        Center(
          child: ElevatedButton(
            child: const Text('button'),
            onPressed: () {},
          ),
        ),
        const SizedBox(
          height: 10,
        ),
        const Center(
          child: ElevatedButton(
            child: Text('normal'),
            onPressed: null,
          ),
        ),
      ],
    );
  }
}
