import 'package:flutter/material.dart';

import 'my_notification.dart';

class UseNotificationWidget extends StatefulWidget {
  const UseNotificationWidget({Key? key}) : super(key: key);

  @override
  _UseNotificationWidgetState createState() => _UseNotificationWidgetState();
}

class _UseNotificationWidgetState extends State<UseNotificationWidget> {
  String _msg = '';

  @override
  Widget build(BuildContext context) {
    return NotificationListener<MyNotification>(
      onNotification: (notification) {
        debugPrint(notification.msg);
        setState(() {
          _msg += notification.msg + '  ';
        });

        return true;
      },
      child: Center(
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: [
            Builder(builder: (context) {
              return ElevatedButton(
                onPressed: () => MyNotification('Hi').dispatch(context),
                child: const Text('Send Nofification'),
              );
            }),
            Text(_msg),
          ],
        ),
      ),
    );
  }
}
