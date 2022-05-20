import 'package:flutter/material.dart';

import 'nav_bar_widget.dart';

class UseColorComputeLuminance extends StatelessWidget {
  const UseColorComputeLuminance({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('color computeLuminance')),
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: const [
          NavBarWidget(color: Colors.blue, title: '测试'),
          SizedBox(height: 10),
          NavBarWidget(color: Colors.white, title: '测试'),
        ],
      ),
    );
  }
}
