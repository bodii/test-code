import 'package:flutter/material.dart';

import 'widgets/use_clip_01_widget.dart';
import 'widgets/use_custom_clipper_01_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('剪裁(Clip)')),
      body: Column(
        children: const [
          UseClip01Widget(),
          UseCustomClipper01Widget(),
        ],
      ),
    );
  }
}
