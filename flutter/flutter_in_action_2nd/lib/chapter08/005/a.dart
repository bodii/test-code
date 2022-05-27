import 'package:flutter/material.dart';

import 'widgets/event_bus.dart';
import 'b.dart';

class A extends StatelessWidget {
  const A({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    bus.on(
      'login',
      (arg) => debugPrint('on'),
    );
    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          const Text('A page'),
          OutlinedButton(
            onPressed: () {
              Navigator.push(
                context,
                MaterialPageRoute(
                  builder: (context) => const B(),
                ),
              );
            },
            child: const Text('go to B page'),
          ),
        ],
      ),
    );
  }
}
