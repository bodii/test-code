import 'package:flutter/material.dart';

import 'package:dio/dio.dart';

class UseDioWidget extends StatefulWidget {
  const UseDioWidget({Key? key}) : super(key: key);

  @override
  _UseDioWidgetState createState() => _UseDioWidgetState();
}

class _UseDioWidgetState extends State<UseDioWidget> {
  final Dio _dio = Dio();

  @override
  Widget build(BuildContext context) {
    return Container(
      alignment: Alignment.center,
      child: FutureBuilder(
        future: _dio.get('https://api.github.com/orgs/google/repos'),
        builder: (BuildContext context, AsyncSnapshot snapshot) {
          // 请求完成
          if (snapshot.connectionState == ConnectionState.done) {
            Response response = snapshot.data;
            // 发生错误
            if (snapshot.hasError) {
              return Text(snapshot.error.toString());
            }

            // 请求成功
            return ListView(
              children: response.data
                  .map(
                    (e) => ListTile(title: Text(e['full_name'])),
                  )
                  .toList(),
            );
          }

          // 请求未完成时弹出loading
          return const CircularProgressIndicator();
        },
      ),
    );
  }
}
