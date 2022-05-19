import 'package:flutter/material.dart';

import 'widgets/use_nested_scroll_view.dart';
import 'widgets/use_snap_app_bar.dart';
import 'widgets/use_nested_tab_bar_view.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const Scaffold(body: UseNestedTabBarView());
  }
}
