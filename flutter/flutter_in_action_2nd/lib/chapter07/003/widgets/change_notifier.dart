import 'package:flutter/material.dart';

class ChangeNotifier implements Listenable {
  List listeners = [];

  @override
  void addListener(VoidCallback listener) {
    // 添加监听器
    listeners.add(listener);
  }

  @override
  void removeListener(VoidCallback listener) {
    // 移除监听器
    listeners.remove(listener);
  }

  void notifyLiseners() {
    // 通知所有监听器，触发监听器回调
    for (var item in listeners) {
      item();
    }
  }
}
