import 'dart:developer';

import 'package:flutter/material.dart';

class UseScrollController01Widget extends StatefulWidget {
  const UseScrollController01Widget({Key? key}) : super(key: key);

  @override
  _UseScrollController01WidgetState createState() =>
      _UseScrollController01WidgetState();
}

class _UseScrollController01WidgetState
    extends State<UseScrollController01Widget> {
  ScrollController _controller = ScrollController();
  bool showToTopBtn = false;

  @override
  void initState() {
    super.initState();
    // 监听滚动事件，打印滚动位置
    _controller.addListener(() {
      log(_controller.offset.toString());
      if (_controller.offset < 1000 && showToTopBtn) {
        setState(() {
          showToTopBtn = false;
        });
      } else if (_controller.offset >= 1000 && showToTopBtn == false) {
        setState(() {
          showToTopBtn = true;
        });
      }
    });
  }

  @override
  void dispose() {
    // 为了避免内存泄露，需要调用_controller.dispose
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('滚动控制')),
      body: Scrollbar(
        child: ListView.builder(
          itemBuilder: (context, index) => ListTile(title: Text('$index')),
          itemCount: 100,
          itemExtent: 50.0,
          controller: _controller,
        ),
      ),
      floatingActionButton: !showToTopBtn
          ? null
          : FloatingActionButton(
              child: const Icon(Icons.arrow_upward),
              onPressed: () {
                // 返回到顶部时执行动画
                _controller.animateTo(
                  .0,
                  duration: const Duration(milliseconds: 200),
                  curve: Curves.ease,
                );
              },
            ),
    );
  }
}
