import 'package:flutter/material.dart';

class AnimatedDecoratedBox02 extends ImplicitlyAnimatedWidget {
  const AnimatedDecoratedBox02({
    Key? key,
    required this.decoration,
    required this.child,
    Curve curve = Curves.linear,
    required Duration duration,
  }) : super(
          key: key,
          curve: curve,
          duration: duration,
        );

  final BoxDecoration decoration;
  final Widget child;

  @override
  _AnimatedDecoratedBox02State createState() {
    return _AnimatedDecoratedBox02State();
  }
}

class _AnimatedDecoratedBox02State
    extends AnimatedWidgetBaseState<AnimatedDecoratedBox02> {
  dynamic _decoration;

  @override
  Widget build(BuildContext context) {
    return DecoratedBox(
      decoration: _decoration.evaluate(animation),
      child: widget.child,
    );
  }

  @override
  void forEachTween(visitor) {
    _decoration = visitor(
      _decoration,
      widget.decoration,
      (value) => DecorationTween(begin: value),
    );
  }
}
