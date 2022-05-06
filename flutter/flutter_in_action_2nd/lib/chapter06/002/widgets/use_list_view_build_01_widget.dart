import 'package:flutter/material.dart';

class UseListViewBuild01Widget extends StatelessWidget {
  const UseListViewBuild01Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
        itemCount: 100,
        itemExtent: 50.0, // 强制高度为50.0
        itemBuilder: (BuildContext context, int index) {
          return ListTile(title: Text("$index"));
        });
  }
}
