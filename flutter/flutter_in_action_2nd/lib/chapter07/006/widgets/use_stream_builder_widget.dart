import 'package:flutter/material.dart';

class UseStreamBuilderWidget extends StatelessWidget {
  const UseStreamBuilderWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return StreamBuilder<int>(
        stream: counter(),
        builder: (BuildContext context, AsyncSnapshot<int> snapshot) {
          if (snapshot.hasError) {
            return Text('Error: ${snapshot.error}');
          }

          switch (snapshot.connectionState) {
            case ConnectionState.none:
              return const Text('没有stream');
            case ConnectionState.waiting:
              return const Text('等待数据...');
            case ConnectionState.active:
              return Text('active: ${snapshot.data}');
            case ConnectionState.done:
              return const Text('Stream 已关闭');
          }
        });
  }
}

Stream<int> counter() {
  return Stream.periodic(const Duration(seconds: 1), (i) {
    return i;
  });
}
