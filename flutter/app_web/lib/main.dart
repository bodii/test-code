import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:app_web/flutter_layout_article/HomePage.dart';

void main() {
  debugPaintSizeEnabled = true;

  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Flutter demo',
      home: Scaffold(
        appBar: AppBar(
          title: const Text('Flutter demo'),
        ),
        body: const HomePage(),
      ),
    );
  }
}
