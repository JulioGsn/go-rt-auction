console.log("hello world!");

const ws = new WebSocket("ws://localhost:8000/ws");

function sendMessage(message) {
  if (!ws.OPEN) {
    console.error("Cannot send message beucause Connection is not working:");
    return;
  }

  ws.send(message);
}

function sendUiMessage() {
  const message = getMessageFromInput();
  if (!message) {
    throw new Error("Message is empty somehow");
  }
  sendMessage(message);
}

function addListenerToButton() {
  const btn = document.getElementById("send-message");
  console.log("legal", btn);
  btn.addEventListener("click", sendUiMessage);
}

function handleMessageEvent(message) {
  const list = document.getElementById("list-messages");
  const li = document.createElement("li");
  li.innerText = message;
  list.appendChild(li);
}

function webSocketConnection() {
  console.log("Connecting to websocket");
  ws.addEventListener("open", (event) => {
    sendMessage("Client connection opened");
    addListenerToButton();
  });
  ws.onclose = function (evt) {
    print("CLOSE");
    ws = null;
  };
  ws.addEventListener("message", (event) => {
    handleMessageEvent(event.data);
  });
  ws.addEventListener("error", (event) => {
    console.log("socket error", event);
    ws.close();
  });
}

function getMessageFromInput() {
  return document.getElementById("input-message").value;
}

webSocketConnection();
