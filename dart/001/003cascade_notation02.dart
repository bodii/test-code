void main() {
  final addressBook = (AddressBookBuilder()
        ..name = 'jenny'
        ..email = 'jenny@example.com'
        ..phone = (PhoneNumberBuilder()
              ..number = '415-555-0100'
              ..label = 'home')
            .build())
      .build();
}

class AddressBookBuilder {
  String name = '';
  String email = '';
  String phone = '';

  String build() {
    return 'name: ${name}, email: $email, phone: $phone';
  }
}

class PhoneNumberBuilder {
  String number = '';
  String label = '';

  String build() {
    return 'number: ${number}, label: $label';
  }
}
