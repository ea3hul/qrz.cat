App.factory('wsData', function($websocket,REST_API) {
  
  var urlbase = REST_API.WsHostname + ":" + REST_API.Port + "/ws";
  
  var dataStream = $websocket(urlbase);

  var msg = "";

  dataStream.onMessage(function(message) {
    msg = message;
  });

  var methods = {
    msg: msg,
    send: function(msg) {
      dataStream.send(msg);
    }
  };

  return methods;  
});