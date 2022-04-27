import 'package:flutter/material.dart';

class UseStackAndPositionedWidget extends StatelessWidget {
  const UseStackAndPositionedWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ConstrainedBox(
      constraints: const BoxConstraints.expand(),
      child: Stack(
        alignment: Alignment.center,
        children: [
          Container(
            child: const Text('Hello world',
                style: TextStyle(color: Colors.white)),
            color: Colors.red,
          ),
          const Positioned(
            left: 18.0,
            child: Text('I am Jack'),
          ),
          const Positioned(
            top: 18.0,
            child: Text('Your friend'),
          ),
        ],
      ),
    );
  }
}
