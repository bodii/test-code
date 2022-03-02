void main() {
  const isPublic = true;
  var visibility = isPublic ? 'public' : 'private';

  String playerName(String? name) => name ?? 'Guest';
  assert(playerName('alice') == 'alice');
  assert(playerName(null) == 'Guest');

  String playerName2(String? name) => name != null ? name : 'Guest';

  String playerName3(String? name) {
    if (name != null) {
      return name;
    }

    return 'Guest';
  }
}
