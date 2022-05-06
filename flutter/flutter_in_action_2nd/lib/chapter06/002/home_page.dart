import 'package:flutter/material.dart';

import 'widgets/use_list_view_01_widget.dart';
import 'widgets/use_list_view_build_01_widget.dart';
import 'widgets/use_list_view_separated_01_widget.dart';
import 'widgets/use_list_view_builder_item_extent_widget.dart';
import 'widgets/use_list_view_test_widget.dart';
import 'widgets/use_list_view_flex_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('ListView')),
      body: const UseListViewFlexWidget(),
    );
  }
}
