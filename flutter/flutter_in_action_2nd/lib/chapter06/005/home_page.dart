import 'package:flutter/material.dart';

import 'widgets/use_grid_view_01_widgets.dart';
import 'widgets/use_grid_view_count_01_widget.dart';
import 'widgets/use_grid_view_02_widget.dart';
import 'widgets/use_grid_view_extent_widget.dart';
import 'widgets/use_grid_view_build_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('GridView')),
      body: const UseGridViewBuildWidget(),
    );
  }
}
