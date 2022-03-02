void main() {
  while (true) {
    if (shutDownRequested()) break;
    processIncomingRequests();
  }

  for (int i = 0; i < 10; i++) {
    var candidate = i;
    if (candidate < 5) {
      continue;
    }

    print(i);
  }

  var command = 'OPEN';
  switch (command) {
    case 'CLOSED':
      print('CLOSED');
      break;
    default:
      print('unknown');
  }
}

bool shutDownRequested() {
  return true;
}

void processIncomingRequests() {}
