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
        }  else if (data !== null &&  typeof data['_open_url'] !== "undefined" && data['_open_url']) {
            var width = data['_open_url'].Width || "85%"
            var height = data['_open_url'].Height || "95%"
            var title = data['_open_url'].Title;
            var type = data['_open_url'].Type || 2;
            close=1;
            if(title==="false") {
                title=false;
                close=0;
            }
            top.layer.open({
                type: parseInt(type),
                title: title,
                shadeClose: false,
                shade: 0.3,
                maxmin: true, //开启最大化最小化按钮
                area: [width, height],
                closeBtn:close,
                content: data['_open_url'].Content,
            });
        }else {
            topIfrGo("", ifrIndex)
        }
    }, 1000);
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

    setTimeout(function(){
        if ( data !== null && typeof data['_url'] !== "undefined" && data['_url'] ) {
            return true
        } else if (data !== null &&  typeof data['_open_url'] !== "undefined" && data['_open_url']) {
            var width = data['_open_url'].Width || "85%"
            var height = data['_open_url'].Height || "95%"
            var title = data['_open_url'].Title;
            var type = data['_open_url'].Type || 2;
            close=1;
            if(title==="false") {
                title=false;
                close=0;
            }
            top.layer.open({
                type: parseInt(type),
                title: title,
                shadeClose: false,
                shade: 0.3,
                maxmin: true, //开启最大化最小化按钮
                area: [width, height],
                closeBtn:close,
                content: data['_open_url'].Content,
            });
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
    }, 1000)
}


function reloadParentLayerIframe(url) {
    var layerIfr = top.$(".layui-layer-iframe")
    var layerIndex = top.layer.getFrameIndex(window.name)
    for (i = 0; i < layerIfr.length; i++ ) {
        if (layerIndex == $(layerIfr[i]).attr("times") ) {
            if (i > 0) {
                _ifr = $(layerIfr[i-1]).find("iframe")
                if (url === undefined || url=== "") {
                    _url = _ifr[0].contentWindow.location.href
                    if (_url === undefined || _url === "") {
                        _url = _ifr.attr("src")
                    }
                    _ifr.attr("src", _url)
                } else {
                    _ifr.attr("src", url)
                }
            }
        }
    }
}

function reloadLayuiShowFrame(url) {
    if (url === "" || url === undefined) {
        url = top.$("#LAY_app_body .layui-show iframe")[0].contentWindow.location.href
        if (url === "" || url === undefined) {
            url = top.$("#LAY_app_body .layui-show iframe").attr('src')
        }
    }
    top.$("#LAY_app_body .layui-show iframe").attr('src', url)
}


//顶部ifr刷新或跳转
function topIfrGo(url, ifrIndex){
    if (ifrIndex !== undefined && ifrIndex !== "") {
        var closeAll = false
        var _ifrIndex = ifrIndex.split("-")
        ifrIndex = _ifrIndex[0]
        if (_ifrIndex.length > 1 && _ifrIndex[1] === "closeAll") {
            closeAll = true
        }
        //  var layerIndex = top.layer.getFrameIndex(window.name)
        if (ifrIndex === 0 || ifrIndex === "0") {//刷新当前 - 不是当前打开的框架 而是打开最上层框架的框架
            reloadLayuiShowFrame(url)
        } else if ( ifrIndex === "both" ) {//刷新当前和父级框架
            reloadParentLayerIframe(url)
            reloadLayuiShowFrame("")
        } else {//刷新父级框架
            reloadParentLayerIframe(url)
        }
        if (closeAll === true){
            top.layer.closeAll()
        } else {
            var layerIndex = top.layer.getFrameIndex(window.name)
            top.layer.close(layerIndex)
        }
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

//这里是在表单改变时 清空错误的提示数据 ---如果有的话
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
* 检测对象是否是空对象(不包含任何可读属性)。 //如你上面的那个对象就是不含任何可读属性
* 方法只既检测对象本身的属性，不检测从原型继承的属性。
*/
function isOwnEmpty(obj) {
    for(var name in obj) {
        if(obj.hasOwnProperty(name)) {
            return false;
        }
    }
    return true;
};