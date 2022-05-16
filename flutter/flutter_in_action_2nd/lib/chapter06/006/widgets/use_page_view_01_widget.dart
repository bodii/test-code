import 'package:flutter/material.dart';

import 'page_text.dart';

class UsePageView01Widget extends StatelessWidget {
  const UsePageView01Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    List<Widget> children = [];

    // 生成6个tab页
    for (int i = 0; i < 6; ++i) {
      children.add(PageText(text: '$i'));
    }

    return PageView(
      scrollDirection: Axis.vertical,
      children: children,
    );
  }
}
