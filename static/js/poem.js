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
function Replace() {
    var points = ["，", "。"];
    var chars = ["一"];
    var brackets = [
        ["（", "︵"],
        ["）", "︶"],
        ["{", "︷"],
        ["}", "︸"],
        ["\\[", "︻"],
        ["\\]", "︼"],
        ["《", "︽"],
        ["》", "︾"],
        ["“", "﹃"],
        ["”", "﹄"],
        ["‘", "﹁"],
        ["’", "﹂"],
        ["——", "｜"]
    ];
    function replace(elements) {
        Array.from(elements).forEach(function(element) {
            var text = element.innerText;
            points.forEach(function(point) {
                text = text.replace(new RegExp(point, "g"),'<span class="poem-point">' + point + '</span>');
            });
            chars.forEach(function(char) {
                text = text.replace(new RegExp(char, "g"),'<span class="poem-char">' + char + '</span>');
            });
            brackets.forEach(function(bracket) {
                text = text.replace(new RegExp(bracket[0], "g"),bracket[1]);
            });
            element.innerHTML = text;
        });
    }
    var paragraphs = document.getElementsByClassName("poem-paragraph");
    var titles = document.getElementsByClassName("poem-title");
    replace(paragraphs);
    replace(titles);
}
$(document).ready(function() {
    var page = 1;
    function Init() {
        $('#fullpage').fullpage({
            controlArrows: false,
            loopHorizontal: false,
            afterSlideLoad: function(anchorLink, index, slideAnchor, slideIndex) {
                if ((slideIndex + 1) % 10 === 0) {
                    page = page + 1;
                    $.get(
                        '/api/poem',
                        {page:page},
                        function(response){
                            var text = "";
                            response.data.forEach(function(item,index) {
                                text = text + $("#template")
                                    .html()
                                    .format(item.title,
                                            item.author,
                                            item.paragraphs.reduce(function(pre,nex) {
                                                return '<p class="poem-paragraph">{0}</p>'.format(pre) + '<p class="poem-paragraph">{0}</p>'.format(nex);
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
    Replace();
});
$(document).keydown(function (e) {
    var keycode = (e.keyCode ? e.keyCode : e.which);
    if (keycode == 32) {
        $.fn.fullpage.moveSlideRight();
    }
});
