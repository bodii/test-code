import 'dart:developer';

import 'package:flutter/material.dart';

import 'keep_alive_wrapper.dart';

class KeepAliveTest extends StatelessWidget {
  const KeepAliveTest({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ListView.builder(
      itemBuilder: (_, index) {
        return KeepAliveWrapper(
          keepAlive: false,
          child: ListItem(index: index),
        );
      },
    );
  }
}

class ListItem extends StatefulWidget {
  const ListItem({Key? key, required this.index}) : super(key: key);
  final int index;

  @override
  _ListItemState createState() => _ListItemState();
}

class _ListItemState extends State<ListItem> {
  @override
  Widget build(BuildContext context) {
    return ListTile(title: Text('${widget.index}'));
  }

  @override
  void dispose() {
    log('dispose ${widget.index}');
    super.dispose();
  }
}
