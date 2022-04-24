import 'package:flutter/material.dart';

import 'widgets/use_elevated_button_widget.dart';
import 'widgets/use_text_button_widget.dart';
import 'widgets/use_outline_button_widget.dart';
import 'widgets/use_icon_button_widget.dart';
import 'widgets/use_elevated_button_icon_widgets.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Button demo')),
      body: Container(
        alignment: Alignment.topLeft,
        padding: const EdgeInsets.only(top: 20),
        child: Column(
          children: const [
            UseElevatedButtonWidget(),
            UseTextButtonWidget(),
            UseOutlineButtonWidget(),
            UseIconButtonWidget(),
            UseElevatedButtonIconWidgets(),
          ],
        ),
      ),
    );
  }
}
