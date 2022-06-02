import 'package:flutter/material.dart';

import 'turn_box.dart';

class UseTurnBoxWidget extends StatefulWidget {
  const UseTurnBoxWidget({Key? key}) : super(key: key);

  @override
  _UseTurnBoxWidgetState createState() => _UseTurnBoxWidgetState();
}

class _UseTurnBoxWidgetState extends State<UseTurnBoxWidget> {
  double _turns = .0;

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        children: [
          TurnBox(
            child: const Icon(
              Icons.refresh,
              size: 50,
            ),
            turns: _turns,
            speed: 500,
          ),
          TurnBox(
            child: const Icon(Icons.refresh, size: 150.0),
            turns: _turns,
            speed: 1000,
          ),
          ElevatedButton(
            onPressed: () {
              setState(() {
                _turns += .2;
              });
            },
            child: const Text('顺时针旋转1/5圈'),
          ),
          const SizedBox(height: 10),
          ElevatedButton(
            onPressed: () {
              setState(() {
                _turns -= .2;
              });
            },
            child: const Text('逆时针旋转1/5圈'),
          ),
        ],
      ),
    );
  }
}
