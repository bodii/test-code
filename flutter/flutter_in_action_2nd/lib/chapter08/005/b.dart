import 'package:flutter/material.dart';

import 'widgets/event_bus.dart';
import 'a.dart';

class B extends StatelessWidget {
  const B({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    bus.emit(
      'login',
      'B',
    );
    return Scaffold(
      appBar: AppBar(title: const Text('事件总线')),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: [
            const Text('B page'),
            OutlinedButton(
              onPressed: () {
                Navigator.of(context).pop();
              },
              child: const Text('go to A page'),
            ),
          ],
        ),
      ),
    );
  }
}
