import 'package:flutter/material.dart';

class AnimatedSwitcherCounterWidget extends StatefulWidget {
  const AnimatedSwitcherCounterWidget({Key? key}) : super(key: key);

  @override
  _AnimatedSwitcherCounterWidgetState createState() =>
      _AnimatedSwitcherCounterWidgetState();
}

class _AnimatedSwitcherCounterWidgetState
    extends State<AnimatedSwitcherCounterWidget> {
  int _count = 0;
  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          AnimatedSwitcher(
            duration: const Duration(milliseconds: 500),
            transitionBuilder: (Widget child, Animation<double> animation) {
              return ScaleTransition(child: child, scale: animation);
            },
            child: Text(
              '$_count',
              // 显示指定key, 不同的key会被认为是不同的Text,这样才能执行动画
              key: ValueKey<int>(_count),
              style: Theme.of(context).textTheme.headline4,
            ),
          ),
          ElevatedButton(
            child: const Text('+1'),
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
