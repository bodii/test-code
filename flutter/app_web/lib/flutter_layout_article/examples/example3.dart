import 'package:flutter/material.dart';
import 'package:app_web/flutter_layout_article/example.dart';

class Example3 extends Example {
  const Example3({Key? key}) : super(key: key);

  @override
  final code = 'Center(\n'
      '   child: Container(width: 100, height: 100, color: red))';

  @override
  final String explanation =
      'The screen forces the Center to be exactly the same size as the screen,'
      'so the Center fills the screen.'
      '\n\n'
      'The Center tells the Container that it can be any size it wants, but not bigger than the screen.'
      'Now the Container can indeed be 100x100.';

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Container(
        width: 100,
        height: 100,
        color: red,
      ),
    );
  }
}
