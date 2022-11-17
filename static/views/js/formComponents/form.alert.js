function loadingShow() {
    layer.closeAll()
    layer.msg("加载中...", {
        icon:16,
        time: -1,
    })
}

function loadingClose() {
    layer.closeAll()
}

function alertSuccess(title, content, time) {
    layer.msg(title+(content?'<br>'+content:''),{time: time||1000}); //, {icon: 6}
}

function alertError(title, content, icon) {
    layer.alert(title+(content?'<br>'+content:''), { icon: icon||5 ,closeBtn: 0, title:false });
}

function alertWarning(title, content, func, func2) {
    layer.confirm(title+(content?'<br>'+content:''),{
        btn: ['确定', '取消'],
        title:false,closeBtn:0,icon:7,
        yes: function(index, layero){
            if ( typeof func !== "undefined" ) {
                func()
            }
            layer.close(index);
        },
        btn2:function(){
            typeof(func2)=='function' && func2();
        }
    });
}