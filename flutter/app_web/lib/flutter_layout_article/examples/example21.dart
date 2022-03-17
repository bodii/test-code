import 'package:flutter/material.dart';
import 'package:app_web/flutter_layout_article/example.dart';

class Example21 extends Example {
  const Example21({Key? key}) : super(key: key);

  @override
  final code = 'Center(\n'
      '   child: Text(\'…\'));';
  @override
  final String explanation = 'If, however, you remove the FittedBox, '
      'the Text gets its maximum width from the screen, '
      'and breaks the line so that it fits the screen.';

  @override
  Widget build(BuildContext context) {
    return const Center(
      child: Text(
          'This is some very very very large text that is too big to fit a regular screen in a single line.'),
    );
  }
}
