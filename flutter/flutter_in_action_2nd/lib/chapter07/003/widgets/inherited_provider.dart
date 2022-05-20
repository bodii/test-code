import 'package:flutter/material.dart';

class InheritedProvider<T> extends InheritedWidget {
  const InheritedProvider({Key? key, required this.data, required Widget child})
      : super(key: key, child: child);

  final T data;

  @override
  bool updateShouldNotify(InheritedProvider<T> oldWidget) {
    return true;
  }
}
