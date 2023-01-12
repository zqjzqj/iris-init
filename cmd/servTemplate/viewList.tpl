{% extends "layouts/main.html" %}
{% block content %}
<div class="layui-fluid">
    <div class="layui-row layui-col-space15">
        <div class="layui-col-md12">
            <div class="layui-card">
                <div class="layui-card-header">列表</div>
                <div class="layui-card-body">
                    <div class="layui-form layui-card-header layuiadmin-card-header-auto" style="margin-bottom: 10px;">
                        <form action="" class="list-search">
                            搜索：
                            <div class="layui-inline">
                                <input class="layui-input" name="ID" placeholder="ID">
                            </div>
                            <button class="layui-btn" type="submit">搜索</button>
                            <button class="layui-btn layui-btn-primary" type="reset">重置</button>
                        </form>
                    </div>

                    <div style="padding-bottom: 10px;">
                        <button data-action="open" href="/{{.Alias}}/item-view.html" data-title="添加" class="layui-btn layuiadmin-btn-list" >
                            添加
                        </button>
                    </div>
                    <table class="layui-table">
                        <colgroup>
                            <col width="150">
                            <col width="150">
                            <col width="200">
                            <col>
                        </colgroup>
                        <thead>
                        <tr>
                            <th>ID</th>
                            <th>创建时间</th>
                            <th>修改时间</th>
                            <th>操作</th>
                        </tr>
                        </thead>
                        <tbody>
                        {% for item in Data.List %}
                        <tr>
                            <td>{{print "{{item.ID}}"}}</td>
                            <td>{{print "{{item.CreatedAt}}"}}</td>
                            <td>{{print "{{item.UpdatedAt}}"}}</td>
                            <td>
                                <button data-perm data-perm-val="GET@/{{.Alias}}/item-view.html"
                                        data-action="open" href="/{{.Alias}}/item-view.html?ID={{print "{{item.ID}}"}}&form-disabled=1"
                                        data-title="详情"
                                        class="layui-btn layui-btn-primary layui-btn-sm">
                                    详情
                                </button>
                                <button data-perm data-perm-val="POST@/{{.Alias}}/edit"
                                        data-action="open" href="/{{.Alias}}/item-view.html?ID={{print "{{item.ID}}"}}"
                                        data-title="编辑账户"
                                        class="layui-btn layui-btn-primary layui-btn-sm">
                                    编辑
                                </button>
                                <button
                                        data-perm data-perm-val="POST@/{{.Alias}}/delete"
                                        data-action="del" href="/{{.Alias}}/delete"
                                        data-id="{{print "{{item.ID}}"}}" class="layui-btn layui-btn-primary layui-btn-sm">
                                    删除
                                </button>
                            </td>
                        </tr>
                        {% endfor %}
                        </tbody>
                    </table>
                    {% include "../partials/pager.html" %}
                </div>
            </div>
        </div>
    </div>
</div>
{% endblock %}