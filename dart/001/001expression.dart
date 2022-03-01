void main() {
  var s = 'string interpolation';

  assert('Dart has $s, which is very handy.' ==
      'Dart has string interpolation, '
          'which is very handy.');

  assert('That deserves all caps. '
          '${s.toUpperCase()} is very handy!' ==
      'That deserves all caps.'
          'STRING INTHRPOLATION is very handy!');

  s = '字符串插值';
  assert('Dart 有$s, 使用起来非常方便' == 'Dart 有字符串插值，使用起来非常方便');
  assert('使用${s.substring(3, 5)}表达式也非常方便' == '使用插值表达式也非常方便。');
}
