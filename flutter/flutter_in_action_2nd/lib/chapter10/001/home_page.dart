import 'package:flutter/material.dart';

import 'widgets/use_gradient_button_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('自定义组件-组合组件')),
      body: const UseGradientButtonWidget(),
    );
  }
}
