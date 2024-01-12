$(function(){
    $("[data-action=import]").click(function(){
        var params = $(".list-search").serialize()
        location.href = "import?"+params
    })
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
        $("input").attr("disabled", true)
        $("textarea").attr("disabled", true)
        $("select").attr("disabled", true)
        //$("button").attr("disabled", "disabled")
        $("button").hide()
        var notDisabled = $(".form-disabled-not")
        notDisabled.show()
        notDisabled.removeAttr("disabled")
        //$(".layui-footer .layui-btn").hide()
        layui.use(['form'], function(){
            var form = layui.form
            form.render('select');
            form.render('radio');

            setTimeout(function(){
                $(".layui-input").removeClass('layui-disabled')
                var  formRadio = $(".layui-form-radio")
                formRadio.removeClass('layui-radio-disbaled')
                formRadio.removeClass('layui-disabled')
            }, 300)
        })
    }

    var listSearch = $(".list-search")
    listSearch.find("input").each(function(k, v){
        $(v).val(getURLString($(v).attr("name")))
    })
    var listSearchSelect = listSearch.find("select")
    listSearchSelect.each(function(k, v){
        var value = getURLString($(v).attr("name"))
        $(v).find("option[value='"+value+"']").attr("selected",true);
    })
    if (listSearchSelect.length > 0) {
        layui.use(['form'], function(){
            var form = layui.form
            form.render('select');
        })
    }
    listSearch.find("button[type=reset]").click(function (){
        var hidden  = {}
        listSearch.find("input[type=hidden]").each(function(k, v){
            hidden[$(v).attr("name")] = $(v).val()
        })
        listSearch.find("input").val("")
        listSearchSelect.find("option").removeAttr("selected")
        for (var k in hidden) {
            $("[name="+k+"]").val(hidden[k])
        }
        if (listSearchSelect.length > 0) {
            layui.use(['form'], function(){
                var form = layui.form
                form.render('select');
            })
        }
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
