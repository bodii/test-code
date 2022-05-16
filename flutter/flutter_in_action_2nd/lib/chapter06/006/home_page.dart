import 'package:flutter/material.dart';

import 'widgets/use_page_view_01_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('PageView')),
      body: const UsePageView01Widget(),
    );
  }
}
