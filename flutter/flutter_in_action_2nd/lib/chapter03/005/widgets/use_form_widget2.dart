import 'package:flutter/material.dart';

class UseFormWidget extends StatefulWidget {
  const UseFormWidget({Key? key}) : super(key: key);

  @override
  State<UseFormWidget> createState() => _UseFormWidgetState();
}

class _UseFormWidgetState extends State<UseFormWidget> {
  final TextEditingController _unameCon = TextEditingController();
  final TextEditingController _pwdCon = TextEditingController();
  final GlobalKey _formKey = GlobalKey<FormState>();

  @override
  Widget build(BuildContext context) {
    return Form(
      key: _formKey,
      autovalidateMode: AutovalidateMode.onUserInteraction,
      child: Column(
        children: [
          TextFormField(
            autofocus: true,
            controller: _unameCon,
            decoration: const InputDecoration(
              labelText: '用户名',
              hintText: '用户名或邮箱',
              icon: Icon(Icons.person),
            ),
            validator: (v) {
              return v!.trim().isNotEmpty ? null : '用户名不能为空';
            },
          ),
          TextFormField(
            controller: _pwdCon,
            decoration: const InputDecoration(
              labelText: '密码',
              hintText: '您的登录密码',
              icon: Icon(Icons.lock),
            ),
            obscureText: true,
            validator: (v) {
              return v!.trim().length > 5 ? null : '密码不能少于6位';
            },
          ),
          Padding(
            padding: const EdgeInsets.only(top: 28.0),
            child: Row(
              children: [
                Expanded(
                  child: Builder(builder: (context) {
                    return ElevatedButton(
                      child: const Padding(
                        padding: EdgeInsets.all(16.0),
                        child: Text('登录'),
                      ),
                      onPressed: () {
                        if (Form.of(context)!.validate()) {
                          // 验证通过提交数据
                        }
                      },
                    );
                  }),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
