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

    if (getURLString("form-disabled") === "1" ) {
        $("input").attr("disabled", "disabled")
        $("textarea").attr("disabled", "disabled")
        $("select").attr("disabled", "disabled")
        $("button").attr("disabled", "disabled")
        $(".layui-footer .layui-btn").hide()
    }

    var listSearch = $(".list-search")
    listSearch.find("input").each(function(k, v){
        $(v).val(getURLString($(v).attr("name")))
    })
    listSearch.find("select").each(function(k, v){
        var value = getURLString($(v).attr("name"))
        $(v).find("option[value='"+value+"']").attr("selected",true);
    })
    var perms =  $("[data-perm]")
    if (perms.length > 0) {
        $.get("/admin/perms", {}, function(resp){
            if (resp.Code === 1) {
                return
            }
            if ($.inArray("*", resp.Data) !== -1) {
                return
            }
            perms.each(function(k, v){
                if ($.inArray($(v).attr("data-perm-val"), resp.Data) === - 1) {
                    $(v).hide()
                }
            })
        }, "json")
    }

    var actionSlice = $("[data-action=slice]")
    actionSlice.each(function(k, v){
        var content = trim($(v).text())
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

    var actionSort = $("[data-action=sort]")
    actionSort.each(function(k, v){
        var _field = $(v).attr("data-field")
        _sort = getURLString(_field)
        if (_sort !== "") {
            $(v).attr("lay-sort", _sort)
        }
    })
    actionSort.click(function(){
        var _sort = $(this).attr("lay-sort")
        var _field = $(this).attr("data-field")
        if (_sort === "" || _sort === undefined) {
            _sort = "asc"
        } else if (_sort === "asc") {
            _sort = "desc"
        } else {
            _sort = ""
        }
        window.location.href = changeURLArg(window.location.href, _field, _sort)
    })
})
