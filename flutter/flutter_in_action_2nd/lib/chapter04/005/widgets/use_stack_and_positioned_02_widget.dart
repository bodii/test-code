import 'package:flutter/material.dart';

class UseStackAndPositioned02Widget extends StatelessWidget {
  const UseStackAndPositioned02Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Stack(
      alignment: Alignment.center,
      fit: StackFit.expand,
      children: [
        const Positioned(
          left: 18.0,
          child: Text('I am Jack'),
        ),
        Container(
          child:
              const Text('Hello world', style: TextStyle(color: Colors.white)),
          color: Colors.red,
        ),
        const Positioned(child: Text('Your friend'), top: 18.0),
      ],
    );
  }
}
