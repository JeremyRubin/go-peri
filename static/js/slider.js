$(document).ready(function(){
	var socket = (function(){
		/*
		Open up a socket for sending/receiving data.
		Only use one socket to minimize system resource consumption
		*/
		var s = new SockJS("http://"+location.host+"/socket");
		s.onopen = function(){};
		s.onclose = function(){};
		s.onmessage = function(evt){};
		return s
	})();

	var rsocket = (function(){
		/*
		Open up a socket for sending/receiving data.
		Only use one socket to minimize system resource consumption
		*/
		SockJS.prototype.regID = function(name, id){
			// register a list of ids locally. Prototyped to avoid closures
			idMap[name] = id;
			console.log(name+id);
		}
		var idMap = {};
		var sId = 0;
		var s = new SockJS("http://"+location.host+"/socket");
		s.onopen = function(){

		};
		s.onclose = function(){};
		s.onmessage = function(evt){
console.log(idMap);
			json = $.parseJSON(evt.data);
			console.log(json);
			$('#'+idMap[json.Name]).html(json.Message);
		};
		return s
	})();

	var make_slider = function(element, index, array){

		var parent = element.id;
		var name = $('#'+parent).data("device");
		var max = 100;
		var min = 0;
		$('#'+parent).append("<div class='well'><input style='display:none' type='text'  value='5' id='"+parent+"_slider''>"+name+"</div>");
		var slider = $('#'+parent+"_slider").slider({"min":min, "max":max});
		slider.on("slideStop", function(evt){
			socket.send(JSON.stringify({"Value":evt.value,"Name":name}));
		});
	}
	var make_toggle =function(element, index, array){

		var parent = element.id;
		var name = $('#'+parent).data("device");

		var state = false;
		var text = {true:name+" On", false: name+" Off"};
		var classes = {true:"btn-success", false: "btn-danger"};
		$('#'+parent).append("<div class='well'><a class='btn "+classes[state]+" span12' id='"+parent+"_btn'>"+text[state]+"</a></div>");
		var btn = $('#'+parent+'_btn');
		btn.click(function (evt) {
			state = ! state;
			btn.html(text[state]);
			btn.removeClass(classes[!state]).addClass(classes[state]);
            var v = 0;
            if(state){
                v =1;
            }
			socket.send(JSON.stringify({"State":[true, state],"Name":name,"Value":v }));
		});

	}
	var make_physical_slider = function(element, index, array){
		var parent = element.id;
		$('#'+parent).addClass('well');
		var name = $('#'+parent).data("device");
		rsocket.regID(name, parent);
		var rid = 0;
		var requestData = function(){
			rsocket.send(JSON.stringify({'Name':name}));
		}
		//rid = setInterval(requestData, 100);

	}

	var make_physical_rotator = function(element, index, array){
		console.log('hello');
		var parent = element.id;
		$('#'+parent).addClass('well');
		var name = $('#'+parent).data("device");
		rsocket.regID(name+'count', parent);
		var rid = 0;
		var requestData = function(){
			rsocket.send(JSON.stringify({'Name':name+'count'}));
			//rsocket.send(JSON.stringify({'name':name+'button'}));
		}
		//rid = setInterval(requestData, 100);
	}



	$(".toggles").toArray().forEach(make_toggle);
	$(".sliders").toArray().forEach(make_slider);
	$(".physical_sliders").toArray().forEach(make_physical_slider);
	$(".physical_rotators").toArray().forEach(make_physical_rotator);

	
});
