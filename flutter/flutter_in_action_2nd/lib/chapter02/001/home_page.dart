import 'package:flutter/material.dart';

import 'cupertino_test_route.dart';
import 'get_state_object_route.dart';

class HomePage extends StatelessWidget {
  const HomePage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return CupertinoTestRoute();
  }
}
