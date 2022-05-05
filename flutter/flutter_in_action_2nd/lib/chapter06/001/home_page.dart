import 'package:flutter/material.dart';

import 'widgets/use_single_child_scroll_view.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('SingleChildScrollView'),
      ),
      body: const UseSingleChildScrollView(),
    );
  }
}
