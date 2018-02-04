if (!String.prototype.format) {
    String.prototype.format = function() {
        var str = this.toString();
        if (!arguments.length) {
            return str;
        }
        var args = typeof arguments[0];
        args = (("string" == args || "number" == args) ? arguments : arguments[0]);
        for (var arg in args) {
            str = str.replace(RegExp("\\{" + arg + "\\}", "gi"), args[arg]);
        }
        return str;
    };
}
$(document).ready(function() {
    var page = 1;
    function Init() {
	    $('#fullpage').fullpage({
            controlArrows: false,
            loopHorizontal: false,
            afterSlideLoad: function(anchorLink, index, slideAnchor, slideIndex) {
                if ((slideIndex + 1) % 10 == 0) {
                    page = page + 1;
                    $.get(
                        '/api/poem',
                        {page:page},
                        function(response){
                            var text = "";
                            response.data.forEach(function(item,index) {
                                text = text + $("#template").html().format(item.Title,
                                                                           item.Author,
                                                                           item.Paragraphs.reduce(function(pre,nex) {
                                                                               return '<p class="paragraphs">{0}</p>'.format(pre) + '<p class="paragraphs">{0}</p>'.format(nex)
                                                                           }));
                            });
                            if(slideIndex == 99) {
                                $(".section").html(text);
                                $.fn.fullpage.destroy('all');
                                $('.slide').eq(1).addClass('active');
                            }else {
                                $(".section").append(text);
                                $.fn.fullpage.destroy('all');
                                $('.slide').eq(slideIndex).addClass('active');
                            }
                            Init();
                        }, 'json');
                }
            }
        });
    }
    Init();
});
$(document).keydown(function (e) {
    var keycode = (e.keyCode ? e.keyCode : e.which);
    if (keycode == 32) {
        $.fn.fullpage.moveSlideRight();
    }
});
