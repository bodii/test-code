import 'package:flutter/material.dart';
import 'package:flutter_in_action_2nd/chapter02/003/echo_route.dart';
import 'package:flutter_in_action_2nd/chapter02/003/home_page02.dart';
import 'package:flutter_in_action_2nd/chapter02/003/new_route.dart';
import 'package:flutter_in_action_2nd/chapter02/003/tip_route.dart';

Map<String, Widget Function(BuildContext)> routes = <String, WidgetBuilder>{
  '/': (context) => const HomePage02(),
  'echo_page': (context) => const EchoRoute(),
  'tip2': (context) {
    return TipRoute(
        text: ModalRoute.of(context)!.settings.arguments.toString());
  },
};
