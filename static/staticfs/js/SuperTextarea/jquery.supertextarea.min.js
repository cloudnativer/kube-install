/**
 * Supertextarea
 * Created by Explosion Pills <explosion-pills@aysites.com>
 * Report Bugs: <bugs@truthanduntruth.com>
 * Documentation: http://explosion-pills.com/development/jquery/plugins/supertextarea/
 * Copyright 2010
 */
(function ($) {
   $.fn.supertextarea = function (faith) {
      /**
       * The supertextarea
       */
      var area = $(this);
      /**
       * The container of the supertextarea
       */
      var cont = area.parent();

      /**
       * Defaults
       */
      var hope = {
         /**
          * Minimum width
          */
         minw: area.width()
         /**
          * maximum width
          */
         , maxw: cont.width()
         /**
          * minimum height
          */
         , minh: area.height()
         /**
          * maximum height
          */
         , maxh: cont.height()
         /**
          * replace tab key presses with actual tabs
          */
         , tabr: {
            /**
             * turn this setting on
             */
            use: true
            /**
             * change tabs into spaces
             */
            , space: true
            /**
             * this many spaces
             */
            , num: 3
         }
         /**
          * supertextarea css
          */
         , css: {'color': 'black'}
         /**
          * max length in characters
          */
         , maxl: 1000
         /**
          * Display the number of characters remaining
          */
         , dsrm: {
            /**
             * turn this setting on
             */
            use: true
            /**
             * The text displayed (e.g. 1000 Remaining)
             * You can use $ to insert the value anywhere
             * Otherwise, it is prepended
             */
            , text: 'Remaining'
            /**
             * css for the display remaining text
             */
            , css: {}
            /**
             * when true, clicking the "display remaining" text removes it
             */
            , rmv: false
         }
         /**
          * Minimum length required for form submission, in characters
          */
         , minl: 0
         /**
          * If true, leading/trailing whitespace is not counted for min length
          */
         , trim: true
         /**
          * Display to-go
          */
         , dstg: {
            /**
             * The text displayed.
             * @see dsrm
             */
            text: 'Required'
            /**
             * Css on the to-go div
             */
            , css: {}
            /**
             * If user tries to submit form with an inadequate number of characters,
             * scroll the textarea into view if this is true
             */
            , slide: true
         }
         /**
          * Placeholder text when supertextarea is empty
          */
         , plch: {
            /**
             * Use this setting
             */
            use: false
            /**
             * The actual placeholder text
             */
            , text: ''
            /**
             * Css of the supertextarea while placeholder is dislpayed
             */
            , css: {'color': 'gray'}
         }
      };
      var love = {};

      /**
       * Combined user settings
       */
      var justice = $.extend(love, hope, faith);

      /**#@+
       * Create sane minimum and maximum width and height
       */
      if (!justice.minh) {
         justice.minh = area.height();
      }
      if (!justice.minw) {
         justice.minw = area.width();
      }
      if (justice.maxh < justice.minh) {
         justice.maxh = justice.minh;
      }
      if (justice.maxw < justice.minw) {
         justice.maxw = justice.minw;
      }

      if (!justice.minh) {
         justice.minh = 0;
      }
      if (!justice.maxw) {
         justice.maxw = 0;
      }

      area.css(justice.css).height(justice.minh).width(justice.minw);
      /**#@-*/

      if (justice.tabr.use && justice.tabr.num < 1) {
         justice.tabr.num = 1;
      }

      var rep_css = [
         'paddingTop',
         'paddingRight',
         'paddingBottom',
         'paddingLeft',
         'fontSize',
         'lineHeight',
         'fontFamily',
         'fontWeight'
      ];

      if (typeof $.fn.supertextarea.counter == 'undefined') {
         $.fn.supertextarea.counter = 0;
      }
      var idcounter = $.fn.supertextarea.counter;
      $.fn.supertextarea.counter++;

      this.each(function () {
         if (this.type != 'textarea') {
            return false;
         }

         //"beholder" shadows the textarea to match size
         var beh = $('<div />').css({'position': 'absolute','display': 'none', 'word-wrap':'break-word'});
         //get the height of the line in pixels, from available source
         var line = parseInt(area.css('line-height')) || parseInt(area.css('font-size'));
         var goalheight = 0;

         beh.appendTo(area.parent());
         for (var i = 0; i < rep_css.length; i++) {
            beh.css(rep_css[i].toString(), area.css(rep_css[i].toString()));
         }
         beh.css('max-width', justice.maxw);

         /**
          * Update height of supertextarea
          */
         function eval_height(height, overflow){
            nh = Math.floor(parseInt(height));
            if (area.height() != nh) {
               area.css({'height': nh + 'px', 'overflow-y':overflow});
            }
         }

         /**
          * Update width of supertextarea
          */
         function eval_width(width, overflow) {
            nw = Math.floor(parseInt(width));
            if (area.width() != nw) {
               area.css({'width': nw + 'px', 'overflow-x': overflow});
            }
         }

         /**
          * Update the textarea size and its contents / indicators
          */
         function update(e) {
            if (justice.tabr.use && e) {
               tab_replace(e);
            }

            /**
             * Handle display-remaining
             */
            if (justice.dsrm.use && justice.maxl && !area.data('rmv')) {
               var dsm;
               if (!$("#textarea_dsrm" + area.data('partner')).length) {
                  dsm = $('<div></div>');
                  dsm.attr('id', "textarea_dsrm" + idcounter);
                  dsm.attr('class', "field_notice");
                  area.after(dsm);
                  area.data('partner', idcounter);
               }
               else {
                  dsm = $("#textarea_dsrm" + area.data('partner'));
               }
               var tl = area.data('plch') ? 0 : area.val().length;
               var txt = justice.maxl - tl;
               txt = txt < 0 ? 0 : txt;
               var rem = tl - justice.minl;
               var remtxt;
               var num;
               var msg;
               if (rem < 0 && justice.dstg.text != undefined) {
                  num = Math.abs(rem);
                  msg = justice.dstg.text;
                  if (justice.dstg.css != undefined) {
                     dsm.css(justice.dstg.css);
                  }
                  else if (justice.dsrm.css != undefined) {
                     dsm.css(justice.dsrm.css);
                  }
               }
               else {
                  num = txt;
                  msg = justice.dsrm.text;
                  if (justice.dsrm.css != undefined) {
                     dsm.css(justice.dsrm.css);
                  }
               }
               if (msg.match(/\$/)) {
                  remtxt = msg.replace('$', ' ' + num + ' ');
               }
               else {
                  remtxt = num + ' ' + msg;
               }
               dsm.text(remtxt);

               if (justice.dsrm.rmv) {
                  dsm.click(function () {
                     $(this).hide();
                     area.data('rmv', true);
                  });
               }
            }
            if (justice.maxl && justice.maxl - tl < 0) {
               area.val(area.val().substring(0, justice.maxl));
            }
            var ac = area.val().replace(/&/g,'&amp;').replace(/  /g, '&nbsp;&nbsp;').replace(/<|>/g, '&gt;').replace(/\n/g, '<br />');
            var bc = beh.html();

            if (ac + '&nbsp;' != bc) {
               beh.html(ac + '&nbsp;&nbsp;');
               if (Math.abs(beh.height() + line - area.height()) > 3
                  || Math.abs(beh.width() + line - area.width()) > 3
               ) {
                  var nh = beh.height() + line;
                  var maxh = justice.maxh;
                  var minh = justice.minh;
                  if (nh >= maxh) {
                     eval_height(maxh, 'auto');
                  }
                  else if (nh <= minh) {
                     eval_height(minh, 'hidden');
                  }
                  else {
                     eval_height(nh, 'hidden');
                  }

                  var nw = beh.width() + line;
                  var maxw = justice.maxw;
                  var minw = justice.minw;
                  if (nw >= maxw) {
                     eval_width(maxw, 'auto');
                  }
                  else if (nw <= minw) {
                     eval_width(minw, 'hidden');
                  }
                  else {
                     if (beh.height() + line > maxh) {
                        eval_width(nw + line, 'hidden');
                     }
                     else {
                        eval_width(nw, 'hidden');
                     }
                  }
               }
            }
         }

         /**
          * Prevent form submission if supertextarea text length is not in the correct limits
          */
         area.parents("form").submit(function (e) {
            var val;
            if (justice.trim) {
               val = $.trim(area.val());
            }
            else {
               val = area.val();
            }
            if (val.length < justice.minl || justice.minl > 0 && area.data('plch')) {
               if (justice.dstg.slide) {
                  $("html, body").animate({scrollTop: area.offset().top});
               }
               e.stopPropagation();
               e.preventDefault();
            }
            else if (area.data('plch')) {
               area.val('');
            }
         });

         /**
          * Replace tab input with a tab character or the correct number of spaces
          */
         function tab_replace(e) {
            var key = e.which;
            var sp = justice.tabr.space ? " " : "\t";
            var str = new Array(justice.tabr.num + 1).join(sp);
            if (key == 9 && !e.shiftKey && !e.ctrlKey && !e.altKey) {
               var os = area.scrollTop();
               var ta = area.get(0);
               if (ta.setSelectionRange) {
                  var ss = ta.selectionStart;
                  var se = ta.selectionEnd;
                  area.val(area.val().substring(0, ss) + str + area.val().substr(se));
                  ta.setSelectionRange(ss + str.length, se + str.length);
                  e.returnValue = false;
               }
               else if (area.createTextRange) {
                  document.selection.createRange().text = str;
                  e.returnValue = false;
               }
               //Fallback if we can't correctly create a range.  Just disallow tab replacement.
               else {
                  return true;
               }
               area.scrollTop(os);
               if (e.preventDefault) {
                  e.preventDefault();
               }
               return false;
            }
            return true;
         }

         /**
          * Placeholder handling.  Remove on focus, add on blur if supertextarea is empty
          */
         if (justice.plch.use) {
            if (!area.val().length) {
               if (justice.plch.css != undefined) {
                  area.css(justice.plch.css);
               }
               area.val(justice.plch.text);
               area.data('plch', true);
            }

            area.focus(function () {
               if (area.data('plch')) {
                  area.val('');
                  if (justice.css != undefined) {
                     area.css(justice.css);
                  }
                  area.data('plch', false);
               }
            });
            area.blur(function () {
               if (!area.val().length) {
                  area.data('plch', true);
                  if (justice.plch.css != undefined) {
                     area.css(justice.plch.css);
                  }
                  area.val(justice.plch.text);
               }
            });
         }

         area.css({'overflow':'auto'})
            .keydown(function (e) { update(e); })
            .keyup(function () { update(); })
            .bind('paste', function () { setTimeout(update, 250); });

         update();
      });
   }
})(jQuery);
