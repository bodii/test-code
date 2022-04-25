import 'package:flutter/material.dart';

import 'widgets/linear_indicators_widget.dart';
import 'widgets/circular_indicator_widgets.dart';
import 'widgets/use_custom_size_linear_indicator_widget.dart';
import 'widgets/use_custom_value_color_animation_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.only(top: 20, left: 100, right: 100),
      child: Column(
        children: const [
          LinearIndicatorswidget(),
          CircularIndicatorWidgets(),
          UseCustomSizeLinearIndicatorWidget(),
          UseCustomValueColorAnimationWidget(),
        ],
      ),
    );
  }
}
