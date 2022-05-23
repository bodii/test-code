import 'package:flutter/material.dart';

class UseFutureBuilderWidget extends StatelessWidget {
  const UseFutureBuilderWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Center(
      child: FutureBuilder<String>(
          future: mockNetworkData(),
          builder: (BuildContext context, AsyncSnapshot snapshot) {
            if (snapshot.connectionState == ConnectionState.done) {
              if (snapshot.hasError) {
                return Text('Error: ${snapshot.error}');
              } else {
                return Text('Contents: ${snapshot.data}');
              }
            } else {
              // 请求未结束，显示loading
              return const CircularProgressIndicator();
            }
          }),
    );
  }
}

Future<String> mockNetworkData() {
  return Future.delayed(const Duration(seconds: 2), () => "我是从互联网上获取的数据");
}
