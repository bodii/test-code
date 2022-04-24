import 'dart:developer';

import 'package:flutter/material.dart';

class TextFieldWidget extends StatelessWidget {
  const TextFieldWidget({
    Key? key,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    var username, password;
    return Column(
      children: [
        TextField(
          autofocus: true,
          decoration: const InputDecoration(
            labelText: '用户名',
            hintText: '用户名或邮箱',
            prefixIcon: Icon(Icons.person),
          ),
          onChanged: (value) {
            username = value;
            log(username);
          },
        ),
        TextField(
          decoration: const InputDecoration(
            labelText: '密码',
            hintText: '您的登录密码',
            prefixIcon: Icon(Icons.lock),
          ),
          obscureText: true,
          onChanged: (value) {
            password = value;
            log(password);
          },
        ),
      ],
    );
  }
}
