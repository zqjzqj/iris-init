$(function(){
    var _pageLimits = $(".layui-laypage-limits select")
    _pageLimits.find("option").each(function(k, v){
        if ($(v).val() === _pageLimits.attr("data-curr")) {
            $(v).attr("selected", "selected")
        }
    })
    _pageLimits.change(function(){
        var url = window.location.href;
        url = changeURLArg(url, "Size", $(this).val())
        window.location.href = url
    })

    $(".layui-laypage-skip .layui-laypage-btn").click(function(){
        var url = window.location.href;
        url = changeURLArg(url, "Page", parseInt($(".layui-laypage-skip .layui-input").val()))
        window.location.href = url
    })

    var listSearch = $(".list-search")
    listSearch.find("input").each(function(k, v){
        $(v).val(getURLString($(v).attr("name")))
    })

    var actionSlice = $("[data-action=slice]")
    actionSlice.each(function(k, v){
        var content = $(v).text()
        var maxLen = $(v).attr("data-length")
        if (maxLen < content.length) {
            $(v).text(cutStr(content, maxLen))
            $(v).attr("data-action", 'tips')
            $(v).attr("data-content", content)
        } else {
            $(v).text(content)
        }
        $(v).css("cursor", "pointer")
        $(v).dblclick(function(){
            layer.alert(content, {
                skin: 'layui-layer-molv', //样式类名
                closeBtn: 0,
                title:"文本详情"
            });
        });
    })
})
