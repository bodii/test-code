import 'package:flutter/material.dart';

import '../../007/widgets/keep_alive_wrapper.dart';

class UseTabBarView02Widget extends StatelessWidget {
  const UseTabBarView02Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    List tabs = ['新闻', '历史', '图片'];
    return DefaultTabController(
      length: tabs.length,
      child: Scaffold(
        appBar: AppBar(
          title: const Text('TabBarView Demo 2'),
          bottom: TabBar(
            tabs: tabs.map((e) => Tab(text: e)).toList(),
          ),
        ),
        body: TabBarView(
          children: tabs.map((e) {
            return KeepAliveWrapper(
              child: Container(
                alignment: Alignment.center,
                child: Text(e, textScaleFactor: 5),
              ),
            );
          }).toList(),
        ),
      ),
    );
  }
}
