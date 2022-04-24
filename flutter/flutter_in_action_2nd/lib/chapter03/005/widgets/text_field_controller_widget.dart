import 'dart:developer';

import 'package:flutter/material.dart';

class TextFieldControllerWidget extends StatefulWidget {
  const TextFieldControllerWidget({Key? key}) : super(key: key);

  @override
  State<TextFieldControllerWidget> createState() =>
      _TextFieldControllerWidgetState();
}

class _TextFieldControllerWidgetState extends State<TextFieldControllerWidget> {
  final TextEditingController _unameController = TextEditingController();
  final TextEditingController _pwordController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        TextField(
          autofocus: true,
          decoration: const InputDecoration(
            labelText: '用户名',
            hintText: '用户名或邮箱',
            prefixIcon: Icon(Icons.person),
          ),
          controller: _unameController,
          onChanged: (value) {
            log(_unameController.text);
          },
        ),
        TextField(
          decoration: const InputDecoration(
            labelText: '密码',
            hintText: '您的登录密码',
            prefixIcon: Icon(Icons.lock),
          ),
          obscureText: true,
          controller: _pwordController,
          onChanged: (value) {
            log(_pwordController.text);
          },
        ),
      ],
    );
  }
}
