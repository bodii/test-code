import 'package:flutter/material.dart';

import 'water_mark_demo.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('水印实例：文本绘制与离屏渲染demo'),
      ),
      body: const WaterMarkDemo(),
    );
  }
}
