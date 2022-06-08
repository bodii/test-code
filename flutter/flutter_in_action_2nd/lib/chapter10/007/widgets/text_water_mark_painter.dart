import 'package:flutter/material.dart';
import 'dart:math' as math;
import 'dart:ui' as ui;

import 'water_mark.dart';

class TextWaterMarkPainter extends WaterMarkPainter {
  TextWaterMarkPainter({
    Key? key,
    double? rotate,
    EdgeInsets? padding,
    TextStyle? textStyle,
    required this.text,
    this.textDirection = TextDirection.ltr,
  })  : assert(rotate == null || rotate >= -90 && rotate <= 90),
        rotate = rotate ?? 0,
        padding = padding ?? const EdgeInsets.all(10.0),
        textStyle = textStyle ??
            const TextStyle(
              color: Color.fromARGB(20, 0, 0, 0),
              fontSize: 14,
            );

  double rotate; // 文本旋转的度数，是角度不是弧度
  TextStyle textStyle; // 文本样式
  EdgeInsets padding; // 文本的padding
  String text; // 文本
  TextDirection textDirection;

  @override
  Size paintUnit(Canvas canvas, double devicePixelRatio) {
    /*
    // 根据屏幕devicePixelRatio对文本样式中长度相关的一些值乘以devicePixelRatio
    final _textStyle = _handleTextStyle(textStyle, devicePixelRatio);
    final _padding = padding * devicePixelRatio;

    // 构建文本段落
    final builder = ui.ParagraphBuilder(_textStyle.getParagraphStyle(
      textDirection: textDirection,
      textAlign: TextAlign.start,
      textScaleFactor: devicePixelRatio,
    ));

    // 添加要绘制的文本及样式
    builder
      ..pushStyle(_textStyle.getTextStyle())
      ..addText(text);

    // layout 文本占用的空间
    ui.Paragraph paragraph = builder.build()
      ..layout(const ui.ParagraphConstraints(width: double.infinity));

    // 文本占用的真实宽度
    final textWidth = paragraph.longestLine.ceilToDouble();
    // 文本占用的真实高度
    final fontSize = paragraph.height;

    // 绘制文本
    canvas.drawParagraph(paragraph, Offset.zero);
    */

    final _textStyle = _handleTextStyle(devicePixelRatio);
    final _padding = padding * devicePixelRatio;
    // 构建文本画笔
    TextPainter painter = TextPainter(
      textDirection: TextDirection.ltr,
      textScaleFactor: devicePixelRatio,
    );

    // 添加文本和样式
    painter.text = TextSpan(text: text, style: _textStyle);
    // 对文本进行布局
    painter.layout();

    // 文本占用真实宽度
    final textWidth = painter.width;
    // 文本占用真实高度
    final textHeight = painter.height;

    // 将弧度转化为度数
    final radians = math.pi * rotate / 180;

    // 通过三角函数计算旋转后的位置和size
    final orgSin = math.sin(radians);
    final sin = orgSin.abs();
    final cos = math.cos(radians).abs();

    final width = textWidth * cos;
    final height = textWidth * sin;
    final adjustWidth = textHeight * sin;
    final adjustHeight = textHeight * cos;

    if (orgSin >= 0) {
      // 旋转角度为正
      canvas.translate(
        adjustWidth + _padding.left,
        _padding.top,
      );
    } else {
      // 旋转角度为负
      canvas.translate(
        _padding.left,
        height + _padding.top,
      );
    }

    canvas.rotate(radians);

    // 绘制文本
    painter.paint(canvas, Offset.zero);
    // 返回水印单元所占的真实空间大小

    return Size(
      width + adjustWidth + _padding.horizontal,
      height + adjustHeight + padding.vertical,
    );
  }

  TextStyle _handleTextStyle(double devicePixelRatio) {
    var style = textStyle;
    double _scale(attr) => attr == null ? 1.0 : devicePixelRatio;
    return style.apply(
      decorationThicknessFactor: _scale(style.decorationThickness),
      letterSpacingFactor: _scale(style.letterSpacing),
      wordSpacingFactor: _scale(style.wordSpacing),
      heightFactor: _scale(style.height),
    );
  }

  @override
  bool shouldRepaint(TextWaterMarkPainter oldPainter) {
    return oldPainter.rotate != rotate ||
        oldPainter.text != text ||
        oldPainter.padding != padding ||
        oldPainter.textDirection != textDirection ||
        oldPainter.textStyle != textStyle;
  }
}
