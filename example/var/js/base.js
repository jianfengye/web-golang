$("#logout").click(function(event){
    event.preventDefault();
    del_cookie("admin_id");
    del_cookie("admin_name");
    window.location.href = "/login/index";
})

function del_cookie(name)
{
    document.cookie = name + '=; expires=Thu, 01 Jan 1970 00:00:01 GMT;path=/;';
}

$("form[data-type=formAction]").submit(function(event){
    event.preventDefault();
    var target = event.target;
    var answer = confirm($(target).attr("data-hint"));
    if (answer) {
        var action = $(target).attr("action");
        $.post(action, $(target).serialize(), function(ret){
            if(ret.Ret != "0") {
                alert(ret.Reason);
            } else {
                alert("操作成功");
                location.href = $(target).attr("data-rediret");
            }
        },"json")
    }
})

$("form[data-type=formActionMulti]").submit(function(event){
    event.preventDefault();
    var target = event.target;
    var answer = confirm($(target).attr("data-hint"));
    if (answer) {
        var action = $(target).attr("action");
        var formData = new FormData($('form')[0]);
        $.ajax({
            type : "POST",
            url : action,
            data : formData,
            contentType: false,
            cache: false,
            processData: false,
            dataType : "json",
            success : function(ret) {
                if(ret.Ret == "0") {
                    alert("操作成功");
                    location.href = $(target).attr("data-rediret");
                } else {
                    alert(ret.Reason);
                }
            }
        })
    }
})

$("input[data-type=redirectInput]").click(function(event){
    event.preventDefault();
    var target = event.target;
    window.location=$(target).attr("data-link");
})