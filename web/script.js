console.log("hello world!");
get = document.getElementById.bind(document);

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

function handleMessageEvent(_message) {
  const message = JSON.parse(_message);
  if (message.type === "bidder") {
    const list = document.getElementById("list-bidders");
    const li = document.createElement("li");
    const { id, name } = message.data;
    li.innerText = `[${id}] - ${name}`;
    list.appendChild(li);
  }
}

function webSocketConnection() {
  console.log("Connecting to websocket");
  ws.addEventListener("open", (event) => {
    //sendMessage("Client connection opened");
    //addListenerToButton();
  });
  ws.onclose = function (evt) {
    console.log("CLOSE");
    ws = null;
  };
  ws.addEventListener("message", (event) => {
    console.log("data coming from the server: ", event.data);
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

async function getData() {
  const url = "http://localhost:8000/auction";
  try {
    const response = await fetch(url);
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }

    const json = await response.json();
    setItem(json);
    console.log(json);
  } catch (error) {
    console.error(error.message);
  }
}

function setItem(item) {
  console.log("set item", item);
  if (!item?.name || !item?.basePrice) {
    throw new Error(`Error in item`);
  }
  get("item-name").innerText = item.name;
  get("item-base-price").innerText = item.basePrice;
}

getData();
webSocketConnection();
