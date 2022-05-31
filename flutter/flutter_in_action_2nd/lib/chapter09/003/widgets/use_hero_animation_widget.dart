import 'package:flutter/material.dart';

import 'hero_animation_widget_b.dart';

class UseHeroAnimationWidget extends StatelessWidget {
  const UseHeroAnimationWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      alignment: Alignment.topCenter,
      child: Column(
        children: [
          InkWell(
            child: Hero(
              tag: 'avatar', // 唯一标记，前后两个路由页Hero的tag必须相同
              child: ClipOval(
                child: Image.asset(
                  'assets/images/avatar.png',
                  width: 50.0,
                ),
              ),
            ),
            onTap: () {
              // 打开B路由
              Navigator.push(
                context,
                PageRouteBuilder(
                    pageBuilder: ((context, animation, secondaryAnimation) {
                  return FadeTransition(
                    opacity: animation,
                    child: Scaffold(
                      appBar: AppBar(title: const Text('原图')),
                      body: const HeroAnimationWidgetB(),
                    ),
                  );
                })),
              );
            },
          ),
          const Padding(
            padding: EdgeInsets.only(top: 8.0),
            child: Text('点击头像'),
          ),
        ],
      ),
    );
  }
}
