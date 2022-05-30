import 'package:flutter/material.dart';

import 'home_page.dart';

Map<String, Widget Function(BuildContext)> routes = {
  '/': (context) => const HomePage()
};
