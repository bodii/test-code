void main() {
  final value = 0x22;
  final bitmask = 0x0f;

  assert((value & bitmask) == 0x02);
  assert((value & ~bitmask) == 0x20);
  assert((value | bitmask) == 0x2f);
  assert((value ^ bitmask) == 0x2d);
  assert((value << 4) == 0x220);
  assert((value >> 4) == 0x02);
  assert((value >>> 4) == 0x02);
  assert((-value >> 4) == -0x03);
  assert((-value >>> 4) > 0);
}
