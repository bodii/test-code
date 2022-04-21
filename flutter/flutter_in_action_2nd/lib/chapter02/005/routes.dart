import 'package:flutter/cupertino.dart';
import 'home_page.dart';

Map<String, Widget Function(BuildContext)> routes =
    <String, Widget Function(BuildContext)>{
  '/': (context) => const HomePage(),
};
