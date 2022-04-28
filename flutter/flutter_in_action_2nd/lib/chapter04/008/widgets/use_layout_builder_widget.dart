import 'package:flutter/material.dart';

import 'responsive_column.dart';
import 'layout_log_print.dart';

class UseLayoutBuilderWidget extends StatelessWidget {
  const UseLayoutBuilderWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    List<Widget> _children = List.filled(6, const Text('A'));
    return Column(
      children: [
        SizedBox(width: 190, child: ResponsiveColumn(children: _children)),
        ResponsiveColumn(children: _children),
        const LayoutLogPrint(child: Text('xx')),
      ],
    );
  }
}
