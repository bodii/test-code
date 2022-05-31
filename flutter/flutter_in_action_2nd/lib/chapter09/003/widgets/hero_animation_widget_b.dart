import 'package:flutter/material.dart';

class HeroAnimationWidgetB extends StatelessWidget {
  const HeroAnimationWidgetB({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Hero(
        tag: 'avatar',
        child: Image.asset('assets/images/avatar.png'),
      ),
    );
  }
}
