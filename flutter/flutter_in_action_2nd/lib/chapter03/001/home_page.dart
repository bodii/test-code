import 'package:flutter/material.dart';

import 'widgets/texts_widget.dart';
import 'widgets/text_span_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Text style demo')),
      body: Container(
        alignment: Alignment.center,
        child: Column(
          children: const [
            TextsWidget(),
            TextSpanWidget(),
          ],
        ),
      ),
    );
  }
}
