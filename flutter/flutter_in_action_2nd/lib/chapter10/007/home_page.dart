import 'package:flutter/material.dart';

import 'widgets/use_text_water_mark_painter_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('custom water mark widget')),
      body: const UseTextWaterMarkPainterWidget(),
    );
  }
}
