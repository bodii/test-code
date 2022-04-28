import 'package:flutter/material.dart';

import 'widgets/use_padding_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Padding')),
      body: Column(
        children: const [
          UsePaddingWidget(),
        ],
      ),
    );
  }
}
