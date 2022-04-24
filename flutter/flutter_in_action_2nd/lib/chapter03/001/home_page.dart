import 'package:flutter/material.dart';
import 'package:flutter_in_action_2nd/chapter03/001/widgets/use_package_font_widget.dart';

import 'widgets/texts_widget.dart';
import 'widgets/text_span_widget.dart';
import 'widgets/default_text_style_widget.dart';

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
            DefaultTextStyleWidget(),
            UsePackageFontWidget(),
          ],
        ),
      ),
    );
  }
}
