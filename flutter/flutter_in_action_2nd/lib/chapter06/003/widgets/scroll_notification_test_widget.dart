import 'dart:developer';

import 'package:flutter/material.dart';

class ScrollNotificationTestWidget extends StatefulWidget {
  const ScrollNotificationTestWidget({Key? key}) : super(key: key);

  @override
  _ScrollNotificationTestWidgetState createState() =>
      _ScrollNotificationTestWidgetState();
}

class _ScrollNotificationTestWidgetState
    extends State<ScrollNotificationTestWidget> {
  String _progress = "0％"; // 保存进度百分比

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('ListView scroll notification')),
      body: Scrollbar(
        // 进度条
        // 监听滚动通知
        child: NotificationListener<ScrollNotification>(
          onNotification: (ScrollNotification notification) {
            double progress = notification.metrics.pixels /
                notification.metrics.maxScrollExtent;

            setState(() {
              _progress = "${(progress * 100).toInt()}%";
            });
            log("BottomEdge: ${notification.metrics.extentAfter == 0}");
            return false;
          },
          child: Stack(
            alignment: Alignment.center,
            children: [
              ListView.builder(
                itemCount: 100,
                itemExtent: 50.0,
                itemBuilder: (context, index) =>
                    ListTile(title: Text("$index")),
              ),
              CircleAvatar(
                radius: 30.0,
                child: Text(_progress),
                backgroundColor: Colors.black54,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
