import 'package:flutter/material.dart';

import 'tip_route.dart';

class RouterTestRoute extends StatelessWidget {
  const RouterTestRoute({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('打开提示页Demo'),
      ),
      body: Center(
        child: ElevatedButton(
          child: const Text('打开提示页'),
          onPressed: () async {
            var result = await Navigator.push(
              context,
              MaterialPageRoute(
                builder: (context) {
                  return const TipRoute(
                    text: '我是提示xxxx',
                  );
                },
              ),
            );
            // ignore: avoid_print
            print('路由返回值: $result');
          },
        ),
      ),
    );
  }
}
