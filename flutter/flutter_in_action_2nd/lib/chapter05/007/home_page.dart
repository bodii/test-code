import 'package:flutter/material.dart';

import 'widgets/use_fitted_box_01_widget.dart';
import 'widgets/use_fitted_box_02_widget.dart';
import 'widgets/use_fitted_box_03_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('空间适配FittedBox')),
      body: Column(
        children: const [
          UseFittedBox01Widget(),
          UseFittedBox02Widget(),
          UseFittedBox03Widget(),
        ],
      ),
    );
  }
}
