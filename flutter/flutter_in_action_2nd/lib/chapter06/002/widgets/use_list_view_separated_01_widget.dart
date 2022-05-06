import 'package:flutter/material.dart';

class UseListViewSeparated01Widget extends StatelessWidget {
  const UseListViewSeparated01Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    const Widget divider1 = Divider(color: Colors.blue);
    const Widget divider2 = Divider(color: Colors.green);
    return ListView.separated(
      itemCount: 100,
      itemBuilder: (BuildContext context, int index) {
        return ListTile(title: Text("$index"));
      },
      separatorBuilder: (BuildContext context, int index) {
        return index & 1 == 0 ? divider1 : divider2;
      },
    );
  }
}
