var ws = new WebSocket("ws://localhost:8081/stream");

ws.onmessage = function (event) {
        var promise = decode.QIMDecode(event.data);
        promise.then(
            function(payload){
                floatframe.wm = payload;
            });
        queue.push(floatframe);
    };


