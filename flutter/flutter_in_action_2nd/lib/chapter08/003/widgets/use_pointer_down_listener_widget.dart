import 'package:flutter/material.dart';

import 'pointer_down_listener.dart';

class UsePointerDownListenerWidget extends StatelessWidget {
  const UsePointerDownListenerWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Center(
      child: PointerDownListener(
        child: const Text('Click me'),
        onPointerDown: (e) => debugPrint('down'),
      ),
    );
  }
}
