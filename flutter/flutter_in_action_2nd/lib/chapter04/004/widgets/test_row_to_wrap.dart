import 'package:flutter/material.dart';

class TestRowToWrap extends StatelessWidget {
  const TestRowToWrap({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    // return Row(
    //   children: [
    //     Text('xxx' * 100),
    //   ],
    // );
    return Wrap(
      children: [
        Text('xxx' * 100),
      ],
    );
  }
}
