import 'package:flutter/material.dart';

class UseAbsorbPointerListenerWidget extends StatelessWidget {
  const UseAbsorbPointerListenerWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Listener(
      child: AbsorbPointer(
        child: Listener(
          child: Container(
            color: Colors.red,
            width: 200.0,
            height: 100.0,
          ),
          onPointerDown: (event) => debugPrint('in'),
        ),
      ),
      onPointerDown: (event) => debugPrint('up'),
    );
  }
}
