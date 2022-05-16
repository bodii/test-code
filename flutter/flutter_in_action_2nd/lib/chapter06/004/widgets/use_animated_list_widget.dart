import 'dart:developer';

import 'package:flutter/material.dart';

class UseAnimatedListWidget extends StatefulWidget {
  const UseAnimatedListWidget({Key? key}) : super(key: key);

  @override
  _UseAnimatedListWidgetState createState() => _UseAnimatedListWidgetState();
}

class _UseAnimatedListWidgetState extends State<UseAnimatedListWidget> {
  var data = <String>[];
  int counter = 5;
  final globalKey = GlobalKey<AnimatedListState>();

  @override
  void initState() {
    for (var i = 0; i < counter; i++) {
      data.add('${i + 1}');
    }
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Stack(
      children: [
        AnimatedList(
          key: globalKey,
          initialItemCount: data.length,
          itemBuilder:
              (BuildContext context, int index, Animation<double> animation) {
            // 添加列表项时会执行渐显动画
            return FadeTransition(
              opacity: animation,
              child: buildItem(context, index),
            );
          },
        ),
        buildAddBtn(),
      ],
    );
  }

  // 创建一个“＋”按钮，点击后会向列表中插入一项
  Widget buildAddBtn() {
    return Positioned(
      child: FloatingActionButton(
        child: const Icon(Icons.add),
        onPressed: () {
          data.add('${++counter}');
          globalKey.currentState!.insertItem(data.length - 1);
          log('添加 $counter');
        },
      ),
      bottom: 30,
      left: 0,
      right: 0,
    );
  }

  // 构建列表项
  Widget buildItem(context, index) {
    String char = data[index];
    return ListTile(
      key: ValueKey(char),
      title: Text(char),
      trailing: IconButton(
        icon: const Icon(Icons.delete),
        onPressed: () => onDelete(context, index),
      ),
    );
  }

  void onDelete(context, index) {
    setState(() {
      globalKey.currentState!.removeItem(
        index,
        (context, animation) {
          var item = buildItem(context, index);
          log('删除 ${data[index]}');
          data.removeAt(index);
          // 删除动画是一个合成动画；渐陷 ＋ 缩小列表项告诉
          return FadeTransition(
            opacity: CurvedAnimation(
              parent: animation,
              curve: const Interval(0.5, 1.0),
            ),
            child: SizeTransition(
              sizeFactor: animation,
              axisAlignment: 0.0,
              child: item,
            ),
          );
        },
        duration: const Duration(milliseconds: 200),
      );
    });
  }
}
