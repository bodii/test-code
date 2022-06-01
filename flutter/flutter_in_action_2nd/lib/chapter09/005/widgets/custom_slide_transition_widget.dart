import 'package:flutter/material.dart';

import 'my_slide_transition_widget.dart';

class CustomSlideTransitionWidget extends StatefulWidget {
  const CustomSlideTransitionWidget({Key? key}) : super(key: key);

  @override
  _CustomSlideTransitionWidgetState createState() =>
      _CustomSlideTransitionWidgetState();
}

class _CustomSlideTransitionWidgetState
    extends State<CustomSlideTransitionWidget> {
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
              var tween = Tween<Offset>(
                  begin: const Offset(1, 0), end: const Offset(0, 0));
              return MySlideTranitionWidget(
                child: child,
                position: tween.animate(animation),
              );
            },
            child: Text(
              '$_count',
              key: ValueKey<int>(_count),
              style: Theme.of(context).textTheme.headline4,
            ),
          ),
          ElevatedButton(
            onPressed: () {
              setState(() {
                _count++;
              });
            },
            child: const Text('Increment'),
          ),
        ],
      ),
    );
  }
}
