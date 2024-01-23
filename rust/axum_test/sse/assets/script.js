var eventSource = new eventSource("sse");

eventSource.onmessage = function (event) {
 console.log("Message from server ", event.data);
}
