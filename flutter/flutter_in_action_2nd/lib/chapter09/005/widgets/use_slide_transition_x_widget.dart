import 'package:flutter/material.dart';

import 'slide_transition_x.dart';

class UseSlideTransitionXWidget extends StatefulWidget {
  const UseSlideTransitionXWidget({Key? key}) : super(key: key);

  @override
  _UseSlideTransitionXWidgetState createState() =>
      _UseSlideTransitionXWidgetState();
}

class _UseSlideTransitionXWidgetState extends State<UseSlideTransitionXWidget> {
  int _count = 0;
  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          AnimatedSwitcher(
            duration: const Duration(milliseconds: 200),
            transitionBuilder: (Widget child, Animation<double> animation) {
              return SlideTransitionX(
                position: animation,
                child: child,
                direction: AxisDirection.down,
              );
            },
            child: Text(
              '$_count',
              key: ValueKey<int>(_count),
              style: Theme.of(context).textTheme.headline4,
            ),
          ),
          ElevatedButton(
            child: const Text('Increment'),
            onPressed: () {
              setState(() {
                _count++;
              });
            },
          ),
        ],
      ),
    );
  }
}
