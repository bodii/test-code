import 'package:flutter/material.dart';

class UseClip01Widget extends StatelessWidget {
  const UseClip01Widget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    Widget avatar = Image.asset('assets/images/avatar.png', width: 60.0);

    return Center(
      child: Column(
        children: [
          avatar,
          ClipOval(child: avatar), // 剪裁为圆形
          ClipRRect(
            // 剪裁为圆角矩形
            borderRadius: BorderRadius.circular(5.0),
            child: avatar,
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Align(
                alignment: Alignment.topLeft,
                widthFactor: .5,
                child: avatar,
              ),
              const Text(
                '你好世界',
                style: TextStyle(color: Colors.green),
              ),
            ],
          ),
          Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              ClipRect(
                // 将溢出部分剪裁
                child: Align(
                  alignment: Alignment.topLeft,
                  widthFactor: .5,
                  child: avatar,
                ),
              ),
              const Text('你好世界', style: TextStyle(color: Colors.green)),
            ],
          ),
        ],
      ),
    );
  }
}
