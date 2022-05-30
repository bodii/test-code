import 'package:flutter/material.dart';

class UseNotificationListenerWidget extends StatelessWidget {
  const UseNotificationListenerWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return NotificationListener(
      onNotification: (notification) {
        switch (notification.runtimeType) {
          case ScrollStartNotification:
            debugPrint('开始滚动');
            break;
          case ScrollUpdateNotification:
            debugPrint('正在滚动');
            break;
          case ScrollEndNotification:
            debugPrint('滚动停止');
            break;
          case OverscrollNotification:
            debugPrint('滚动到边界');
            break;
        }

        return false;
      },
      child: ListView.builder(
        itemCount: 100,
        itemBuilder: (context, index) {
          return ListTile(title: Text('$index'));
        },
      ),
    );
  }
}
