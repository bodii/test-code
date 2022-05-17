import 'package:flutter/material.dart';

import 'widgets/two_list_view_widget.dart';
import 'widgets/two_sliver_list_widget.dart';
import 'widgets/use_custom_scroll_view_01_widget.dart';
import 'widgets/persistent_header_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('CustomScrollView Demo')),
      body: const PersistentHeaderWidget(),
    );
  }
}
