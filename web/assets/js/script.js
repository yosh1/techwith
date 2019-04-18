  // animation
  $("body").css({
      // left:"-100000px",
      opacity: "0",
  }).animate({
      left: "50px",
      opacity: "1"
  }, 1500);

  // <!-- make ace editor -->

  var editor = ace.edit("editor");
  editor.$blockScrolling = Infinity;
  editor.setOptions({
      exitOnOverlayClick: false,
      enableBasicAutocompletion: true,
      enableSnippets: true,
      enableLiveAutocompletion: true,
      fontSize: "11pt"
  });
  editor.setTheme("ace/theme/monokai");
  editor.setShowPrintMargin(false);
  editor.getSession().setMode("ace/mode/php");
  var EditSession = require("ace/edit_session").EditSession;

  // focus
  focusEditor();

  function getContent() {
      var content = editor.getValue();
      $('#terminal-body').text(content);
  }

  function showAlertPanel() {
      //$('.alert').alert('close');
      $('.alert').collapse('show');
  }


  function runReset() {
      editor.setValue("<?php\n\n?>");
      focusEditor();
  }


  function closeHint() {
      document.getElementById("answer-false").classList.remove("show");
  }

  var php = "";

  function suspend() {
      php = new EditSession(editor.getValue());
  }


  function suspend() {

  }


  function showLGTM() {

      // Get the snackbar DIV
      var toast = document.getElementById("answer-true");
      // Add the "show" class to DIV
      toast.className = "show";

  }


  function showPTAL() {
      // Get the snackbar DIV
      var toast = document.getElementById("answer-false");

      // Add the "show" class to DIV
      toast.className = "show";

      // After 3 seconds, remove the show class from DIV
      setTimeout(function () {
          toast.className = toast.className.replace("show", "");
      }, 5000);
  }


  // jQuery Key bind
  $(function ($) {
      //Ctrl+Enter or Command+Enter
      $(window).keydown(function (e) {
          if (event.ctrlKey || event.commandKey) {
              if (e.keyCode === 13) {
                  nextStep();
                  return false;
              }
          }
      });
  });

  function nextStep() {
      if ($('#answer-true').hasClass('show')) {
          // getNextQuestion();
          document.getElementById("next-button").click();
      } else {
          postProgram();
      }
  };

  function windowClose() {
      if ($('#answer-false').hasClass('show')) {
          closeHint();
      } else {
          postProgram();
      }
  };

  function focusEditor() {
      editor.focus();
      editor.gotoLine(2);
  }
