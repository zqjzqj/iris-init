function Callback(resp, ifrIndex) {
    if (resp.Code != ReqStatusSuccess) {
        errorCallback(resp.Msg, resp.Data, ifrIndex)
        return
    }
    successCallback(resp.Msg, resp.Data, ifrIndex)
}

function successCallbackUrl(url, message, data, ifrIndex) {
    if ( data !== null ) {
        data["_url"] = url
    } else {
        data = {"_url":url}
    }
    successCallback(message, data, ifrIndex)
}

function errCallbackUrl(url, message, data, ifrIndex) {
    if ( data !== null ) {
        data["_url"] = url
    } else {
        data = {"_url":url}
    }
    errorCallback(message, data, ifrIndex)
}

function successCallback(message, data, ifrIndex) {
    alertSuccess(message);
    setTimeout(function () {
        if (data !== null &&  typeof data['_url'] !== "undefined" && data['_url'] ) {
            topIfrGo(data["_url"], ifrIndex)
        } else {
            topIfrGo("", ifrIndex)
        }
    }, 500);
}


function sendRequest(url, method, data, func, json) {
    var csrfContent = $("meta[name=csrf-token]").attr('content');
    loadingShow();
    var params = {
        url : url,
        method : method,
        data : data,
        dataType: 'json',
        headers: {"X-CSRF-Token":csrfContent},
        success: function(resp){
            loadingClose()
            func(resp)
        }
    }
    if (json) {
        params["contentType"] = "application/json;charset=utf-8"
    }
    $.ajax(params)
}

function sendRequestSync(url, method, data, func, json) {
    var csrfContent = $("meta[name=csrf-token]").attr('content');
    loadingShow();
    var params = {
        url : url,
        method : method,
        data : data,
        dataType: 'json',
        async:false,
        headers: {"X-CSRF-Token":csrfContent},
        success: function(resp){
            loadingClose()
            func(resp)
        }
    }
    if (json) {
        params["contentType"] = "application/json;charset=utf-8"
    }
    $.ajax(params)
}

function errorCallback(message, data, ifrIndex) {

    alertError(message)

    if ( data !== null && typeof data['_url'] !== "undefined" && data['_url'] ) {
        topIfrGo(data['_url'], ifrIndex)
        return true
    }

    if (data !== null &&  typeof data !== "undefined" && !isEmptyObject(data) ) {
        $.each(data, function (key, item) {
            if ( typeof item === "object" || typeof item === 'Array' ) {
                item = item[0]
            }

            $("[name='"+key+"']").parents(".form-group").addClass("has-error has-danger");/*
            $("[name='"+key+"']").addClass("has-danger")*/

            var errorElem= $("[data-mark='"+key+"']");
            errorElem.text(item);
            errorElem.css("color", "#a94442");

        });
    }
}


function reloadParentLayerIframe(url) {
    var layerIfr = top.$(".layui-layer-iframe")
    var layerIndex = top.layer.getFrameIndex(window.name)
    for (i = 0; i < layerIfr.length; i++ ) {
        if (layerIndex == $(layerIfr[i]).attr("times") ) {
            if (i > 0) {
                console.log("p000")
                _ifr = $(layerIfr[i-1]).find("iframe")
                if (url === undefined || url=== "") {
                    _ifr.attr("src", _ifr.attr("src"))
                } else {
                    _ifr.attr("src", url)
                }
            }
        }
    }
}

function reloadLayuiShowFrame(url) {
    if (url === "" || url === undefined) {
        url = top.$("#LAY_app_body .layui-show iframe").attr('src')
    }
    top.$("#LAY_app_body .layui-show iframe").attr('src', url)
}

//??????ifr???????????????
function topIfrGo(url, ifrIndex){
    if (ifrIndex !== undefined && ifrIndex !== "") {
        var layerIndex = top.layer.getFrameIndex(window.name)
        if (ifrIndex === 0 || ifrIndex === "0") {
            reloadLayuiShowFrame(url)
        } else if ( ifrIndex === "both" ) {
            reloadParentLayerIframe(url)
            reloadLayuiShowFrame("")
        } else {
            reloadParentLayerIframe(url)
        }
        top.layer.close(layerIndex)
        return
    }
    if (url === undefined || url=== "") {
        location.reload(true)
        return
    }
    location.href = url
}

function isEmptyObject(obj) {
    for (var i in obj) {
        return false
    }
    return true
}

//??????????????????????????? ??????????????????????????? ---???????????????
$("input").change(function(){
    var _this_div = $(this).parents(".form-group");
    _this_div.removeClass("has-error")
    _this_div.removeClass("has-danger")

    var name = $(this).attr("data-name");
    if ( !name ) {
        name = $(this).attr("name");
    }
    $("[data-mark='"+name+"']").text("")
});
$("button").click(function(){
    var _this_div = $(this).parents(".form-group");
    _this_div.removeClass("has-error")
    _this_div.removeClass("has-danger")

    var name = $(this).attr("data-name");
    if ( !name ) {
        name = $(this).attr("name");
    }
    $("[data-mark='"+name+"']").text("")
});

$("select").change(function(){
    var _this_div = $(this).parents(".form-group");
    _this_div.removeClass("has-error")
    _this_div.removeClass("has-danger")

    var name = $(this).attr("data-name");
    if ( !name ) {
        name = $(this).attr("name");
    }
    $("[data-mark='"+name+"']").text("")
});

$("input[type=checkbox]").click(function(){
    var _this_div = $(this).parents(".form-group");
    _this_div.removeClass("has-error")
    _this_div.removeClass("has-danger")

    var name = $(this).prop("data-name");
    if ( !name ) {
        name = $(this).prop("name");
    }
    $("[data-mark='"+name+"']").text("")
});

$("textarea").change(function(){
    var _this_div = $(this).parents(".form-group");
    _this_div.removeClass("has-error")
    _this_div.removeClass("has-danger")

    var name = $(this).attr("data-name");
    if ( !name ) {
        name = $(this).attr("name");
    }
    $("[data-mark='"+name+"']").text("");
});


/*
* ??????????????????????????????(???????????????????????????)??? //?????????????????????????????????????????????????????????
* ??????????????????????????????????????????????????????????????????????????????
*/
function isOwnEmpty(obj) {
    for(var name in obj) {
        if(obj.hasOwnProperty(name)) {
            return false;
        }
    }
    return true;
};