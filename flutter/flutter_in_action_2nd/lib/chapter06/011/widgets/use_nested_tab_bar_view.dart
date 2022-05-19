import 'package:flutter/material.dart';

class UseNestedTabBarView extends StatelessWidget {
  const UseNestedTabBarView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    final List<String> _tabs = ['猜你喜欢', '今日特价', '发现更多'];

    return DefaultTabController(
      length: _tabs.length,
      child: NestedScrollView(
        headerSliverBuilder: (context, innerBoxIsScrolled) {
          return [
            SliverOverlapAbsorber(
              handle: NestedScrollView.sliverOverlapAbsorberHandleFor(context),
              sliver: SliverAppBar(
                  title: const Text('商城'),
                  floating: true,
                  snap: true,
                  forceElevated: innerBoxIsScrolled,
                  bottom: TabBar(
                    tabs: _tabs.map((String e) => Tab(text: e)).toList(),
                  )),
            ),
          ];
        },
        body: TabBarView(
          children: _tabs.map((String name) {
            return Builder(
              builder: (context) {
                return CustomScrollView(
                  key: PageStorageKey<String>(name),
                  slivers: [
                    SliverOverlapInjector(
                      handle: NestedScrollView.sliverOverlapAbsorberHandleFor(
                          context),
                    ),
                    SliverPadding(
                        padding: const EdgeInsets.all(8.0),
                        sliver: buildSliverList(50)),
                  ],
                );
              },
            );
          }).toList(),
        ),
      ),
    );
  }

  Widget buildSliverList([int count = 5]) {
    return SliverFixedExtentList(
      delegate: SliverChildBuilderDelegate(
        (context, index) {
          return ListTile(title: Text('$index'));
        },
        childCount: count,
      ),
      itemExtent: 50,
    );
  }
}
