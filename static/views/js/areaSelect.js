
$(function(){
    var provinceEl = $(".province")
    var cityClass = provinceEl.attr("data-city-class") || "city"
    var regionClass = provinceEl.attr("data-region-class") || "region"
    var cityEl = $("."+cityClass)
    var regionEl = $("."+regionClass)

    layui.use(['form'], function(){
        provinceEl.change(function(event, cityId, regionId){
            var id = $(this).find("option:selected").val()
            $.get("/areas/list", {PID:id}, function(resp){
                if (resp.Code != ReqStatusSuccess) {
                    alertError(resp.Msg)
                    return
                }
                setAresSelectOption(cityEl, resp.Data, cityId)
                if (regionId > 0) {
                    cityEl.trigger("change", [regionId])
                }
            })
        })
        cityEl.change(function(event, _id){
            var id = $(this).find("option:selected").val()
            $.get("/areas/list", {PID:id}, function(resp){
                if (resp.Code != ReqStatusSuccess) {
                    alertError(resp.Msg)
                    return
                }
                setAresSelectOption(regionEl, resp.Data, _id)
            })
        })
        var form = layui.form
        var selectedProvinceID = provinceEl.find("option:selected").val()
        if (selectedProvinceID > 0) {
            provinceEl.trigger("change", [provinceEl.attr("data-city-id"), provinceEl.attr("data-region-id")])
        }
        form.on('select(province)', function (data) {
            provinceEl.trigger("change", [provinceEl.attr("data-city-id"), provinceEl.attr("data-region-id")])
        })
        form.on('select('+cityClass+')', function (data) {
            cityEl.trigger("change", [provinceEl.attr("data-city-id"), provinceEl.attr("data-region-id")])
        })
        form.on('select('+regionClass+')', function (data) {
            regionEl.trigger("change", [provinceEl.attr("data-city-id"), provinceEl.attr("data-region-id")])
        })
    })
})


function setAresSelectOption(el, data, id) {
    el.empty()
    el.append("<option value='0'>请选择</option>")
    data.forEach(function(v){
        if (v.ID == id) {
            el.append("<option value='"+v.ID+"' selected>"+v.Name+"</option>")
        } else {
            el.append("<option value='"+v.ID+"'>"+v.Name+"</option>")
        }
    })
    layui.use(['form'], function(){
        var form = layui.form
        form.render('select');
    })
}