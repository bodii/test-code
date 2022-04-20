import 'package:flutter/cupertino.dart';

class CupertinoTestRoute extends StatelessWidget {
  const CupertinoTestRoute({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return CupertinoPageScaffold(
      navigationBar: const CupertinoNavigationBar(
        middle: Text('Cupertino Demo'),
      ),
      child: Center(
        child: CupertinoButton(
          child: const Text('Press'),
          color: CupertinoColors.activeBlue,
          onPressed: () {},
        ),
      ),
    );
  }
}
