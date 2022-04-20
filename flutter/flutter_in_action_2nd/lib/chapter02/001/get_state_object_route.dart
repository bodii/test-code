import 'package:flutter/material.dart';

class GetStateObjectRoute extends StatefulWidget {
  const GetStateObjectRoute({Key? key}) : super(key: key);

  @override
  _GetStateObjectRouteState createState() => _GetStateObjectRouteState();
}

class _GetStateObjectRouteState extends State<GetStateObjectRoute> {
  static final GlobalKey<ScaffoldState> _globalKey = GlobalKey();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      key: _globalKey,
      appBar: AppBar(title: const Text("子树中获取state对象")),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
          children: [
            Builder(
              builder: (context) {
                return ElevatedButton(
                  child: const Text('打开抽屉菜单1'),
                  onPressed: () {
                    ScaffoldState _state =
                        context.findAncestorStateOfType<ScaffoldState>()!;
                    // 打开抽屉菜单
                    _state.openDrawer();
                  },
                );
              },
            ),
            Builder(
              builder: (context) {
                return ElevatedButton(
                  child: const Text('打开抽屉菜单2'),
                  onPressed: () {
                    ScaffoldState _state = Scaffold.of(context);
                    _state.openDrawer();
                  },
                );
              },
            ),
            Builder(
              builder: (context) {
                return ElevatedButton(
                  child: const Text('打开抽屉菜单3'),
                  onPressed: () {
                    _globalKey.currentState?.openDrawer();
                  },
                );
              },
            ),
            Builder(
              builder: (context) {
                return ElevatedButton(
                  child: const Text('我是SnackBar'),
                  onPressed: () {
                    ScaffoldMessenger.of(context).showSnackBar(
                      const SnackBar(
                        content: Text('我是SnackBar'),
                      ),
                    );
                  },
                );
              },
            ),
          ],
        ),
      ),
      drawer: const Drawer(),
    );
  }
}
