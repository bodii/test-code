import 'package:flutter/material.dart';

class DefaultTextStyleWidget extends StatelessWidget {
  const DefaultTextStyleWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return DefaultTextStyle(
      // 默认样式
      style: const TextStyle(
        color: Colors.red,
        fontSize: 20.0,
      ),
      textAlign: TextAlign.start,
      child: Column(
        children: const [
          Text('hello world'),
          Text('I am Jack'),
          Text(
            'I am Lucy',
            style: TextStyle(
              inherit: false,
              color: Colors.grey,
            ),
          )
        ],
      ),
    );
  }
}
