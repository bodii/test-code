import 'package:flutter/material.dart';

import 'widgets/test_row_to_wrap.dart';
import 'widgets/use_wrap_chip_widget.dart';
import 'widgets/use_flow_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('流式布局 Wrap、Flow demo')),
      body: Column(
        children: const [
          TestRowToWrap(),
          UseWrapChipWidget(),
          UseFlowWidget(),
        ],
      ),
    );
  }
}
