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
    if (getURLString("form-disabled") === "1" ) {
        $("input").attr("disabled", "disabled")
        $("textarea").attr("disabled", "disabled")
        $("select").attr("disabled", "disabled")
        $(".layui-footer .layui-btn").hide()
    }
    var perms =  $("[data-perms]")
    if (perms.length > 0) {
        $.get("/admin/perms", {}, function(resp){
            if (resp.Code === 0 && resp.Data == null) {
                return
            }
            perms.each(function(k, v){
                if ($.inArray($(v).attr("data-perm-val"), resp.Data) === - 1) {
                    $(v).hide()
                }
            })
        }, "json")
    }
})
