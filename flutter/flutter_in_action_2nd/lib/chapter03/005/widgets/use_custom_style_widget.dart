import 'package:flutter/material.dart';

class UseCustomStyleWidget extends StatelessWidget {
  const UseCustomStyleWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const TextField(
      decoration: InputDecoration(
        labelText: '请输入用户名',
        prefixIcon: Icon(Icons.person),
        // 未获得焦点下划线设为灰色
        enabledBorder: UnderlineInputBorder(
          borderSide: BorderSide(color: Colors.grey),
        ),
        // 获得焦点下划线设为蓝色
        focusedBorder: UnderlineInputBorder(
          borderSide: BorderSide(color: Colors.blue),
        ),
      ),
    );
  }
}
