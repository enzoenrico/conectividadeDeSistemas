const input = document.querySelector('#textarea')
const messages = document.querySelector('#messages')
const username = document.querySelector('#username')
const send = document.querySelector('#send')

const url = "wss://" + "localhost:9999" + "/ws";
const ws = new WebSocket(url);

ws.onmessage = (msg)=>{
    //quando receber mensagem fazer parse pra json
    insertMessage(JSON.parse(msg.data))
}

send.onclick = () =>{
    const message = {
        username: username.value,
        content: input.value
    }

    //broadcast message
    ws.send(JSON.stringify(message))
    input.value = ''
}

send.addEventListener("keypress", (e)=>{
    if(e.key === "Enter"){
        const message = {
            username: username.value,
            content: input.value
        }
    
        //broadcast message
        ws.send(JSON.stringify(message))
        input.value = ''
    }
})

function insertMessage(messageObj){
    //create message div
    const message = document.createElement('div')
    message.setAttribute('class', 'chat-message')
    
    //quando mensagem e recebida
    console.log(`[${messageObj.username}]: ` + messageObj.content)
    message.textContent = `$[{messageObj.username}]: ${messageObj.content}`

    messages.appendChild(message)

    //primeira mensagem do chat
    messages.insertBefore(message, messages.firstChild)
}