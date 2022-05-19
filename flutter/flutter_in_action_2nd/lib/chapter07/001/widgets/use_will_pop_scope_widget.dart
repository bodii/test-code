import 'package:flutter/material.dart';

class UseWillPopScopeWidget extends StatefulWidget {
  const UseWillPopScopeWidget({Key? key}) : super(key: key);

  @override
  _UseWillPopScopeWidgetState createState() => _UseWillPopScopeWidgetState();
}

class _UseWillPopScopeWidgetState extends State<UseWillPopScopeWidget> {
  DateTime? _lastPressedAt; // 上次点击时间

  @override
  Widget build(BuildContext context) {
    return WillPopScope(
      onWillPop: () async {
        if (_lastPressedAt == null ||
            DateTime.now().difference(_lastPressedAt!) >
                const Duration(seconds: 1)) {
          _lastPressedAt = DateTime.now();
          return false;
        }
        return true;
      },
      child: Container(
        alignment: Alignment.center,
        child: const Text("1秒内连续按两次返回键退出"),
      ),
    );
  }
}
