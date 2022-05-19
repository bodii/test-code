import 'package:flutter/material.dart';

import 'widgets/use_will_pop_scope_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('导航返回拦截WillPopScope'),
        leading: IconButton(
          icon: const Icon(Icons.keyboard_arrow_left),
          onPressed: () {
            Navigator.of(context).pop(true);
          },
        ),
      ),
      body: const UseWillPopScopeWidget(),
    );
  }
}
