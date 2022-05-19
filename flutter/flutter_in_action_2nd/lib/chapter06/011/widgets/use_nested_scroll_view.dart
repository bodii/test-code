import 'package:flutter/material.dart';

class UseNestedScrollView extends StatelessWidget {
  const UseNestedScrollView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Material(
      child: NestedScrollView(
        headerSliverBuilder: (BuildContext context, bool innerBoxIsScrolled) {
          return [
            SliverAppBar(
              title: const Text('嵌套ListView'),
              pinned: true,
              forceElevated: innerBoxIsScrolled,
            ),
            buildSliverList(),
          ];
        },
        body: ListView.builder(
          padding: const EdgeInsets.all(8),
          physics: const ClampingScrollPhysics(),
          itemCount: 30,
          itemBuilder: (context, index) {
            return SizedBox(
              child: Center(child: Text('Item $index')),
              height: 50,
            );
          },
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
