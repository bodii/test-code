import 'package:flutter/material.dart';

class UseElevatedButtonIconWidgets extends StatelessWidget {
  const UseElevatedButtonIconWidgets({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      alignment: Alignment.center,
      child: Column(
        mainAxisAlignment: MainAxisAlignment.spaceEvenly,
        children: [
          ElevatedButton.icon(
            icon: const Icon(Icons.send),
            label: const Text('发送1'),
            onPressed: () {},
          ),
          Container(
            height: 40.0,
            width: 120.0,
            margin: const EdgeInsets.symmetric(vertical: 8),
            child: ElevatedButton.icon(
              icon: const Icon(Icons.send),
              label: const Text('发送2'),
              onPressed: () {},
            ),
          ),
          OutlinedButton.icon(
            icon: const Icon(Icons.add),
            label: const Text('添加'),
            onPressed: () {},
          ),
          TextButton.icon(
            icon: const Icon(Icons.info),
            label: const Text('详情'),
            onPressed: () {},
          ),
        ],
      ),
    );
  }
}
