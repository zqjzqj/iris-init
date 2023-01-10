{% extends "layouts/main.html" %}
{% block content %}
<div class="layui-fluid">
    <div class="layui-card">
        <div class="layui-card-body" style="padding: 15px;">
            <form class="layui-form" action="/{{.Alias}}/edit" method="post" data-action="form" data-ifr-index="0" lay-filter="component-form-group">
                <div class="layui-form-item">
                    <label class="layui-form-label">ID</label>
                    <div class="layui-input-block">
                        <input type="text" name="ID" value="{{printf "{{Data.item.ID}}"}}" disabled autocomplete="off" placeholder="ID" class="layui-input">
                    </div>
                </div>
                <div class="layui-form-item">
                    <label class="layui-form-label">创建时间</label>
                    <div class="layui-input-block">
                        <input type="text" name="CreatedAt" value="{{printf "{{Data.item.CreatedAt}}"}}" disabled autocomplete="off" placeholder="创建时间" class="layui-input">
                    </div>
                </div>
                <div class="layui-form-item">
                    <label class="layui-form-label">更新时间</label>
                    <div class="layui-input-block">
                        <input type="text" name="UpdatedAt" value="{{printf "{{Data.item.UpdatedAt}}"}}" disabled autocomplete="off" placeholder="更新时间" class="layui-input">
                    </div>
                </div>
                <div class="layui-form-item layui-layout-admin">
                    <div class="layui-input-block">
                        <div class="layui-footer" style="left: 0;">
                            <input type="hidden" name="ID" value="{{printf "{{Data.item.ID}}"}}">
                            <button class="layui-btn">立即提交</button>
                            <button type="reset" class="layui-btn layui-btn-primary">重置</button>
                        </div>
                    </div>
                </div>
            </form>
        </div>
    </div>
</div>
{% endblock %}