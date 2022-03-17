import 'package:flutter/material.dart';

import 'package:app_web/flutter_layout_article/example.dart';

class Example2 extends Example {
  const Example2({Key? key}) : super(key: key);

  @override
  final code = 'Container(width: 100, height: 100, color: red)';

  @override
  final String explanation =
      'The red Container wants to be 100x100, but it can\'t, '
      'because the screen forces it to be exactly the same size as the screen.'
      '\n\n'
      'So the Container fills the screen.';

  @override
  Widget build(BuildContext context) {
    return Container(width: 100, height: 100, color: red);
  }
}
