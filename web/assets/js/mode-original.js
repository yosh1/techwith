

/*
define(function(require, exports, module) {
"use strict";
*/

define("ace/mode/doc_comment_highlight_rules",["require","exports","module","ace/lib/oop","ace/mode/text_highlight_rules"], function(require, exports, module) {
"use strict";


var oop = require("../lib/oop");
var TextHighlightRules = require("ace/mode/text_highlight_rules").TextHighlightRules;

var MyNewHighlightRules = function() {
    // regexp must not have capturing parentheses. Use (?:) instead.
    // regexps are ordered -> the first match is used
   this.$rules = {
        "start" : [
            {
              token: "string.文字列",
              regrex: /　+/
            }
        ]
    };
};

oop.inherits(MyNewHighlightRules, TextHighlightRules);
exports.MyNewHighlightRules = MyNewHighlightRules;
});
