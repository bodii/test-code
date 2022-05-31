import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';
import 'package:flutter/scheduler.dart';

class CustomHeroAnimationWidget extends StatefulWidget {
  const CustomHeroAnimationWidget({Key? key}) : super(key: key);

  @override
  _CustomHeroAnimationWidgetState createState() =>
      _CustomHeroAnimationWidgetState();
}

class _CustomHeroAnimationWidgetState extends State<CustomHeroAnimationWidget>
    with SingleTickerProviderStateMixin {
  late AnimationController _controller;
  bool _animating = false;
  AnimationStatus? _lastAnimationStatus;
  late Animation _animation;

  // 两个组件在Stack中的rect
  Rect? child1Rect;
  Rect? child2Rect;

  @override
  void initState() {
    _controller = AnimationController(
      duration: const Duration(milliseconds: 200),
      vsync: this,
    );
    _animation = CurvedAnimation(
      parent: _controller,
      curve: Curves.easeIn,
    );

    _controller.addListener(() {
      if (_controller.isCompleted || _controller.isDismissed) {
        if (_animating) {
          setState(() {
            _animating = false;
          });
        }
      } else {
        _lastAnimationStatus = _controller.status;
      }
    });

    super.initState();
  }

  Widget wChild1() {
    return GestureDetector(
      onTap: () {
        setState(() {
          _animating = true;
          _controller.forward();
        });
      },
      child: SizedBox(
          width: 50,
          child: ClipOval(child: Image.asset('assets/images/avatar.png'))),
    );
  }

  Widget wChild2() {
    return GestureDetector(
      onTap: () {
        setState(() {
          _animating = true;
          _controller.reverse();
        });
      },
      child: Image.asset(
        'assets/images/avatar.png',
        width: 400,
      ),
    );
  }

  Rect _getRect(RenderAfterLayout renderAfterLayout) {
    return renderAfterLayout.localToGlobal(
          Offset.zero,
          ancestor: context.findRenderObject(),
        ) &
        renderAfterLayout.size;
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    // 小头像
    final Widget child1 = wChild1();
    // 大头像
    final Widget child2 = wChild2();

    // 是否展示小头像；只有在动画执行时，初始状态或者刚从大图变小图时才应该显示小头像
    bool showChild1 =
        !_animating && _lastAnimationStatus != AnimationStatus.forward;

    // 执行动画时的目标组件，如果是从小图变为大图，则目标组件是大图；反之则是小图
    Widget targetWidget;
    if (showChild1 || _controller.status == AnimationStatus.reverse) {
      targetWidget = child1;
    } else {
      targetWidget = child2;
    }

    return LayoutBuilder(builder: (context, constraints) {
      return SizedBox(
        // 让Stack填满屏幕剩余空间
        width: constraints.maxWidth,
        height: constraints.maxHeight,
        child: Stack(
          alignment: AlignmentDirectional.topCenter,
          children: [
            if (showChild1)
              AfterLayout(
                callback: (value) => child1Rect = _getRect(value),
                child: child1,
              ),
            if (!showChild1)
              AnimatedBuilder(
                animation: _animation,
                builder: (context, child) {
                  final rect = Rect.lerp(
                    child1Rect,
                    child2Rect,
                    _animation.value,
                  );

                  return Positioned.fromRect(rect: rect!, child: child!);
                },
                child: targetWidget,
              ),

            // 用于测量child2的大小，设置为全透明并且不能响应事件
            IgnorePointer(
              child: Center(
                child: Opacity(
                  opacity: 0,
                  child: AfterLayout(
                    callback: (value) => child2Rect = _getRect(value),
                    child: child2,
                  ),
                ),
              ),
            ),
          ],
        ),
      );
    });
  }
}

class RenderAfterLayout extends RenderProxyBox {
  RenderAfterLayout(this.callback);

  ValueSetter<RenderAfterLayout> callback;

  @override
  void performLayout() {
    super.performLayout();
    SchedulerBinding.instance
        .addPostFrameCallback((timeStamp) => callback(this));
  }

  Offset get offset => localToGlobal(Offset.zero);

  Rect get rect => offset & size;
}

class AfterLayout extends SingleChildRenderObjectWidget {
  const AfterLayout({
    Key? key,
    required this.callback,
    Widget? child,
  }) : super(key: key, child: child);

  final ValueSetter<RenderAfterLayout> callback;

  @override
  RenderObject createRenderObject(BuildContext context) {
    return RenderAfterLayout(callback);
  }

  @override
  void updateRenderObject(
      BuildContext context, RenderAfterLayout renderObject) {
    renderObject.callback = callback;
    super.updateRenderObject(context, renderObject);
  }
}
