import 'package:flutter/material.dart';

import '../../../chapter04/008/widgets/layout_log_print.dart';

class UseListViewBuilderItemExtentWidget extends StatelessWidget {
  const UseListViewBuilderItemExtentWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
        // prototypeItem: const ListTile(title: Text('1')),
        itemExtent: 56,
        itemBuilder: (BuildContext context, int index) {
          return LayoutLogPrint(
            tag: index,
            child: ListTile(title: Text("$index")),
          );
        });
  }
}
