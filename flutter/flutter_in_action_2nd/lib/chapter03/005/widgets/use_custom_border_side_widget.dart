import 'package:flutter/material.dart';

class UseCustomBorderSideWidget extends StatelessWidget {
  const UseCustomBorderSideWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: const TextField(
        keyboardType: TextInputType.emailAddress,
        decoration: InputDecoration(
          labelText: 'Email',
          hintText: '电子邮件地址',
          prefixIcon: Icon(Icons.email),
          border: InputBorder.none,
        ),
      ),
      decoration: BoxDecoration(
        border: Border(
          bottom: BorderSide(color: Colors.grey[200]!, width: 1.0),
        ),
      ),
    );
  }
}
