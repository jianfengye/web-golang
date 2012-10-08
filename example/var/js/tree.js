(function(){
    var tree = $("table[data-type=tree]");
    renderTree(tree);
    
    $(tree).children("div[data-type=tree-parent]").click(function(event){
        var target = event.target;
        $(target).parent().children("div[data-type=tree-child]").toggle();
    })
    
    function renderTree(tree) {
        var dataurl = $(tree).attr("tree-data");
        $.ajax({ 
            url: dataurl, 
            dataType: "json",
            success:function(msg){
                var innerContent = "";
                var parentid = 0;
                for(key in msg) {
                    parentid++;
                    var iframeTarget = "#";
                    if (msg[key]['iframeTarget'] != "") {
                        iframeTarget = msg[key]['iframeTarget'];
                    }
                    innerContent = innerContent + '<tr data-type="tree-parent" parent-id="'
                        + parentid
                        +'" class="left_tr"><th class="left_th"><div class="left_item"><a href="#" tree-role="parentlink" iframe-target="'
                        + iframeTarget
                        +'" >'
                        + key 
                        +'</a></div></th></tr>';
                    if(typeof(msg[key]) == "object") {
                        for(key2 in msg[key]) {
                            if (key2 == "child") {
                                for(key3 in msg[key][key2]) {
                                    innerContent = innerContent + '<tr data-type="tree-child" data-parent="'
                                    + parentid
                                    + '" class="left_tr"><th class="left_th_2"><div class="left_item"><a href="/admin/" tree-role="childlink" iframe-target="'
                                    + msg[key][key2][key3]
                                    +'">'
                                    + key3
                                    + '</a></div></th></tr>';
                                }
                            }
                        }
                    }
                }
                $(tree).html(innerContent);
                $("tr[data-type=tree-child]").hide();
                
                // 增加toggle
                $("a[tree-role=parentlink]").click(function(event){
                    var target = event.target;
                    $("tr[data-type=tree-child]").hide();
                    var targettr = $(target).parents("tr[data-type=tree-parent]").first();
                    var nodeid = $(targettr).attr("parent-id");
                    $("tr[data-type=tree-child][data-parent="+ nodeid +"]").toggle();
                });
                
                $("a[iframe-target]").click(function(event){
                    event.preventDefault();
                    var target = event.target;
                    var targetsrc = $(target).attr("iframe-target");
                    var targetTree = $(target).parents("table[data-type=tree]").first();
                    var targetIframe = $(targetTree).attr("target-iframe");
                    $("#" + targetIframe).attr("src",targetsrc);
                });
                
                $("tr[data-type=tree-parent]").click(function(event){
                    event.preventDefault();
                    $("tr").removeClass("left_cur");
                    $(event.target).parents("tr[data-type=tree-parent]").addClass("left_cur");
                })
                
                $("tr[data-type=tree-child]").click(function(event){
                    event.preventDefault();
                    $("tr").removeClass("left_cur");
                    $(event.target).parents("tr[data-type=tree-child]").addClass("left_cur");
                    $(event.target).parents("tr[data-type=tree-child]").prevAll("tr[data-type=tree-parent]").first().addClass("left_cur");
                })
        }})
    }
})()