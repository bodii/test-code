import 'package:flutter/material.dart';
import 'dart:developer';

import 'sliver_flexible_header.dart';

class UseCustomScrollViewWidget extends StatelessWidget {
  const UseCustomScrollViewWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return CustomScrollView(
      physics:
          const BouncingScrollPhysics(parent: AlwaysScrollableScrollPhysics()),
      slivers: [
        SliverFlexibleHeader(
          visibleExtent: 200, // 初始状态在列表中占用的布局高度
          builder: (context, avaliableHeight, direction) {
            return GestureDetector(
              onTap: () => log('tap'),
              child: Image(
                image: const AssetImage('assets/images/avatar.png'),
                width: 50.0,
                height: avaliableHeight,
                alignment: Alignment.bottomCenter,
                fit: BoxFit.cover,
              ),
            );
          },
        ),
        // 构建一个list
        buildSliverList(30),
      ],
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
