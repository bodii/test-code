import 'package:flutter/material.dart';

import 'animated_decorated_box.dart';

class UseAnimatedsWidgets extends StatefulWidget {
  const UseAnimatedsWidgets({Key? key}) : super(key: key);

  @override
  _UseAnimatedsWidgetsState createState() => _UseAnimatedsWidgetsState();
}

class _UseAnimatedsWidgetsState extends State<UseAnimatedsWidgets> {
  double _padding = 10;
  Alignment _align = Alignment.topRight;
  double _height = 100;
  double _left = 0;
  Color _color = Colors.red;
  TextStyle _style = const TextStyle(color: Colors.black);
  Color _decorationColor = Colors.blue;

  @override
  Widget build(BuildContext context) {
    Duration duration = const Duration(seconds: 5);
    return SingleChildScrollView(
      child: Column(
        children: [
          ElevatedButton(
            onPressed: () {
              setState(() {
                _padding = 20;
              });
            },
            child: AnimatedPadding(
              duration: duration,
              padding: EdgeInsets.all(_padding),
              child: const Text('AnimatedPadding'),
            ),
          ),
          SizedBox(
            height: 50,
            child: Stack(
              children: [
                AnimatedPositioned(
                  duration: duration,
                  left: _left,
                  child: ElevatedButton(
                    child: const Text('AnimatedPositioned'),
                    onPressed: () {
                      setState(() {
                        _left = 100;
                      });
                    },
                  ),
                ),
              ],
            ),
          ),
          Container(
            height: 100,
            color: Colors.grey,
            child: AnimatedAlign(
              duration: duration,
              alignment: _align,
              child: ElevatedButton(
                child: const Text('AnimatedAlign'),
                onPressed: () {
                  setState(() {
                    _align = Alignment.center;
                  });
                },
              ),
            ),
          ),
          AnimatedContainer(
            duration: duration,
            height: _height,
            color: _color,
            child: TextButton(
              child: const Text(
                'AnimatedContainer',
                style: TextStyle(color: Colors.white),
              ),
              onPressed: () {
                setState(() {
                  _height = 150;
                  _color = Colors.blue;
                });
              },
            ),
          ),
          AnimatedDefaultTextStyle(
            child: GestureDetector(
                child: const Text('AnimatedDefaultTextStyle'),
                onTap: () {
                  setState(() {
                    _style = const TextStyle(
                      color: Colors.blue,
                      decorationStyle: TextDecorationStyle.solid,
                      decorationColor: Colors.blue,
                    );
                  });
                }),
            style: _style,
            duration: duration,
          ),
          AnimatedDecoratedBox(
            duration: duration,
            decoration: BoxDecoration(color: _decorationColor),
            child: TextButton(
              child: const Text('AnimatedDecoratedBox',
                  style: TextStyle(color: Colors.white)),
              onPressed: () {
                setState(() {
                  _decorationColor = Colors.red;
                });
              },
            ),
          ),
        ]
            .map(
              (e) => Padding(
                padding: const EdgeInsets.symmetric(vertical: 16),
                child: e,
              ),
            )
            .toList(),
      ),
    );
  }
}
