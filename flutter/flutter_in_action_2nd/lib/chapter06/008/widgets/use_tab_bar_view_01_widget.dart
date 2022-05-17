import 'package:flutter/material.dart';

import '../../007/widgets/keep_alive_wrapper.dart';

class UseTabBarView01Widget extends StatefulWidget {
  const UseTabBarView01Widget({Key? key}) : super(key: key);

  @override
  _UseTabBarView01WidgetState createState() => _UseTabBarView01WidgetState();
}

class _UseTabBarView01WidgetState extends State<UseTabBarView01Widget>
    with SingleTickerProviderStateMixin {
  late TabController _tabController;
  List tabs = ['新闻', '历史', '图片'];

  @override
  void initState() {
    super.initState();
    _tabController = TabController(length: tabs.length, vsync: this);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('TabBarView demo'),
        bottom: TabBar(
          controller: _tabController,
          tabs: tabs.map((e) => Tab(text: e)).toList(),
        ),
      ),
      body: TabBarView(
        controller: _tabController,
        children: tabs.map((e) {
          return KeepAliveWrapper(
            child: Container(
              alignment: Alignment.center,
              child: Text(
                e,
                textScaleFactor: 5,
              ),
            ),
          );
        }).toList(),
      ),
    );
  }

  @override
  void dispose() {
    // 释放资源
    _tabController.dispose();
    super.dispose();
  }
}
