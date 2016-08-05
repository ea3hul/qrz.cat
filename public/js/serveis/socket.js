App.factory('wsData', function($websocket,REST_API) {
  
  var urlbase = REST_API.WsHostname + ":" + REST_API.Port + "/ws";
  
  var dataStream = $websocket(urlbase);

  var collection = [];

  dataStream.onMessage(function(message) {
        collection.push(message.data);
  });

  var methods = {
    collection: collection,
    send: function(msg) {
      dataStream.send(msg);
    }
  };

  return methods;  
});