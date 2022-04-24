import 'package:flutter/material.dart';

import 'widgets/text_field_widget.dart';
import 'widgets/use_focus_widget.dart';
import 'widgets/use_custom_style_widget.dart';
import 'widgets/use_theme_style_widget.dart';
import 'widgets/use_custom_border_side_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('input text field and form')),
      body: Column(
        children: const [
          // TextFieldWidget(),
          UseFocusWidget(),
          UseCustomStyleWidget(),
          UseThemeStyleWidget(),
          UseCustomBorderSideWidget(),
        ],
      ),
    );
  }
}
