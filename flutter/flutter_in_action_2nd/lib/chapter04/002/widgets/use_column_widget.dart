import 'package:flutter/material.dart';

class UseColumnWidget extends StatelessWidget {
  const UseColumnWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ConstrainedBox(
      constraints: const BoxConstraints(minWidth: double.infinity),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.center,
        children: const [
          Text('hi'),
          Text('world'),
        ],
      ),
    );
  }
}
