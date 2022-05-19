import 'package:flutter/material.dart';

class UseSnapAppBar extends StatelessWidget {
  const UseSnapAppBar({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: NestedScrollView(
        headerSliverBuilder: (BuildContext context, bool innerBoxIsScrolled) {
          return [
            SliverOverlapAbsorber(
              handle: NestedScrollView.sliverOverlapAbsorberHandleFor(context),
              sliver: SliverAppBar(
                floating: true,
                snap: true,
                expandedHeight: 200,
                forceElevated: innerBoxIsScrolled,
                flexibleSpace: FlexibleSpaceBar(
                  background: Image.asset(
                    'assets/images/sea.png',
                    fit: BoxFit.cover,
                  ),
                ),
              ),
            ),
          ];
        },
        body: Builder(builder: (BuildContext context) {
          return CustomScrollView(
            slivers: [
              SliverOverlapInjector(
                handle:
                    NestedScrollView.sliverOverlapAbsorberHandleFor(context),
              ),
              buildSliverList(100),
            ],
          );
        }),
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
