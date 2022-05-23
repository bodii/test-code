import 'package:flutter/material.dart';

import 'widgets/use_future_builder_widget.dart';
import 'widgets/use_stream_builder_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('FutureBuilder和StreamBuilder')),
      body: const UseStreamBuilderWidget(),
    );
  }
}
