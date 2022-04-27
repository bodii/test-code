import 'package:flutter/material.dart';

import 'widgets/use_align_01_widget.dart';
import 'widgets/use_align_02_widget.dart';
import 'widgets/use_align_03_widget.dart';
import 'widgets/use_align_04_widget.dart';
import 'widgets/use_center_01_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('对齐与相对定位(Align)')),
      body: Container(
        margin: const EdgeInsets.all(10),
        child: Wrap(
          spacing: 10,
          runSpacing: 10,
          children: const [
            UseAlign01Widget(),
            UseAlign02Widget(),
            UseAlign03Widget(),
            UseAlign04Widget(),
            UseCenter01Widget(),
          ],
        ),
      ),
    );
  }
}
