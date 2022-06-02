import 'package:flutter/material.dart';

import 'widgets/custom_hero_animation_widget.dart';
import 'widgets/use_hero_animation_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Hero动画')),
      body: const UseHeroAnimationWidget(),
    );
  }
}