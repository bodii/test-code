import 'dart:developer';

import 'package:flutter/material.dart';

class LayoutLogPrint<T> extends StatelessWidget {
  const LayoutLogPrint({
    Key? key,
    this.tag,
    required this.child,
  }) : super(key: key);
  final Widget child;
  final T? tag;

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(builder: (_, constraints) {
      assert(() {
        log('${tag ?? key ?? child}: $constraints');
        return true;
      }());
      return child;
    });
  }
}
