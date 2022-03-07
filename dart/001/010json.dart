import 'dart:convert';

void main() {
  test01();
  test02();
}

void test01() {
  var jsonString = '''
    [
      {"score": 40},
      {"score": 80}
    ]
  ''';

  var scores = jsonDecode(jsonString);
  assert(scores is List);

  var firstScore = scores[0];
  assert(firstScore is Map);
  assert(firstScore['score'] == 40);
  var jsonString = '''
	  {
		  "name": "jack",
		  "email": "jack@email.com",
	  }
  '''

  Map<String, dynamic> userMap = jsonDecode(jsonString);
  var user = User.fromJson(userMap);
  print('Howdy, ${user.name}!');
  print('We sent the verification link to ${user.email}.');
}

void test02() {
  var scores = [
    {'score': 40},
    {'score': 80},
    {'score': 100, 'overtime': true, 'special_guest': null}
  ];

  var jsonText = jsonEncode(scores);
  print(jsonText);
  assert(jsonText ==
      '[{"score":40},{"score":80},'
          '{"score":100,"overtime":true,'
          '"special_guest":null}]');
}

class User {
	final String name;
	final String email;

	User(this.name, this.email);

	User.fromJson(Map<String, dynamic> json)
		: name = json['name'],
		email = json['email'];

	Map<String, dynamic> toJson() => {
		'name' : name,
		'email':email,
	};
}
