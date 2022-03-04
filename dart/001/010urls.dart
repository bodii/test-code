void main() {
  var uri = 'https://example.org/api?foo=some message';
  var encoded = Uri.encodeFull(uri);

  assert(encoded == 'https://example.org/api?foo=some%20message');

  var decoded = Uri.decodeFull(encoded);
  assert(uri == decoded);

  var uri2 = 'https://example.org/api?foo=some message';
  var encode2 = Uri.encodeComponent(uri2);
  assert(encode2 == 'https%3A%2F%2Fexample.org%2Fapi%3Ffoo%3Dsome%20message');

  var decoded2 = Uri.decodeComponent(encode2);
  assert(uri2 == decoded2);

  var uri3 = Uri.parse('https://example.org:8080/foo/bar#frag');
  assert(uri3.scheme == 'https');
  assert(uri3.host == 'example.org');
  assert(uri3.path == '/foo/bar');
  assert(uri3.fragment == 'frag');
  assert(uri3.origin == 'https://example.org:8080');

  var uri4 = Uri(
      scheme: 'https', host: 'example.org', path: '/foo/bar', fragment: 'frag');

  assert(uri4.toString() == 'https://example.org/foo/bar#frag');
}
