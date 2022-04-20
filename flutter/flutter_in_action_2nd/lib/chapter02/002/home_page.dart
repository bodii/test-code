import 'package:flutter/material.dart';
import 'parent_widget_c.dart';
import 'parent_widget.dart';
import 'tap_box_a.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const ParentWidgetC();
  }
}
