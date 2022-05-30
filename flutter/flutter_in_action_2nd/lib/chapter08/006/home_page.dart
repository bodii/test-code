import 'package:flutter/material.dart';

import 'widgets/use_notification_listener_widget.dart';
import 'widgets/use_notification_widget.dart';
import 'widgets/use_notification_02_widget.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('通知Notification')),
      body: const UseNotification02Widget(),
    );
  }
}
