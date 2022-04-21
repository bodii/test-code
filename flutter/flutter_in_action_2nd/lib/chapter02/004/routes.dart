import 'package:flutter/material.dart';
import 'home_page.dart';

Map<String, Widget Function(BuildContext)> routes = <String, WidgetBuilder>{
  '/': (context) => const HomePage(),
};
