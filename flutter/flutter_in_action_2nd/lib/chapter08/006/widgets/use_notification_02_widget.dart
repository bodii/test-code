import 'package:flutter/material.dart';

import 'my_notification.dart';

class UseNotification02Widget extends StatefulWidget {
  const UseNotification02Widget({Key? key}) : super(key: key);

  @override
  _UseNotificationWidgetState createState() => _UseNotificationWidgetState();
}

class _UseNotificationWidgetState extends State<UseNotification02Widget> {
  String _msg = '';

  @override
  Widget build(BuildContext context) {
    return NotificationListener<MyNotification>(
      onNotification: (notification) {
        debugPrint(notification.msg);
        return false;
      },
      child: NotificationListener<MyNotification>(
        onNotification: (notification) {
          setState(() {
            _msg += notification.msg + '  ';
          });
          return false;
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
      ),
    );
  }
}
