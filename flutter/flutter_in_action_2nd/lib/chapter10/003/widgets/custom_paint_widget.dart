import 'package:flutter/material.dart';
import 'dart:math';

class CustomPaintWidget extends StatelessWidget {
  const CustomPaintWidget({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        children: [
          RepaintBoundary(
            child: CustomPaint(
              size: const Size(300, 300), // 指定画布大小
              painter: MyPainter(),
            ),
          ),
          // CustomPaint或刷新按钮添加RepaintBoundary, 防止棋盘内容重绘
          ElevatedButton(
            onPressed: () {},
            child: const Text('刷新'),
          ),
        ],
      ),
    );
  }
}

class MyPainter extends CustomPainter {
  @override
  void paint(Canvas canvas, Size size) {
    debugPrint('paint');

    var rect = Offset.zero & size;

    // 画棋盘
    drawChessboard(canvas, rect);

    // 画棋子
    drawPieces(canvas, rect);
  }

  @override
  bool shouldRepaint(CustomPainter oldDelegate) => false;
}

/// 实现棋盘绘制
void drawChessboard(Canvas canvas, Rect rect) {
  // 棋盘背景
  Paint paint = Paint()
    ..isAntiAlias = true
    ..style = PaintingStyle.fill
    ..color = const Color(0xFFDCC48C);
  canvas.drawRect(rect, paint);

  // 画棋盘风格
  paint
    ..style = PaintingStyle.stroke // 线
    ..color = Colors.black38
    ..strokeWidth = 1.0; // 描边的宽度

  // 画横线
  for (int i = 0; i <= 15; ++i) {
    double dy = rect.top + rect.height / 15 * i;
    //画线
    canvas.drawLine(Offset(rect.left, dy), Offset(rect.right, dy), paint);
  }

  // 画竖线
  for (int i = 0; i <= 15; ++i) {
    double dx = rect.left + rect.width / 15 * i;
    canvas.drawLine(Offset(dx, rect.top), Offset(dx, rect.bottom), paint);
  }
}

/// 实现棋子绘制
void drawPieces(Canvas canvas, Rect rect) {
  double eWidth = rect.width / 15;
  double eHeight = rect.height / 15;

  // 画一个黑子
  Paint paint = Paint()
    ..style = PaintingStyle.fill
    ..color = Colors.black;
  canvas.drawCircle(
    Offset(rect.center.dx - eWidth / 2, rect.center.dy - eHeight / 2),
    min(eWidth / 2, eHeight / 2) - 2,
    paint,
  );

  paint.color = Colors.white;
  canvas.drawCircle(
    Offset(rect.center.dx + eWidth / 2, rect.center.dy - eHeight / 2),
    min(eWidth / 2, eHeight / 2) - 2,
    paint,
  );
}
